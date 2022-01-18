package postgresql

import (
	"fmt"

	"github.com/forbole/egldjuno/types"
)

func (db *Db) SaveAccount(account []types.Account) error {
	stmt := `INSERT INTO account(address,nonce,balance,root_hash,tx_count,scr_count,shard,developer_reward) VALUES `

	var params []interface{}

	for i, rows := range account {
		ai := i * 8
		stmt += fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d),", ai+1, ai+2, ai+3, ai+4, ai+5, ai+6, ai+7, ai+8)

		params = append(params, rows.Address, rows.Nonce, rows.Balance, rows.RootHash, rows.TxCount, rows.ScrCount, rows.Shard, rows.DeveloperReward)

	}
	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT DO NOTHING`

	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return err
	}

	return nil
}

func (db *Db) SaveToken(token []types.Token) error {
	stmt := `INSERT INTO token(identifier,name,ticker,owner,minted,burnt,decimals,is_paused,can_upgrade,can_mint,can_burn,can_change_owner,can_pause,can_freeze,can_wipe,balance) VALUES `

	var params []interface{}

	for i, rows := range token {
		ai := i * 16
		stmt += fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d),", ai+1, ai+2, ai+3, ai+4, ai+5, ai+6, ai+7, ai+8, ai+9, ai+10, ai+11, ai+12, ai+13, ai+14, ai+15, ai+16)

		params = append(params, rows.Identifier, rows.Name, rows.Ticker, rows.Owner, rows.Minted, rows.Burnt, rows.Decimals, rows.IsPaused, rows.CanUpgrade, rows.CanMint, rows.CanBurn, rows.CanChangeOwner, rows.CanPause, rows.CanFreeze, rows.CanWipe, rows.Balance)

	}
	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT DO NOTHING`

	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return err
	}

	return nil
}

func (db *Db) SaveTokenBalance(tokenBalance []types.Token, address string) error {
	stmt := `INSERT INTO token_balance(address,identifier,balance) VALUES `

	var params []interface{}

	for i, rows := range tokenBalance {
		ai := i * 3
		stmt += fmt.Sprintf("($%d,$%d,$%d),", ai+1, ai+2, ai+3)

		params = append(params, address, rows.Identifier, rows.Balance)

	}
	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT DO NOTHING`

	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return err
	}

	return nil
}

func (db *Db) SaveAccountNft(accountNft []types.NFT, address string) error {
	stmt := `INSERT INTO account_nft(address,identifier) VALUES `

	var params []interface{}

	for i, rows := range accountNft {
		ai := i * 2
		stmt += fmt.Sprintf("($%d,$%d),", ai+1, ai+2)

		params = append(params, address, rows.Identifier)

	}
	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT DO NOTHING`

	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return err
	}

	return nil
}

func (db *Db) SaveAccountContract(accountContract []types.AccountContract, address string) error {
	stmt := `INSERT INTO account_contract(address,contract_address,deploy_tx_hash,timestamp) VALUES `

	var params []interface{}

	for i, rows := range accountContract {
		ai := i * 4
		stmt += fmt.Sprintf("($%d,$%d,$%d,$%d),", ai+1, ai+2, ai+3, ai+4)

		params = append(params, address, rows.Address, rows.DeployTxHash, rows.Timestamp)

	}
	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT DO NOTHING`

	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return err
	}

	return nil
}
