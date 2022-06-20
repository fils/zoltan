package bothook

import (
	"encoding/json"
	"fmt"
)

// sparql struct

type SPARQLresp struct {
	Head    Head    `json:"head"`
	Results Results `json:"results"`
}
type Head struct {
	Vars []string `json:"vars"`
}
type Subj struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
type Score struct {
	Datatype string `json:"datatype"`
	Type     string `json:"type"`
	Value    string `json:"value"`
}
type Name struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
type Desc struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
type Lit struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
type Bindings struct {
	Subj  Subj  `json:"subj"`
	Score Score `json:"score"`
	Name  Name  `json:"name"`
	Desc  Desc  `json:"desc"`
	Lit   Lit   `json:"lit"`
}
type Results struct {
	Bindings []Bindings `json:"bindings"`
}

// slack struct

type Autogenerated struct {
	Blocks []Blocks `json:"blocks"`
}
type Text struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
type Blocks struct {
	Type string `json:"type"`
	Text Text   `json:"text,omitempty"`
}

func BuildReply(postBody []byte) []byte {
	res := SPARQLresp{}

	if err := json.Unmarshal(postBody, &res); err != nil {
		panic(err)
	}

	for x := range res.Results.Bindings {
		fmt.Printf("------------ Result %d --------------\n%s\n%s\n\n", x, res.Results.Bindings[x].Subj, res.Results.Bindings[x].Name)
	}

	return []byte("{'head':'test'}")
}

func reitem() string {
	return "*<fakeLink.toHotelPage.com|Windsor Court Hotel>*\\n$340 per night\\nRated: 9.4 - Excellent"
}

func header() string {
	return "Thanks for using GeoCODES.  The following are the top 3 results  See more results at <https://geocodes.org|GeoCODES>"
}
