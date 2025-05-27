package checker

import (
	"fmt"
	"net/http"
	"time"
)

type CheckResult struct {
	Target string
	Status string
	Err    error
}

func CheckURL(url string, results chan<- CheckResult) {
	client := http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := client.Get(url)

	if err != nil {
		results <- CheckResult{
			Target: url,
			Err:    fmt.Errorf("Request failed: %w, err"),
		}
		return
	}

	defer resp.Body.Close()
	results <- CheckResult{
		Target: url,
		Status: resp.Status,
	}
}

//func checkURL(url, string, result chan <- checkResult) {
//    resp, err := http.Get(url)
//    result <- checkResult{url, resp.Status, err}
//}
