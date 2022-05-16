package bothook

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func Reply(resurl []string) {
	fmt.Printf("Will try to send to the responce URL: %s \n", resurl[0])

	postBody, _ := json.Marshal(map[string]string{
		"name":  "GeoCODES user",
		"email": "gcu@example.com",
	})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(resurl[0], "application/json", responseBody)
	//resp.Body

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}