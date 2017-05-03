package filecreate

import (
	"os"

	"io"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// log is the default package logger
var log = logger.GetLogger("activity-createreport-file")

const (
	ivReportname = "reportName"
	ovCreated  = "isCreated"
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

	fileName := context.GetInput(ivReportname).(string)
	content := "Hello file activity"
	file, err := os.Create(fileName)
	defer file.Close()
	io.WriteString(file, content)

	if err != nil {
		context.SetOutput(ovCreated, false)
		return true, err
	}
	context.SetOutput(ovCreated, true)

	return true, nil
}
