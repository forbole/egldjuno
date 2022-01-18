package utils

import (
	"fmt"

	"github.com/forbole/egldjuno/types"

	"github.com/forbole/egldjuno/client"
	db "github.com/forbole/egldjuno/db/postgresql"
)

func GetNewBlocks(client client.Proxy) ([]types.Block, error) {
	var blocks []types.Block
	err := client.RestRequestGetDecoded("blocks", nil, &blocks)
	if err != nil {
		return nil, err
	}
	return blocks, nil
}

func GetNewAccounts(db *db.Db, client client.Proxy) error {

	size := 25
	addresses, err := getAccountsList(client, size)
	if err != nil {
		return err
	}

	for _, a := range addresses {
		fmt.Println(fmt.Sprintf("Address:%s", a))
		account, err := getAccountDetails(a, client)
		if err != nil {
			return err
		}
		err = db.SaveAccount([]types.Account{*account})

		tokens, err := getTokensBalance(a, client)
		if err != nil {
			return err
		}

		err = db.SaveTokenBalance(tokens, a)
		if err != nil {
			return err
		}

		nfts, err := getNFTs(a, client)
		if err != nil {
			return err
		}

		err = db.SaveAccountNft(nfts, account.Address)
		if err != nil {
			return err
		}

		contracts, err := getContracts(a, client)
		if err != nil {
			return err
		}

		err = db.SaveAccountContract(contracts, a)
		if err != nil {
			return err
		}

	}
	return nil
}

func getAccountsList(client client.Proxy, size int) ([]string, error) {
	txsParams := map[string]string{
		"size":   string(size),
		"fields": "address",
	}
	type AddressRow []struct {
		Address string `json:"address"`
	}
	var addresses AddressRow
	err := client.RestRequestGetDecoded("accounts", txsParams, &addresses)
	if err != nil {
		return nil, err
	}
	returnAddresses := make([]string, 25)
	for i, a := range addresses {
		returnAddresses[i] = a.Address
	}
	return returnAddresses, nil
}

func getAccountDetails(address string, client client.Proxy) (*types.Account, error) {
	var account types.Account
	err := client.RestRequestGetDecoded(fmt.Sprintf("accounts/%s", address), nil, &account)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func getScResults(address string, client client.Proxy) ([]types.SCResult, error) {
	var scResult []types.SCResult
	err := client.RestRequestGetDecoded(fmt.Sprintf("accounts/%s/sc-results", address), nil, &scResult)
	if err != nil {
		return nil, err
	}
	return scResult, nil
}

func getTokensBalance(address string, client client.Proxy) ([]types.Token, error) {
	var token []types.Token
	err := client.RestRequestGetDecoded(fmt.Sprintf("accounts/%s/tokens", address), nil, &token)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func getNFTs(address string, client client.Proxy) ([]types.NFT, error) {
	var nfts []types.NFT
	err := client.RestRequestGetDecoded(fmt.Sprintf("accounts/%s/nfts?includeFlagged=true", address), nil, &nfts)
	if err != nil {
		return nil, err
	}
	return nfts, nil
}

func getContracts(address string, client client.Proxy) ([]types.AccountContract, error) {
	var contracts []types.AccountContract
	err := client.RestRequestGetDecoded(fmt.Sprintf("accounts/%s/contracts", address), nil, &contracts)
	if err != nil {
		return nil, err
	}
	return contracts, nil
}
