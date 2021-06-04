package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	MODULE := "account"
	ACTION := "txlistinternal"
	ADDRESS := "0x107F4dC628EE75EbcB50953f8a72EbddD280869C"

	url := "https://bkcscan.com/api?module=" + MODULE + "&action=" + ACTION + "&address=" + ADDRESS
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
