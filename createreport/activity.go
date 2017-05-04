package createreport

import (
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
	reportID := "9106427CF8384AE9B2E5"
	url := fmt.Sprintf("https://www.concursolutions.com/api/v3.0/expense/reports/%s", reportID)
	return true, nil
}
