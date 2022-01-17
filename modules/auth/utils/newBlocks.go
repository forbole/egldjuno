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

func getScResults(address string,client client.Proxy) ([]types.SCResult,error){
	var scResult []types.SCResult
	err := client.RestRequestGetDecoded(fmt.Sprintf("accounts/%s/sc-results",address), nil, &scResult)
	if err!=nil{
		return nil,err
	}
	return scResult,nil
}

func getTokens(address string,client client.Proxy) ([]types.Token,error){
	var scResult []types.Token
	err := client.RestRequestGetDecoded(fmt.Sprintf("accounts/%s/tokens",address), nil, &scResult)
	if err!=nil{
		return nil,err
	}
	return scResult,nil
}

func getNFTs(address string,client client.Proxy) ([]types.NFT,error){
	var nfts []types.NFT
	err := client.RestRequestGetDecoded(fmt.Sprintf("accounts/%s/nfts?includeFlagged=true",address), nil, &nfts)
	if err!=nil{
		return nil,err
	}
	return nfts,nil
}

func getContracts(address string,client client.Proxy) ([]types.NFT,error){
	var nfts []types.NFT
	err := client.RestRequestGetDecoded(fmt.Sprintf("accounts/%s/contracts",address), nil, &nfts)
	if err!=nil{
		return nil,err
	}
	return nfts,nil
}