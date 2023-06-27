package main

import (
	"log"
	"net/http"

	"github.com/druvainc/Platform/reporting-api/gosdk/reports"
	"github.com/druvainc/Platform/reporting-api/gosdk/spec"
	"github.com/druvainc/gorestlib/restliberror"
)

func main() {
	// Replace with your credentials
	apiURL := "https://apis-us0.druva.com"
	clientID := "API credential Client ID"
	secretKey := "API credential Secret Key"

	reportsClient, err := reports.GetReportsClientFromCredentials(apiURL, clientID, secretKey)
	if err != nil {
		log.Printf("Error in GetReportsClientFromCredentials: %v", err)
		return
	}

	// Get Report List
	reportListAPIVersion := "v1"
	err = GetReportList(reportsClient, reportListAPIVersion, nil)
	if err != nil {
		log.Printf("Error in GetReportList: %v", err)
		return
	}

	// Get Report Data
	// Replace with required report ID
	reportID := "epLastBackupStatuss"
	reportVersion := "v1"
	pageSize := 500

	// Filters
	filters := spec.Filter{
		PageSize: &pageSize,
		FilterBy: []spec.FilterBy{
			{
				FieldName: "status",
				Value:     "Backup Failed",
				Operator:  "EQUAL",
			},
		},
	}

	err = GetReportData(reportsClient, apiURL, clientID, secretKey, reportID, reportVersion, filters)
	if err != nil {
		log.Printf("Error in GetReportData: %v", err)
		return
	}
}

func GetReportList(reportsClient *reports.ReportsClient, version string, queryParameters map[string]string) error {
	// GetReportList using reports client
	getReportListResponse, err := reportsClient.GetReportList(version, queryParameters)
	if err != nil {
		log.Printf("Error in getting report list: %v", err)
		return err
	}
	log.Printf("ReportList: %s", getReportListResponse)
	return nil
}

func GetReportData(reportsClient *reports.ReportsClient, apiURL string, clientID string, secretKey string, reportID string, version string, filter spec.Filter) error {
	// GetReportData using reports client
	getReportResponse, err := reportsClient.GetReport(reportID, version, "", filter)
	if err != nil {
		log.Printf("GetReportData: Error in getting report data: %v", err)
		return err
	}
	log.Printf("getReportResponse: %+v", getReportResponse)

	pageToken := getReportResponse.NextPageToken

	for len(pageToken) > 0 {
		getReportResponseWithPageToken, err := reportsClient.GetReport(reportID, version, pageToken, filter)
		if err != nil {
			if apiError, ok := err.(restliberror.RestLibError); ok {
				log.Printf("GetReportData: Error in getting report data with page token: %s, ErrorCode: %d", pageToken, apiError.Code)
				if apiError.Code != http.StatusForbidden {
					return err
				}
				log.Print("GetReportData: http.StatusForbidden, Retrying by re-authenticating and creating new reports client..", apiError.Code)
				reportsClient, err = reports.GetReportsClientFromCredentials(apiURL, clientID, secretKey)
				if err != nil {
					log.Printf("GetReportData: reports.GetReportsClientFromCredentials")
					return err
				}
				getReportResponseWithPageToken, err = reportsClient.GetReport(reportID, version, pageToken, filter)
				if err != nil {
					log.Printf("GetReportData: Error in getting report data with new reports client and page token: %+v", err)
					return err
				}
			} else {
				log.Printf("GetReportData: Error in getting report data with page token: %+v", err)
				return err
			}
		}
		log.Printf("PageToken: %s, getReportResponseWithPageToken: %+v", pageToken, getReportResponseWithPageToken)
		pageToken = getReportResponseWithPageToken.NextPageToken
	}

	return nil
}
