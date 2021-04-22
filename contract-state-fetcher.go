package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const requestURL = "https://eos.eoscafeblock.com/v1/chain/get_account"

type requestStateBody struct {
	ContractName string `json:"account_name"`
}

func stateFetcher(contractName string) (response Account, err error) {
	var contractResponse *http.Response

	requestBody := requestStateBody{
		ContractName: contractName,
	}

	var requestBodyBytes []byte
	if requestBodyBytes, err = json.Marshal(requestBody); err != nil {
		return Account{}, err
	}

	if contractResponse, err = http.Post(requestURL, "application/json", bytes.NewReader(requestBodyBytes)); err != nil {
		return Account{}, err
	}
	defer contractResponse.Body.Close()

	if err = json.NewDecoder(contractResponse.Body).Decode(&response); err != nil {
		return Account{}, err
	}

	return
}
