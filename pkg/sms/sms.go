package sms

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"fmt"
)

const (

	fromPhone = "+12139479583"
	toPhone   = "+998901233323"

	accountSid = "AC1503e08dd56cf5209b3503b03e7c1d1c"
	authToken  = "626725ed325ae38a4e6170bf2b578da3"

)

type Sms struct {
	client http.Client
}

func New(client http.Client) *Sms {
	return &Sms{
		client: client,
	}
}
func (s Sms) SendToPhone(body string) error {
	// Set account keys & information
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	msgData := url.Values{}
	msgData.Set("To",toPhone)
	msgData.Set("From", fromPhone)
	msgData.Set("Body", body)
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data["sid"])
		}
		return err
	} else {
		fmt.Println(resp.Status)
	}
	return nil
}