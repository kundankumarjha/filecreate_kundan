package createreport

var jsonMetadata = `{
  "name": "createreport-file",
  "version": "0.0.1",
  "description": "Test activity to create a report",
  "inputs":[
	{
      "name": "reportID",
      "type": "string",
	  "required": true
    }
  ],
  "outputs": [
    {
      "name": "status",
      "type": "string"
    }
  ]
}`
