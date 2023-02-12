package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const API string = `https://evilinsult.com/generate_insult.php?lang=en&type=json`

type evilInsults struct {
	Number    string
	Language  string
	Insult    string
	Created   string
	Shown     string
	Createdby string
	Active    string
	Comment   string
}

func main() {
	evilJokes := new(evilInsults)

	// fetch data
	req, err := http.NewRequest(http.MethodGet, API, nil)
	if err != nil {
		log.Fatal(err)
	}

	// get response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// convert to bytes
	toBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// unmarshal response
	if err = json.Unmarshal(toBytes, &evilJokes); err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(evilJokes.Insult))

}
