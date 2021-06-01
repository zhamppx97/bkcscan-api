package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://bkcscan.com/api?module=account&action=eth_get_balance&address=0x107F4dC628EE75EbcB50953f8a72EbddD280869C"
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
