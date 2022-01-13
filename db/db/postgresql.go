package database

import (
	"database/sql"
	"fmt"

	"github.com/forbole/egldjuno/logging"

	"github.com/rs/zerolog/log"

	"github.com/cosmos/cosmos-sdk/simapp/params"

	"github.com/forbole/egldjuno/db"
	"github.com/forbole/egldjuno/types"
)

// Builder creates a database connection with the given database connection info
// from config. It returns a database connection handle or an error if the
// connection fails.
func Builder(cfg types.DatabaseConfig, encodingConfig *params.EncodingConfig) (db.Database, error) {
	sslMode := "disable"
	if cfg.GetSSLMode() != "" {
		sslMode = cfg.GetSSLMode()
	}

	schema := "public"
	if cfg.GetSchema() != "" {
		schema = cfg.GetSchema()
	}

	connStr := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s sslmode=%s search_path=%s",
		cfg.GetHost(), cfg.GetPort(), cfg.GetName(), cfg.GetUser(), sslMode, schema,
	)

	if cfg.GetPassword() != "" {
		connStr += fmt.Sprintf(" password=%s", cfg.GetPassword())
	}

	postgresDb, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Set max open connections
	postgresDb.SetMaxOpenConns(cfg.GetMaxOpenConnections())
	postgresDb.SetMaxIdleConns(cfg.GetMaxIdleConnections())

	return &Database{Sql: postgresDb, EncodingConfig: encodingConfig}, nil
}

// type check to ensure interface is properly implemented
var _ db.Database = &Database{}

// Database defines a wrapper around a SQL database and implements functionality
// for data aggregation and exporting.
type Database struct {
	Sql            *sql.DB
	EncodingConfig *params.EncodingConfig
	Logger         logging.Logger
}

// LastBlockHeight implements db.Database
func (db *Database) LastBlockHeight() (int64, error) {
	var height int64
	err := db.Sql.QueryRow(`SELECT coalesce(MAX(height),0) AS height FROM block;`).Scan(&height)
	return height, err
}

// HasBlock implements db.Database
func (db *Database) HasBlock(height int64) (bool, error) {
	var res bool
	err := db.Sql.QueryRow(`SELECT EXISTS(SELECT 1 FROM block WHERE height = $1);`, height).Scan(&res)
	return res, err
}

func (db *Database) SaveBlock(block []types.Block) error {
	stmt := `INSERT INTO block(hash,epoch,nonce,prev_hash,proposer,pub_key_bitmap,round,shard,size,size_txs,state_root_hash,time_stamp,tx_count,gas_consumed,gas_refunded,gas_penalized,max_gas_limit) VALUES `

	var params []interface{}

	for i, rows := range block {
		ai := i * 17
		stmt += fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d),", ai+1, ai+2, ai+3, ai+4, ai+5, ai+6, ai+7, ai+8, ai+9, ai+10, ai+11, ai+12, ai+13, ai+14, ai+15, ai+16, ai+17)

		params = append(params, rows.Hash, rows.Epoch, rows.Nonce, rows.PrevHash, rows.Proposer, rows.PubKeyBitmap, rows.Round, rows.Shard, rows.Size, rows.SizeTxs, rows.StateRootHash, rows.TimeStamp, rows.TxCount, rows.GasConsumed, rows.GasRefunded, rows.GasPenalized, rows.MaxGasLimit)

	}
	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT DO NOTHING`

	_, err := db.Sql.Exec(stmt, params...)
	if err != nil {
		return err
	}

	return nil
}


// HasValidator implements db.Database
func (db *Database) HasValidator(addr string) (bool, error) {
	var res bool
	stmt := `SELECT EXISTS(SELECT 1 FROM validator WHERE consensus_address = $1);`
	err := db.Sql.QueryRow(stmt, addr).Scan(&res)
	return res, err
}

// SaveValidators implements db.Database
func (db *Database) SaveValidators(validators []*types.Validator) error {
	if len(validators) == 0 {
		return nil
	}

	stmt := `INSERT INTO validator (consensus_address, consensus_pubkey) VALUES `

	var vparams []interface{}
	for i, val := range validators {
		vi := i * 2

		stmt += fmt.Sprintf("($%d, $%d),", vi+1, vi+2)
		vparams = append(vparams, val.ConsAddr, val.ConsPubKey)
	}

	stmt = stmt[:len(stmt)-1] // Remove trailing ,
	stmt += " ON CONFLICT DO NOTHING"
	_, err := db.Sql.Exec(stmt, vparams...)
	return err
}

func (db *Database) SaveNodeInfos(infos []*types.StakerNodeInfo) error {
	if len(infos) == 0 {
		return nil
	}

	stmt := `INSERT INTO node_info (
		id ,role,networkingAddress,networkingKey ,stakingKey ,tokensStaked ,
		tokensCommitted ,tokensUnstaking ,tokensUnstaked ,
		tokensRewarded ,delegators ,delegatorIDCounter ,
		tokensRequestedToUnstake, initialWeight
	) VALUES `

	var vparams []interface{}
	for i, val := range infos {
		vi := i * 13

		stmt += fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d),",
			vi+1, vi+2, vi+3, vi+4, vi+5, vi+6,
			vi+7, vi+8, vi+9, vi+10, vi+11, vi+12, vi+13)
		vparams = append(vparams, val.Id, val.Role, val.NetworkingAddress, val.NetworkingKey, val.StakingKey,
			val.TokensStaked, val.TokensCommitted, val.TokensUnstaking, val.TokensUnstaked, val.TokensRewarded,
			val.Delegators, val.DelegatorIDCounter, val.TokensRequestedToUnstake, val.InitialWeight)
	}

	stmt = stmt[:len(stmt)-1] // Remove trailing ,
	stmt += " ON CONFLICT DO NOTHING"
	_, err := db.Sql.Exec(stmt, vparams...)
	return err
}

// SaveCommitSignatures implements db.Database
func (db *Database) SaveCommitSignatures(signatures []*types.CommitSig) error {
	if len(signatures) == 0 {
		return nil
	}

	stmt := `INSERT INTO pre_commit (validator_address, height, timestamp, voting_power, proposer_priority) VALUES `

	var sparams []interface{}
	for i, sig := range signatures {
		si := i * 5

		stmt += fmt.Sprintf("($%d, $%d, $%d, $%d, $%d),", si+1, si+2, si+3, si+4, si+5)
		sparams = append(sparams, sig.ValidatorAddress, sig.Height, sig.Timestamp, sig.VotingPower, sig.ProposerPriority)
	}

	stmt = stmt[:len(stmt)-1]
	stmt += " ON CONFLICT (validator_address, timestamp) DO NOTHING"
	_, err := db.Sql.Exec(stmt, sparams...)
	return err
}


// Close implements db.Database
func (db *Database) Close() {
	err := db.Sql.Close()
	if err != nil {
		log.Error().Str("module", "psql database").Err(err).Msg("error while closing connection")
	}
}

// -------------------------------------------------------------------------------------------------------------------

// GetLastPruned implements db.PruningDb
func (db *Database) GetLastPruned() (int64, error) {
	var lastPrunedHeight int64
	err := db.Sql.QueryRow(`SELECT coalesce(MAX(last_pruned_height),0) FROM pruning LIMIT 1;`).Scan(&lastPrunedHeight)
	return lastPrunedHeight, err
}

// StoreLastPruned implements db.PruningDb
func (db *Database) StoreLastPruned(height int64) error {
	_, err := db.Sql.Exec(`DELETE FROM pruning`)
	if err != nil {
		return err
	}

	_, err = db.Sql.Exec(`INSERT INTO pruning (last_pruned_height) VALUES ($1)`, height)
	return err
}

// Prune implements db.PruningDb
func (db *Database) Prune(height int64) error {
	_, err := db.Sql.Exec(`DELETE FROM pre_commit WHERE height = $1`, height)
	if err != nil {
		return err
	}

	_, err = db.Sql.Exec(`
