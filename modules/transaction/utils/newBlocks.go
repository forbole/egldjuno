package utils

import (
	"encoding/json"
	"fmt"

	"github.com/forbole/egldjuno/types"
	"github.com/tidwall/gjson"

	"github.com/forbole/egldjuno/client"
)

func GetNewTransactions(client client.Proxy) ([]types.Tx,error){
	txsParams:=map[string]string{
		"size":"25",
		"fields":"txHash",
	}
	txHashesRaw,err:=client.RestRequestGet("transactions", txsParams)
	if err!=nil{
		return nil,err
	}
	type txHashes []struct {
		TxHash string `json:"txHash"`
	}
	var txhash txHashes
	err=json.Unmarshal(txHashesRaw,&txhash)
	if err!=nil{
		return nil,err
	}

	mainTxs:=make([]types.Tx,25)
	for i,tx:=range txhash{
		fmt.Println(tx)
		jsonstr,err:=client.RestRequestGet(fmt.Sprintf("transactions/%s",tx.TxHash),nil)
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
	fmt.Println(jsStr)
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
