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
	apiToken := "a181147c613dcd3e82908899084e5d39214e4a345eebae54"
	exotelSid := "ekagga2"
	apiKey := "741cbdd276bd1d0d81b3e3694c3205bbbf9d93794f96bb21"

	baseURL := fmt.Sprintf("https://%s:%s@api.exotel.com/v1/accounts/%s/Calls/connect.json", apiKey, apiToken, exotelSid)

	callData := url.Values{}
	callData.Set("To", "09686419568")
	callData.Set("From", "09686419568")
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
