package types

type Account struct {
	Address         string `json:"address"`
	Nonce           int    `json:"nonce"`
	Balance         string `json:"balance"`
	RootHash        string `json:"rootHash"`
	TxCount         int    `json:"txCount"`
	ScrCount        int    `json:"scrCount"`
	Shard           int    `json:"shard"`
	DeveloperReward string `json:"developerReward"`
}
