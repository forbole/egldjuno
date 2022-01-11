package utils
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"io"
	"github.com/tidwall/gjson"
	"strings"


	"github.com/rs/zerolog/log"

	db "github.com/forbole/egldjuno/db/postgresql"

	"github.com/forbole/egldjuno/client"

	elasticsearch "github.com/elastic/go-elasticsearch/v7"


)

// GetHistoricBlocks scroll blocks from elastic endpoint from the latest to oldest 
func GetHistoricBlocks(db *db.Db, client client.Proxy)error{

	var (
		batchNum int
		scrollID string
	)
	
	cfg := elasticsearch.Config{
		Addresses: []string{
		  "https://index.elrond.com/",
		},
		Transport: &http.Transport{
		  MaxIdleConnsPerHost:   10,
		  ResponseHeaderTimeout: time.Second,
		  },
		}
	es,err:=elasticsearch.NewClient(cfg)

	if err != nil {
		return err
	}

	// run a boolean search query
    res, err := es.Search(
		es.Search.WithIndex("blocks"),
		es.Search.WithSort("_doc"),
		es.Search.WithSize(10),
		es.Search.WithScroll(time.Minute),
	)

	if err!=nil{
		return err
	}

	json := read(res.Body)
	res.Body.Close()

	scrollID = gjson.Get(json, "_scroll_id").String()


	for {
		batchNum++

		// Perform the scroll request and pass the scrollID and scroll duration
		//
		res, err := es.Scroll(es.Scroll.WithScrollID(scrollID), es.Scroll.WithScroll(time.Hour))
		if err != nil {
			panic( err)
		}
		if res.IsError() {
			panic(res)
		}

		json := read(res.Body)
		res.Body.Close()

		// Extract the scrollID from response
		//
		scrollID = gjson.Get(json, "_scroll_id").String()

		// Extract the search results
		//
		hits := gjson.Get(json, "hits.hits")
		
		// Break out of the loop when there are no results
		if len(hits.Array()) < 1{
			fmt.Println("Finished scrolling") 
			break
		}

		for i,hit:=range hits.Array(){
			fmt.Println(i)
			_ =gjson.Get(hit.String(), "_source").String()
			//fmt.Println(r)
			//blockstr <- r

			/* TODO: Make it a type array*/
		}

		//
		
		/* TODO: INSERT INTO DB */
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