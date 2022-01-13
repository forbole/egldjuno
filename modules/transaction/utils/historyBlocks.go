package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/tidwall/gjson"

	db "github.com/forbole/egldjuno/db/postgresql"
	"github.com/forbole/egldjuno/types"
	"github.com/rs/zerolog/log"

	"github.com/forbole/egldjuno/client"
)

// GetHistoricBlocks scroll blocks from elastic endpoint from the latest to oldest
func GetHistoricBlocks(db *db.Db, client client.Proxy) error {
	log.Info().Msg("Enqueuing historic block")
	var (
		batchNum int
		scrollID string
	)
	es := client.ElasticClient()

	// run a boolean search query
	res, err := es.Search(
		es.Search.WithIndex("blocks"),
		es.Search.WithSort("_doc"),
		es.Search.WithSize(10),
		es.Search.WithScroll(time.Minute),
	)

	if err != nil {
		return err
	}

	jsonstr := read(res.Body)
	res.Body.Close()

	scrollID = gjson.Get(jsonstr, "_scroll_id").String()
	fmt.Println(jsonstr)

	hits := gjson.Get(jsonstr, "hits.hits")

	blocks := make([]types.Block, len(hits.Array()))
	for i, hit := range hits.Array() {
		//fmt.Println(i)
		blockjson := gjson.Get(hit.String(), "_source").String()
		//fmt.Println(blockjson)
		/* TODO: Make it a type array*/
		var block types.Block
		json.Unmarshal([]byte(blockjson), &block)
		blocks[i] = block
		//fmt.Println(block.Epoch)
	}
	err = db.SaveBlock(blocks)
	if err != nil {
		return err
	}

	for {
		batchNum++

		// Perform the scroll request and pass the scrollID and scroll duration
		//
		res, err := es.Scroll(es.Scroll.WithScrollID(scrollID), es.Scroll.WithScroll(time.Hour))
		if err != nil {
			panic(err)
		}
		if res.IsError() {
			panic(res)
		}

		jsonstr := read(res.Body)
		res.Body.Close()

		// Extract the scrollID from response
		//
		scrollID = gjson.Get(jsonstr, "_scroll_id").String()

		// Extract the search results
		//
		hits := gjson.Get(jsonstr, "hits.hits")

		// Break out of the loop when there are no results
		if len(hits.Array()) < 1 {
			fmt.Println("Finished scrolling")
			break
		}

		blocks := make([]types.Block, len(hits.Array()))
		for i, hit := range hits.Array() {
			//fmt.Println(i)
			blockjson := gjson.Get(hit.String(), "_source").String()

			/* TODO: Make it a type array*/
			var block types.Block
			json.Unmarshal([]byte(blockjson), &block)
			blocks[i] = block
			//fmt.Println(block.Epoch)
		}
		/* TODO: INSERT INTO DB */
		err = db.SaveBlock(blocks)
		if err != nil {
			return err
		}
	}
	defer res.Body.Close()
	fmt.Println("resultBody \n")
	fmt.Println(res)

	return nil
}

func read(r io.Reader) string {
	var b bytes.Buffer
	b.ReadFrom(r)
	return b.String()
}