DELETE FROM message 
USING transaction 
WHERE message.transaction_hash = transaction.hash AND transaction.height = $1
`, height)
	return err
}

func (db *Database) SaveEvents(events []types.Event) error {
	if len(events) == 0 {
		return nil
	}

	stmt := `INSERT INTO event (
		height,type,transaction_id,transaction_index,event_index,value
	) VALUES `

	var vparams []interface{}
	for i, event := range events {
		vi := i * 6

		stmt += fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d),",
			vi+1, vi+2, vi+3, vi+4, vi+5, vi+6)
		vparams = append(vparams, event.Height, event.Type, event.TransactionID, event.TransactionIndex, event.EventIndex, event.Value.String())
	}

	stmt = stmt[:len(stmt)-1] // Remove trailing ,
	stmt += " ON CONFLICT DO NOTHING"
	_, err := db.Sql.Exec(stmt, vparams...)
	return err
}

func (db *Database) SaveCollection(collection []types.Collection) error {
	stmt := `INSERT INTO collection(height,id,processed,transaction_id) VALUES `

	var params []interface{}

	i := 0
	for _, rows := range collection {
		for _, txid := range rows.TransactionIds {
			ai := i * 4
			stmt += fmt.Sprintf("($%d,$%d,$%d,$%d),", ai+1, ai+2, ai+3, ai+4)
			params = append(params, rows.Height, rows.Id, rows.Processed, txid.String())
			i++
		}
	}
	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT DO NOTHING`

	_, err := db.Sql.Exec(stmt, params...)
	if err != nil {
		return err
	}

	return nil
}
func (db *Database) SaveTransactionResult(transactionResult []types.TransactionResult, height uint64) error {
	stmt := `INSERT INTO transaction_result(height,transaction_id,status,error) VALUES `

	var params []interface{}

	for i, rows := range transactionResult {
		ai := i * 4
		stmt += fmt.Sprintf("($%d,$%d,$%d,$%d),", ai+1, ai+2, ai+3, ai+4)

		params = append(params, height, rows.TransactionId, rows.Status, rows.Error)

	}
	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT DO NOTHING`

	_, err := db.Sql.Exec(stmt, params...)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) SaveTxs(transaction types.Txs) error {
    stmt:= `INSERT INTO tx(tx_hash,gas_limit,gas_price,gas_used,mini_block_hash,nonce,receiver,receiver_shard,round,sender,sender_shard,signature,status,value,fee,timestamp,data,token_identifier,action,scam_info) VALUES `

    var params []interface{}

	  for i, rows := range transaction{
      ai := i * 20
      stmt += fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d),", ai+1,ai+2,ai+3,ai+4,ai+5,ai+6,ai+7,ai+8,ai+9,ai+10,ai+11,ai+12,ai+13,ai+14,ai+15,ai+16,ai+17,ai+18,ai+19,ai+20)
      
      params = append(params,rows.TxHash,rows.GasLimit,rows.GasPrice,rows.GasUsed,rows.MiniBlockHash,rows.Nonce,rows.Receiver,rows.ReceiverShard,rows.Round,rows.Sender,rows.SenderShard,rows.Signature,rows.Status,rows.Value,rows.Fee,rows.Timestamp,rows.Data,rows.TokenIdentifier,rows.Action,rows.ScamInfo)

    }
	  stmt = stmt[:len(stmt)-1]
    stmt += ` ON CONFLICT DO NOTHING` 

    _, err := db.Sql.Exec(stmt, params...)
    if err != nil {
      return err
    }

    return nil 
    }
     