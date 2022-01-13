package postgresql

import (
	"fmt"

	"github.com/forbole/egldjuno/types"
)

func (db *Db) SaveTxs(transaction types.Txs) error {
    stmt:= `INSERT INTO tx(tx_hash,gas_limit,gas_price,gas_used,mini_block_hash,nonce,receiver,receiver_shard,round,sender,sender_shard,signature,status,value,fee,timestamp,data) VALUES `

    var params []interface{}

	  for i, rows := range transaction{
      ai := i * 20
      stmt += fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d),", ai+1,ai+2,ai+3,ai+4,ai+5,ai+6,ai+7,ai+8,ai+9,ai+10,ai+11,ai+12,ai+13,ai+14,ai+15,ai+16,ai+17)
      
      params = append(params,rows.TxHash,rows.GasLimit,rows.GasPrice,rows.GasUsed,rows.MiniBlockHash,rows.Nonce,rows.Receiver,rows.ReceiverShard,rows.Round,rows.Sender,rows.SenderShard,rows.Signature,rows.Status,rows.Value,rows.Fee,rows.Timestamp,rows.Data)

    }
	  stmt = stmt[:len(stmt)-1]
    stmt += ` ON CONFLICT DO NOTHING` 

    _, err := db.Sqlx.Exec(stmt, params...)
    if err != nil {
      return err
    }

    return nil 
    }

	func (db *Db) SaveSmartContractResult(smartContractResult []types.SmartContractResult) error {
		stmt:= `INSERT INTO smart_contract_result(tx_hash,hash,timestamp,nonce,gas_limit,gas_price,value,sender,receiver,relayed_value,data,prev_tx_hash,original_tx_hash,call_type,logs) VALUES `
	
		var params []interface{}
	
		  for i, rows := range smartContractResult{
		  ai := i * 15
		  stmt += fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d),", ai+1,ai+2,ai+3,ai+4,ai+5,ai+6,ai+7,ai+8,ai+9,ai+10,ai+11,ai+12,ai+13,ai+14,ai+15)
		  
		  params = append(params,rows.TxHash,rows.Hash,rows.Timestamp,rows.Nonce,rows.GasLimit,rows.GasPrice,rows.Value,rows.Sender,rows.Receiver,rows.RelayedValue,rows.Data,rows.PrevTxHash,rows.OriginalTxHash,rows.CallType,rows.Logs)
	
		}
		  stmt = stmt[:len(stmt)-1]
		stmt += ` ON CONFLICT DO NOTHING` 
	
		_, err := db.Sqlx.Exec(stmt, params...)
		if err != nil {
		  return err
		}
	
		return nil 
		}
     