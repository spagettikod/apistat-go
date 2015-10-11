package apistat

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	endpoint string = "https://sxfr45o9b3.execute-api.eu-west-1.amazonaws.com/prod"
)

type apiStatError struct {
	ErrorMessage string `json:"errorMessage"`
}

type Stat struct {
	HTTPMethod   string `json:"httpMethod"`
	URL          string `json:"url"`
	Status       int    `json:"status"`
	ResponseTime int64  `json:"responseTime"`
	BytesRead    int64  `json:"bytesRead"`
	BytesWritten int64  `json:"bytesWritten"`
	UserID       int64  `json:"userId"`
	APIKey       string `json:"apiKey"`
}

func Post(s Stat) error {
	var b []byte
	var err error
	b, err = json.Marshal(s)
	var resp *http.Response
	resp, err = http.Post(endpoint, "application/json", bytes.NewReader(b))
	if err != nil {
		return fmt.Errorf("Error while posting to ApiStat: %v\n", err)
	}
	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("Reading response body failed: %v", err)
		}
		e := apiStatError{}
		err = json.Unmarshal(body, &e)
		if err != nil {
			return fmt.Errorf("Unmarshal of ApiError failed: %v", err)
		}
		return errors.New(e.ErrorMessage)
	}
	return nil
}
