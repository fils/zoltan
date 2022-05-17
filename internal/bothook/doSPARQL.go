package bothook

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func DoSPARQL(query []string) []byte {

	fmt.Printf("Try to do SPARQL call with %s\n", query[0])

	d := fmt.Sprintf(`prefix schema: <http://schema.org/>
				SELECT distinct ?subj ?score ?name ?desc ?lit	
						WHERE {
							?lit bds:search "%s" .
							?lit bds:relevance ?score .
							filter( ?score > 0.4).
							?subj ?p ?lit .
							?subj schema:name ?name . 
							?subj schema:description ?desc . 
				}
				   ORDER BY DESC(?score)
				LIMIT 5`, query[0])
	//d := fmt.Sprint("SELECT distinct ?subj  WHERE {?lit bds:search 'test' . ?subj ?p ?lit .}")

	spql := "https://graph.geodex.org/blazegraph/namespace/nabu/sparql"

	pab := []byte("")
	params := url.Values{}
	params.Add("query", d)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?%s", spql, params.Encode()), bytes.NewBuffer(pab))
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Accept", "application/sparql-results+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	//defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(strings.Repeat("ERROR", 5))
		log.Println("response Status:", resp.Status)
		log.Println("response Headers:", resp.Header)
		log.Println("response Body:", string(body))
	}

	return body
}
