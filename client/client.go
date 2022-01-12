package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"net/http"

	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/forbole/egldjuno/types"

	"google.golang.org/grpc"

	elasticsearch "github.com/elastic/go-elasticsearch/v7"
)

// Proxy implements a wrapper around both a Tendermint RPC client and a
// Cosmos Sdk REST client that allows for essential data queries.
type Proxy struct {
	ctx            context.Context
	encodingConfig *params.EncodingConfig
	apiClient *httpClient
	elasticClient *elasticsearch.Client
	grpConnection   *grpc.ClientConn
	txServiceClient tx.ServiceClient
	genesisHeight   uint64
}

type httpClient struct{
	Client *http.Client
	ApiAdress string
}

// NewClientProxy allows to build a new Proxy instance
func NewClientProxy(cfg types.Config, encodingConfig *params.EncodingConfig) (*Proxy, error) {
	address:=cfg.GetRPCConfig().GetAddress()
	client := &http.Client{Timeout: 10 * time.Second}
	
	httpClient  := &httpClient{
		Client:client,
		ApiAdress:address,
	}

	elasticCfg := elasticsearch.Config{
		Addresses: []string{
		  "https://index.elrond.com/",
		},
		Transport: &http.Transport{
		  MaxIdleConnsPerHost:   10,
		  ResponseHeaderTimeout: time.Hour,
		  },
		}
	es,err:=elasticsearch.NewClient(elasticCfg)

	if err != nil {
		panic(err)
	}

	return &Proxy{
		encodingConfig:  encodingConfig,
		ctx:             context.Background(),
		apiClient:       httpClient,
		elasticClient: es,
		grpConnection:   nil,
		txServiceClient: nil,
		genesisHeight:   cfg.GetCosmosConfig().GetGenesisHeight(),
	}, nil
}

// QueryLCD queries the LCD at the given endpoint, and deserializes the result into the given pointer.
// If an error is raised, returns the error.
func (cp Proxy) restRequestGet(endpoint string,values map[string]string)([]byte,error) {
	jsonData, err := json.Marshal(values)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s",cp.apiClient.ApiAdress,endpoint), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil,err
	}

	response, err := cp.apiClient.Client.Do(req)
	if err != nil {
		return nil,err
	}

	// Close the connection to reuse it
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil,err
	}

	return body,nil
}

func (cp Proxy) RestRequestGetDecoded(endpoint string,values map[string]string,ptr interface{})error{
	bz,err:= cp.restRequestGet(endpoint,values)
	if err!=nil{
		return err
	}
	
	return json.Unmarshal(bz, &ptr)
}

// GetGeneisisBlock parse the specific block as genesis block
func (cp *Proxy) GetGenesisHeight() uint64 {
	return cp.genesisHeight
}


func (cp *Proxy) ElasticClient()*elasticsearch.Client{
	return cp.elasticClient
}



// LatestHeight returns the latest block height on the active chain. An error
// is returned if the query fails.
func (cp *Proxy) LatestHeight() (int64, error) {

	var blocks []types.Block
	err := cp.RestRequestGetDecoded("blocks",nil,&blocks)
	if err != nil {
		return -1, err
	}

	height := int64(blocks[0].Round)
	return height, nil
}
/*
// Block queries for a block by height. An error is returned if the query fails.
func (cp *Proxy) Block(height int64) (*flow.Block, error) {
	params:=map[string]string{
		"nonce": strconv.FormatInt(height,10),
	}
	block:=types.Block{}
	err := cp.RestRequestGetDecoded("blocks",params,block)
	if err != nil {
		return nil, err
	}
	return block, nil
}

// GetTransaction queries for a transaction by hash. An error is returned if the
// query fails.
func (cp *Proxy) GetTransaction(hash string) (*flow.Transaction, error) {
	transaction, err := cp.flowClient.GetTransaction(cp.ctx, flow.HashToID([]byte(hash)))
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
 */
/*
// NodeOperators returns all the known flow node operators for a given block
// height. An error is returned if the query fails.
func (cp *Proxy) NodeOperators(height int64) (*types.NodeOperators, error) {
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(): [FlowIDTableStaking.NodeInfo] {
		let nodes:[FlowIDTableStaking.NodeInfo]=[]
		for node in FlowIDTableStaking.getStakedNodeIDs() {
			nodes.append(FlowIDTableStaking.NodeInfo(node))
		}
		return nodes
	}`, cp.contract.StakingTable)

	result, err := cp.flowClient.ExecuteScriptAtLatestBlock(cp.ctx, []byte(script), nil)
	if err != nil {
		return nil, err
	}
	value := result.ToGoValue()
	nodes, ok := value.([]interface{})
	if !ok {
		return nil, fmt.Errorf("candance value cannot change to valid []interface{}")
	}
	nodeInfos := make([]*types.NodeInfo, len(nodes))
	for i, node := range nodes {
		nodeInfo, err := types.NewNodeInfoFromCandance(node)
		if err != nil {
			return nil, err
		}
		nodeInfos[i] = &nodeInfo
	}

	nodeOperators := types.NewNodeOperators(height, nodeInfos)
	return &nodeOperators, nil
} */

/* 
func (cp *Proxy) GetChainID() string {
	// There is GetNetworkParameters rpc method that not implenment yet in flow-go-sdk.
	return cp.contract.ChainID
} 
 */
