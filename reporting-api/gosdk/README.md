# Druva Cloud Platform Reports Go SDK

Before you begin, you must have the API Credentials, which is a combination of Client ID and Secret Key, for the application or tool which you intend to integrate with the Druva products. API Credentials can be created from the [Druva Cloud Platform Console](https://login.druva.com/).

If you do not have the API Credentials, you can request your Druva Cloud administrator to provide you the API Credentials. Refer to [Integration Workflow](https://developer.druva.com/docs#section-integration-workflow) for the steps.

**Reference**: [Developers Portal](https://developer.druva.com/reference)
## Configuring the Reports SDK for Go

### Install below go packages and SDK

```
go get golang.org/x/oauth2
go get github.com/druvainc/gorestlib
go get github.com/druvainc/Platform/reporting-api/gosdk
```

### Create reports client by providing your API credentials. 
Replace apiURL, clientID and secretKey with your parameters
```
apiURL := "https://apis-us0.druva.com"
clientID := "API credential Client ID"
secretKey := "API credential Secret Key"

reportsClient, err := reports.GetReportsClientFromCredentials(apiURL, clientID, secretKey)
if err != nil {
	log.Printf("Error in GetReportsClientFromCredentials: %v", err)
	return
}
```

### Get Report List using reports client

Get list of Reports by providing 
1. version: The version supported for the report list API. If empty string provided, default version will be "v1"
2. queryParameters: Additional API URL query parameters if any

Reference: [Developers Portal](https://developer.druva.com/reference)
```
getReportListResponse, err := reportsClient.GetReportList(version, queryParameters)
if err != nil {
	log.Printf("Error in getting report list: %v", err)
	return err
}
log.Printf("getReportListResponse: %v", getReportListResponse)
```

### Get Report data using reports client
Get report data by providing
1. reportID: The unique ID given for the report
2. version: The version supported for the report. Example - "v1". If empty string provided, default version will be "v1"
3. pageToken: Should be empty string for first call. The token to access the next page of results. Use the token value received in the previous response's parameter 'nextPageToken'
4. filter: filter to be applied

Reference: [Developers Portal](https://developer.druva.com/reference)
```
getReportResponse, err := reportsClient.GetReport(reportID, version, "", filter)
if err != nil {
	log.Printf("Error in getting report data: %v", err)
	return err
}
log.Printf("getReportResponse: %+v", getReportResponse)

pageToken := getReportResponse.NextPageToken

for len(pageToken) > 0 {
	getReportResponseWithPageToken, err := reportsClient.GetReport(reportID, version, pageToken, filter)
	if err != nil {
		log.Printf("Error in getting report data with page token: %+v", err)
		return err
	}
	log.Printf("PageToken: %s, getReportResponseWithPageToken: %+v", pageToken, getReportResponseWithPageToken)
	pageToken = getReportResponseWithPageToken.NextPageToken
}
```
