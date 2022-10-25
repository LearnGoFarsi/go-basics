package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	if _, err := ReadDBPort(); err != nil {
		fmt.Println(err)
	}
}

// --> Error handling strategies

func error_propagation() (*http.Response, error) {

	var resp *http.Response

	err := func() error {
		url := "ftp://anyhost.com"
		var err error
		resp, err = http.Get(url)
		if err != nil {
			// construct a new error message that includes reasoning and the underlying error
			return fmt.Errorf("failed to call the %s: %v", url, err)
		}
		return nil
	}()

	if err != nil {
		// construct a new error message that includes reasoning and the underlying error
		return nil, fmt.Errorf("something went wrong: %v", err)
	}
	return resp, nil
}

func retry_on_error() (resp *http.Response, err error) {

	deadline := time.Now().Add(time.Second * 15)

	for tries := 1; time.Now().Before(deadline); tries++ {
		resp, err = http.Get("ftp://anyhost.com")
		if err != nil {
			// exponential back-off
			time.Sleep(time.Second << uint(tries))
			continue
		}

		return resp, nil
	}

	return nil, fmt.Errorf("failed to call the API: %v", err)
}

func log_and_exit() *http.Response {
	resp, err := http.Get("ftp://anyhost.com")
	if err != nil {
		log.Fatalf("failed to call the API %v. exit", err)
	}

	return resp
}

func log_and_continue() *http.Response {
	resp, err := http.Get("ftp://anyhost.com")
	if err != nil {
		fmt.Printf("failed to call the API: %v \n", err)
		return nil
	}

	return resp
}

// Error handling strategies <--

// --> panic and recover

func ReadDBPassOrDie() string {

	defer func() {
		fmt.Println("Call me before you kill the process.")
	}()

	pass, ok := os.LookupEnv("db_password")
	if !ok {
		panic("db_password is missing.")
	}
	return pass
}

func ReadDBPort() (port string, err error) {

	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("failed to read db_port %v \n", p)
		}
	}()

	port, ok := os.LookupEnv("db_port")
	if !ok {
		panic("db_port is missing.")
	}
	return port, nil
}

// panic and recover <--
