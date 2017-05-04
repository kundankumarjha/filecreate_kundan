package filecreate

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = `{
  "name": "createreport-file",
  "type": "flogo:activity",
  "ref": "github.com/kundankumarjha/createfile_kundan/createreport",
  "version": "0.0.1",
  "title": "Test activity to create a report",
  "description": "Test activity to create a report",
  "homepage": "https://github.com/kundankumarjha/filecreate_kundan/tree/master/createreport",
  "inputs":[
  ],
  "outputs": [
    {
      "name": "status",
      "type": "string"
    }
  ]
}
`

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(NewActivity(md))
}