/*
// Genesis returns the genesis state
func (cp *Proxy) Genesis() (*tmctypes.ResultGenesis, error) {
	return cp.flowClient.Genesis(cp.ctx)
}

// ConsensusState returns the consensus state of the chain
func (cp *Proxy) ConsensusState() (*constypes.RoundStateSimple, error) {
	state, err := cp.flowClient.ConsensusState(context.Background())
	if err != nil {
		return nil, err
	}

	var data constypes.RoundStateSimple
	err = tmjson.Unmarshal(state.RoundState, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
*/
// SubscribeEvents subscribes to new events with the given query through the RPC
// client with the given subscriber name. A receiving only channel, context
// cancel function and an error is returned. It is up to the caller to cancel
// the context and handle any errors appropriately.
/* func (cp *Proxy) SubscribeEvents(subscriber, query string) (<-chan tmctypes.ResultEvent, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	eventCh, err := cp.flowClient.event
	return eventCh, cancel, err
}

// SubscribeNewBlocks subscribes to the new block event handler through the RPC
// client with the given subscriber name. An receiving only channel, context
// cancel function and an error is returned. It is up to the caller to cancel
// the context and handle any errors appropriately.
func (cp *Proxy) SubscribeNewBlocks(subscriber string) (<-chan tmctypes.ResultEvent, context.CancelFunc, error) {
	return cp.SubscribeEvents(subscriber, "tm.event = 'NewBlock'")
} */

// Collections get all the collection from block
/* func (cp *Proxy) Collections(block *flow.Block) []types.Collection {
	collectionsID := block.CollectionGuarantees
	collections := make([]types.Collection, len(block.CollectionGuarantees))
	for i, c := range collectionsID {
		processed := true
		collection, err := cp.flowClient.GetCollection(cp.ctx, c.CollectionID)

		if err != nil {
			processed = false
		}

		collections[i] = types.NewCollection(block.Height, collection.ID().String(), processed, collection.TransactionIDs)
	}
	return collections
} */

// Txs queries for all the transactions in a block. Transactions are returned
// in the TransactionResult format which internally contains an array of Transactions. An error is
// returned if any query fails.
/* func (cp *Proxy) Txs(block *flow.Block) (types.Txs, error) {

	var transactionIDs []flow.Identifier
	collections := cp.Collections(block)
	for _, collection := range collections {
		transactionIDs = append(transactionIDs, (collection.TransactionIds)...)
	}

	txResponses := make([]types.Tx, len(transactionIDs))
	for i, txID := range transactionIDs {
		transaction, err := cp.flowClient.GetTransaction(cp.ctx, txID)
		if err != nil {
			return nil, err
		}

		authoriser := make([]string, len(transaction.Authorizers))
		for i, auth := range transaction.Authorizers {
			authoriser[i] = auth.String()
		}

		payloadSignitures, err := json.Marshal(transaction.PayloadSignatures)
		if err != nil {
			return nil, err
		}

		envelopeSigniture, err := json.Marshal(transaction.EnvelopeSignatures)
		if err != nil {
			return nil, err
		}

		tx := types.NewTx(block.Height, txID.String(), transaction.Script, transaction.Arguments,
			transaction.ReferenceBlockID.String(), transaction.GasLimit, transaction.ProposalKey.Address.String(), transaction.Payer.String(),
			authoriser, payloadSignitures, envelopeSigniture)

		txResponses[i] = tx
	}
	return txResponses, nil
}

func (cp *Proxy) TransactionResult(transactionIds []flow.Identifier) ([]types.TransactionResult, error) {
	if len(transactionIds) == 0 {
		return nil, nil
	}

	txResults := make([]types.TransactionResult, len(transactionIds))
	for i, txid := range transactionIds {
		result, err := cp.flowClient.GetTransactionResult(cp.ctx, txid)
		if err != nil {
			return nil, err
		}
		errStr := ""
		if result.Error != nil {
			errStr = result.Error.Error()
		}
		txResults[i] = types.NewTransactionResult(txid.String(), result.Status.String(), errStr)
	}
	return txResults, nil
}

func (cp *Proxy) EventsInBlock(block *flow.Block) ([]types.Event, error) {
	txs, err := cp.Txs(block)
	if err != nil {
		return nil, err
	}
	var event []types.Event
	for _, tx := range txs {
		ev, err := cp.Events(tx.TransactionID, int(tx.Height))
		if err != nil {
			return []types.Event{}, err
		}
		event = append(event, ev...)
	}
	return event, nil
}

func (cp *Proxy) EventsInTransaction(tx types.Tx) ([]types.Event, error) {
	var event []types.Event

	ev, err := cp.Events(tx.TransactionID, int(tx.Height))
	if err != nil {
		return []types.Event{}, err
	}
	event = append(event, ev...)
	return event, nil
}

// Events get events from a transaction ID
func (cp *Proxy) Events(transactionID string, height int) ([]types.Event, error) {
	transactionResult, err := cp.flowClient.GetTransactionResult(cp.ctx, flow.HexToID(transactionID))
	if err != nil {
		return []types.Event{}, err
	}

	ev := make([]types.Event, len(transactionResult.Events))
	for i, event := range transactionResult.Events {
		ev[i] = types.NewEvent(height, event.Type, event.TransactionID.String(), event.TransactionIndex,
			event.EventIndex, event.Value)
	}
	return ev, nil
}
*/
// Stop defers the node stop execution to the RPC client.
func (cp *Proxy) Stop() {
	

}
 