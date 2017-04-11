package openwhisk

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type OpenWhisk struct {
	endpoint  string
	token     string
	namespace string

	client http.Client
}

func New(endpoint string, token string, namespace string) *OpenWhisk {
	ow := OpenWhisk{
		endpoint, //https://openwhisk.ng.bluemix.net/api/v1",
		token,
		namespace,
		http.Client{},
	}
	return &ow
}

func (ow *OpenWhisk) TriggerAction(action string, args string) error {
	url := ow.endpoint + "/namespaces/" + ow.namespace + "/actions/" + action
	jsonStr := []byte(args)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}

	addHeader(req, ow.Token)
	client := ow.client
	err := doRequest(req, client)
	if err != nil {
		return err
	}

	return nil
}

func doRequest(req *http.Request, client http.Client) error {
	resp, err := client.Do(req)
	if err != nil {
		return error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", body)

	return nil
}

func addHeader(req *http.Request, token string) {
	auth := "Basic " + token
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", auth)
}
