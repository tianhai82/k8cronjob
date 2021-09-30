package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func send(url string, user string, pw string, jsonStr []byte, method string) (map[string]interface{}, error) {
	fmt.Println("URL:>", url)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	req.SetBasicAuth(user, pw)
	if method != "PATCH" {
		req.Header.Set("Content-Type", "application/json")
	} else {
		req.Header.Set("Content-Type", "application/merge-patch+json")
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var result map[string]interface{}
	json.Unmarshal(buf, &result)

	fmt.Println("response Status:", resp.Status)
	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("call failed")
	}
	return result, nil
}
