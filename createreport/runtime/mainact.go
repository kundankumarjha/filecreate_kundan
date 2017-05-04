package main

import (
	"fmt"
	"net/http"
)

// Eval implements activity.Activity.Eval
func main(){
	reportID := "9106427CF8384AE9B2E5"
	url := fmt.Sprintf("https://www.concursolutions.com/api/v3.0/expense/reports/%s", reportID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.Header.Set("Authorization", "OAuth 0_yUp5EggL5HKz8pXwbNllocNrM=")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	
	fmt.Println("Response = ", resp.Status)

	return
}
