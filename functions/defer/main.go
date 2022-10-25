package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Rate struct {
	Code, Name string
	Rate       float32
}

func main() {
	rate, _ := BTCUSD()
	fmt.Printf("Today's BTC price is %0.2f USD \n", rate.Rate)

	order()

	square(6)
}

func BTCUSD() (*Rate, error) {
	resp, err := http.Get("https://bitpay.com/api/rates")
	if err != nil {
		return nil, fmt.Errorf("Failed to call the API")
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("Failed to read response")
	}

	rates := []Rate{}
	if err := json.Unmarshal(body, &rates); err != nil {
		return nil, fmt.Errorf("Failed to read the rates. %v", err)
	}

	for _, r := range rates {
		if r.Code == "USD" {
			return &r, nil
		}
	}

	return nil, nil
}

func order() {
	defer func() {
		fmt.Println("1")
	}()

	defer func() {
		fmt.Println("2")
	}()

	defer func() {
		fmt.Println("3")
	}()
}

func square(x int) (r int) {

	defer func() {
		fmt.Printf("Square of %d is %d \n", x, r)
	}()

	r = x * x
	return
}
