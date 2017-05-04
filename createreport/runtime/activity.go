package createreport

import (
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

type fileActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new fileActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &fileActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *fileActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *fileActivity) Eval(context activity.Context) (done bool, err error) {
	url := "https://www.concursolutions.com/api/v3.0/expense/reports/9106427CF8384AE9B2E5"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return true, nil
	}
	req.Header.Set("Authorization", "OAuth 0_yUp5EggL5HKz8pXwbNllocNrM=")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return true, nil
	}
	defer resp.Body.Close()
	fmt.Println("Response = ", resp.Status)
	return true, nil
}
