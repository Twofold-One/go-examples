package basic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func HTTPClientGithubAPI() {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	// zenRequest(c)
	// userReq(c)
	// postGistReq(c)
	// updateGistReq(c)
	deleteGistReq(c)

}

func zenRequest(c http.Client) {
	req, err := http.NewRequest("GET", "https://api.github.com/zen", nil)
	if err != nil {
		fmt.Printf("ERROR: %s", err)
		return
	}

	req.Header.Add("Accept", `application/json`)
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("ERROR: %s", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ERROR: %s", err)
		return
	}

	fmt.Printf("Body: %s\n", body)
	fmt.Printf("Response status: %s\n", resp.Status)
}

func userReq(c http.Client) {
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		fmt.Printf("ERROR: %s", err)
		return
	}
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Authorization", fmt.Sprintf("token %s", os.Getenv("TOKEN")))
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("ERROR: %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ERROR: %s", err)
		return
	}
	fmt.Printf("Body: %s\n", body)
	fmt.Printf("Response Status: %s\n", resp.Status)

}

type GistRequest struct {
	Files       map[string]File `json:"files"`
	Description string
	Public      bool
}

type File struct {
	Content string `json:"content"`
}

func postGistReq(c http.Client) {
	files := map[string]File{
		"main.go": File{"test"},
	}
	gistReq := GistRequest{
		Files:       files,
		Description: "this is a test",
		Public:      false,
	}
	gistReqJSON, err := json.Marshal(gistReq)
	if err != nil {
		logError(err)
	}

	req, err := http.NewRequest("POST", "https://api.github.com/gists", bytes.NewBuffer(gistReqJSON))
	if err != nil {
		logError(err)
	}
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Authorization", fmt.Sprintf("token %s", os.Getenv("TOKEN")))
	resp, err := c.Do(req)
	if err != nil {
		logError(err)
	}
	defer resp.Body.Close()
	fmt.Printf("Response status: %s\n", resp.Status)
}

func updateGistReq(c http.Client) {
	files := map[string]File{
		"main.go": File{"test updated"},
	}
	gistReq := GistRequest{
		Files:       files,
		Description: "this is a test",
		Public:      false,
	}
	gistRequestJSON, err := json.Marshal(gistReq)
	if err != nil {
		logError(err)
	}

	req, err := http.NewRequest("PATCH", "https://api.github.com/gists/a16b15aa29e459caecd59c907fda727c", bytes.NewBuffer(gistRequestJSON))
	if err != nil {
		logError(err)
	}
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Authorization", fmt.Sprintf("token %s", os.Getenv("TOKEN")))
	resp, err := c.Do(req)
	if err != nil {
		logError(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logError(err)
	}
	fmt.Printf("Body: %s\n", body)
	fmt.Printf("Response status: %s\n", resp.Status)
}

func logError(err error) {
	fmt.Printf("ERROR: %s\n", err)
	return
}

func deleteGistReq(c http.Client) {
	req, err := http.NewRequest("DELETE", "https://api.github.com/gists/a16b15aa29e459caecd59c907fda727c", nil)
	if err != nil {
		logError(err)
	}
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Authorization", fmt.Sprintf("token %s", os.Getenv("TOKEN")))
	resp, err := c.Do(req)
	if err != nil {
		logError(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logError(err)
	}
	fmt.Printf("Body: %s\n", body)
	fmt.Printf("Response status: %s\n", resp.Status)
}

// func errorChecker(err error) {
// 	if err != nil {
// 		fmt.Printf("ERROR: %s", err)
// 	}
// 	return
// }
