package basic

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func HTTPClientExmaple() {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	// Get response
	// resp, err := c.Get("https://www.google.com/")
	// if err != nil {
	// 	fmt.Printf("Error: %s", err)
	// 	return
	// }
	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// fmt.Printf("Body: %s", body)
	googleRequest(c)
}

func googleRequest(c http.Client) {
	req, err := http.NewRequest("GET", "https://www.google.com/", nil)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	req.Header.Add("Accept", "application/json")
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Body: %s", body)
}
