package utils

import (
	"fmt"

	"github.com/forbole/egldjuno/types"

	"github.com/forbole/egldjuno/client"
)

func GetNewBlocks(client client.Proxy) ([]types.Block, error) {
	var blocks []types.Block
	err := client.RestRequestGetDecoded("blocks", nil, &blocks)
	if err != nil {
		return nil, err
	}
	return blocks, nil
}

func GetNewAccounts(client client.Proxy) ([]types.Account, error) {
	txsParams := map[string]string{
		"size":   "25",
		"fields": "address",
	}
	type AddressRow []struct {
		Address string `json:"address"`
	}
	var addresses AddressRow
	err := client.RestRequestGetDecoded("accounts", txsParams, &addresses)

	accounts:=make([]types.Account,25)
	for i,a:=range addresses{
		fmt.Println(fmt.Sprintf("Address:%s",a.Address))
		var account types.Account
		err = client.RestRequestGetDecoded(fmt.Sprintf("accounts/%s",a.Address), nil, &account)
		if err!=nil{
			return nil,err
		}
		accounts[i]=account
	}
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
