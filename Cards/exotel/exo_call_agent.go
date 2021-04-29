package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	apiToken := ""
	exotelSid := ""
	apiKey := ""

	baseURL := fmt.Sprintf("https://%s:%s@api.exotel.com/v1/accounts/%s/Calls/connect.json", apiKey, apiToken, exotelSid)

	callData := url.Values{}
	callData.Set("To", "")
	callData.Set("From", "")
	callData.Set("CallerType", "trans")

	callDataReader := strings.NewReader(callData.Encode())
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, baseURL, callDataReader)
	req.SetBasicAuth(exotelSid, apiToken)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
