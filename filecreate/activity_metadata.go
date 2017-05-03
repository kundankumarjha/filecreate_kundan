package filecreate

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = `{
  "name": "pankaj-file",
  "type": "flogo:activity",
  "ref": "github.com/kundankumarjha/createfile_kundan/filecreate",
  "version": "0.0.1",
  "title": "Test activity to create a file",
  "description": "Test activity to create a file",
  "homepage": "https://github.com/kundankumarjha/filecreate_kundan/tree/master/filecreate",
  "inputs":[
    {
      "name": "fileName",
      "type": "string",
      "required": true
    }
  ],
  "outputs": [
    {
      "name": "isCreated",
      "type": "boolean"
    }
  ]
}
`

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(NewActivity(md))
}
