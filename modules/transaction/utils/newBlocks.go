package utils

import (
	"fmt"

	"github.com/forbole/egldjuno/types"
	"github.com/tidwall/gjson"


	"github.com/forbole/egldjuno/client"
)

func GetNewTransactions(client client.Proxy) ([]types.Tx,error){
	type txHash struct{
		txHash string
	} 
	
	var txHashes []txHash
	txsParams:=map[string]string{
		"size":"25",
		"fields":"txHash",
	}
	err:=client.RestRequestGetDecoded("transactions", txsParams, &txHashes)

	mainTxs:=make([]types.Tx,25)
	for i,tx:=range txHashes{
		jsonstr,err:=client.RestRequestGet(fmt.Sprintf("transactions/%s",tx.txHash),nil)
		if err!=nil{
			return nil,err
		}
		mainTxs[i]=decodeTx(jsonstr)
		
	}

	if err!=nil{
		return nil,err
	}
	return mainTxs,nil
}

// decodeTx Get a json str from transactions/{hash} endpoint and get the data that exist on every tx
func decodeTx(jsonstr []byte)types.Tx{
	jsStr:=string(jsonstr)
	tx:=types.NewTx(
		gjson.Get(jsStr, "txHash").String(),
		gjson.Get(jsStr, "gasLimit").Int(),
		gjson.Get(jsStr, "gasPrice").Int(),
		gjson.Get(jsStr, "gasUsed").Int(),
		gjson.Get(jsStr, "miniBlockHash").String(),
		gjson.Get(jsStr, "nonce").Int(),
		gjson.Get(jsStr, "receiver").String(),
		gjson.Get(jsStr, "receiverShard").Int(),
		gjson.Get(jsStr, "round").Int(),
		gjson.Get(jsStr, "sender").String(),
		gjson.Get(jsStr, "senderShard").Int(),
		gjson.Get(jsStr, "signature").String(),
		gjson.Get(jsStr, "status").String(),
		gjson.Get(jsStr, "value").String(),
		gjson.Get(jsStr, "fee").String(),
		gjson.Get(jsStr, "timestamp").Int(),
		gjson.Get(jsStr, "data").String(),
	)
	return tx
}
