package utils

import (
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
