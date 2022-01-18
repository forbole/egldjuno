package types

type Account struct {
	Address  string `json:"address"`
	Balance  string `json:"balance"`
	Nonce    int    `json:"nonce"`
	Shard    int    `json:"shard"`
	ScamInfo struct {
	} `json:"scamInfo"`
	Code                     string `json:"code"`
	CodeHash                 string `json:"codeHash"`
	RootHash                 string `json:"rootHash"`
	TxCount                  int    `json:"txCount"`
	ScrCount                 int    `json:"scrCount"`
	Username                 string `json:"username"`
	DeveloperReward          string `json:"developerReward"`
	OwnerAddress             string `json:"ownerAddress"`
	DeployedAt               int    `json:"deployedAt"`
	IsUpgradeable            bool   `json:"isUpgradeable"`
	IsReadable               bool   `json:"isReadable"`
	IsPayable                bool   `json:"isPayable"`
	IsPayableBySmartContract bool   `json:"isPayableBySmartContract"`
}

type SCResult struct {
	Hash           string `json:"hash"`
	Timestamp      int    `json:"timestamp"`
	Nonce          int    `json:"nonce"`
	GasLimit       int    `json:"gasLimit"`
	GasPrice       int    `json:"gasPrice"`
	Value          string `json:"value"`
	Sender         string `json:"sender"`
	Receiver       string `json:"receiver"`
	RelayedValue   string `json:"relayedValue"`
	Data           string `json:"data"`
	PrevTxHash     string `json:"prevTxHash"`
	OriginalTxHash string `json:"originalTxHash"`
	CallType       string `json:"callType"`
	Logs           struct {
		Address string   `json:"address"`
		Events  []string `json:"events"`
	} `json:"logs"`
	ReturnMessage struct {
	} `json:"returnMessage"`
}
type TokenBalance struct {
	Address    string
	Identifier string
	Balance    string
}

// Equal tells whether v and w represent the same rows
func (v TokenBalance) Equal(w TokenBalance) bool {
	return v.Address == w.Address &&
		v.Identifier == w.Identifier &&
		v.Balance == w.Balance
}

// TokenBalance allows to build a new TokenBalance
func NewTokenBalance(
	address string,
	identifier string,
	balance string) TokenBalance {
	return TokenBalance{
		Address:    address,
		Identifier: identifier,
		Balance:    balance,
	}
}

type Token struct {
	Identifier     string `json:"identifier"`
	Name           string `json:"name"`
	Ticker         string `json:"ticker"`
	Owner          string `json:"owner"`
	Minted         string `json:"minted"`
	Burnt          string `json:"burnt"`
	Decimals       int    `json:"decimals"`
	IsPaused       bool   `json:"isPaused"`
	CanUpgrade     bool   `json:"canUpgrade"`
	CanMint        bool   `json:"canMint"`
	CanBurn        bool   `json:"canBurn"`
	CanChangeOwner bool   `json:"canChangeOwner"`
	CanPause       bool   `json:"canPause"`
	CanFreeze      bool   `json:"canFreeze"`
	CanWipe        bool   `json:"canWipe"`
	Balance        string `json:"balance"`
}

type NFT struct {
	Identifier string   `json:"identifier"`
	Collection string   `json:"collection"`
	Attributes string   `json:"attributes"`
	Nonce      int      `json:"nonce"`
	Type       string   `json:"type"`
	Name       string   `json:"name"`
	Creator    string   `json:"creator"`
	Royalties  int      `json:"royalties"`
	Uris       []string `json:"uris"`
	URL        string   `json:"url"`
	Media      []struct {
		URL          string `json:"url"`
		OriginalURL  string `json:"originalUrl"`
		ThumbnailURL string `json:"thumbnailUrl"`
		FileType     string `json:"fileType"`
		FileSize     int    `json:"fileSize"`
	} `json:"media"`
	IsWhitelistedStorage bool `json:"isWhitelistedStorage"`
	Metadata             struct {
	} `json:"metadata"`
	ScamInfo struct {
		Type string `json:"type"`
		Info string `json:"info"`
	} `json:"scamInfo"`
	Balance string `json:"balance"`
	Ticker  string `json:"ticker"`
}

type AccountContract struct {
	Address      string `json:"address"`
	DeployTxHash string `json:"deployTxHash"`
	Timestamp    int    `json:"timestamp"`
}
