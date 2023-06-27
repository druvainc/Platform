package reports

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/druvainc/Platform/reporting-api/gosdk/spec"
	"github.com/druvainc/gorestlib"
	"github.com/druvainc/gorestlib/restliberror"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type ReportsClient struct {
	AccessToken string
	RestClient  gorestlib.RestClientInterface
}

// NewLogger returns new logger with UTC timezone and package prefix
func NewLogger() *log.Logger {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.LUTC)
	logger.SetPrefix("reportspkg ")
	return logger
}

// NewReportsClient returns new instance of reporting client
func NewReportsClient(accessToken string, apiURL string) (*ReportsClient, error) {
	logger := NewLogger()
	if len(apiURL) == 0 {
		logger.Printf("NewReportsClient: Empty apiURL, using default URL: %s", spec.APIHost)
		apiURL = spec.APIHost
	}
	if len(accessToken) == 0 {
		logger.Printf("NewReportsClient: Failed to initialize reporting client, Empty Access Token")
		return nil, errors.New("NewReportsClient: Empty Access Token")
	}
	return &ReportsClient{
		AccessToken: accessToken,
		RestClient:  gorestlib.NewRestClient(apiURL),
	}, nil
}

// GetReportList returns the list of accessible reports for provided list of productIDs.
//
// # Inputs
//
// version: The version supported for the report list API. If empty string provided, default version will be "v1"
//
// queryParameters: Additional API URL query parameters if any
//
// Refer: https://developer.druva.com/reference
func (client *ReportsClient) GetReportList(version string, queryParameters map[string]string) (*spec.GetReportListResponse, error) {
	logger := NewLogger()
	if len(version) == 0 {
		version = "v1"
	}
	response := spec.GetReportListResponse{}
	getReportListPath := spec.ReportingPrefix + fmt.Sprintf(spec.ListReportsResource, version)

	headers := map[string]string{
		spec.HeaderAuthorization: GetBearerTokenHeader(client.AccessToken),
		spec.HeaderContentType:   spec.HeaderContentTypeJSON,
	}
	err := client.RestClient.Get(getReportListPath, &response, queryParameters, headers)
	if err != nil {
		if apiError, ok := err.(restliberror.RestLibError); ok {
			if apiError.Code == http.StatusTooManyRequests {
				logger.Printf("GetReportList: client.RestClient.Post: http.StatusTooManyRequests encountered, Retrying....")
				rand.Seed(time.Now().UnixNano())
				delay := 0
				retryCount := 1
				for retryCount <= 3 {
					delay = delay + rand.Intn(30) + 1
					logger.Printf("GetReportList: Sleeping for %v seconds", delay)
					time.Sleep(time.Duration(delay * int(time.Second)))
					err := client.RestClient.Get(getReportListPath, &response, queryParameters, headers)
					if err != nil {
						if apiError, ok = err.(restliberror.RestLibError); ok {
							if apiError.Code != http.StatusTooManyRequests {
								logger.Printf("GetReportList: Error in client.RestClient.Post: %v", err)
								return nil, err
							}
							logger.Printf("Retry: %d, GetReportList: client.RestClient.Post: http.StatusTooManyRequests encountered: %v", retryCount, err)
							delay = 30 * retryCount
							retryCount++
							continue
						} else {
							logger.Printf("GetReportList: Error in client.RestClient.Post: %v", err)
							return nil, err
						}
					}
					break
				}
			}
		} else {
			logger.Printf("GetReport: Error in client.RestClient.Post: %v", err)
			return nil, err
		}
	}
	return &response, nil
}

// GetReport gets the data for mentioned report
//
// # Inputs
//
// reportID : The unique ID given for the report.
//
// version : The version supported for the report. Example - "v1". If empty string provided, default version will be "v1"
//
// pageToken: Should be empty string for first call. The token to access the next page of results. Use the token value received in the previous response's parameter 'nextPageToken'.
//
// filter: filter to be applied
//
// Refer: https://developer.druva.com/reference
func (client *ReportsClient) GetReport(reportID string, version string, pageToken string, filter spec.Filter) (*spec.GetReportResponse, error) {
	logger := NewLogger()
	if len(reportID) == 0 {
		logger.Printf("GetReport: Validation Error, Empty reportUniqueID")
		return nil, errors.New("Validation Error, Invalid input reportUniqueID")
	}
	if len(version) == 0 {
		version = "v1"
	}

	response := spec.GetReportResponse{}
	request := spec.GetReportRequest{
		Filters: filter,
	}
	if len(pageToken) != 0 {
		request.PageToken = pageToken
	}

	headers := map[string]string{
		spec.HeaderAuthorization: GetBearerTokenHeader(client.AccessToken),
		spec.HeaderContentType:   spec.HeaderContentTypeJSON,
	}

	getReportResourcePath := spec.ReportingPrefix + fmt.Sprintf(spec.GetReportResource, version, reportID)

	err := client.RestClient.Post(getReportResourcePath, request, &response, headers)
	if err != nil {
		if apiError, ok := err.(restliberror.RestLibError); ok {
			if apiError.Code == http.StatusTooManyRequests {
				logger.Printf("GetReport: client.RestClient.Post: http.StatusTooManyRequests encountered, Retrying....")
				rand.Seed(time.Now().UnixNano())
				delay := 0
				retryCount := 1
				for retryCount <= 3 {
					delay = delay + rand.Intn(30) + 1
					logger.Printf("GetReport: Sleeping for %v seconds", delay)
					time.Sleep(time.Duration(delay * int(time.Second)))
					err := client.RestClient.Post(getReportResourcePath, request, &response, headers)
					if err != nil {
						if apiError, ok = err.(restliberror.RestLibError); ok {
							if apiError.Code != http.StatusTooManyRequests {
								logger.Printf("GetReport: Error in client.RestClient.Post: %v", err)
								return nil, err
							}
							logger.Printf("Retry: %d, GetReport: client.RestClient.Post: http.StatusTooManyRequests encountered: %v", retryCount, err)
							delay = 30 * retryCount
							retryCount++
							continue
						} else {
							logger.Printf("GetReport: Error in client.RestClient.Post: %v", err)
							return nil, err
						}
					}
					break
				}
			}
		} else {
			logger.Printf("GetReport: Error in client.RestClient.Post: %v", err)
			return nil, err
		}
	}
	return &response, nil
}

// GetReportsClientFromCredentials authenticates client credentials provided with oauth2 and returns reports client
//
// # Inputs
//
// clientID: API credential Client ID
//
// secretKey: API credential Secret Key
//
// Refer: https://developer.druva.com/reference
func GetReportsClientFromCredentials(apiURL string, clientID string, secretKey string) (*ReportsClient, error) {
	logger := NewLogger()
	if len(apiURL) == 0 {
		logger.Printf("GetReportsClientFromCredentials: Empty apiURL, using default URL: %s", spec.APIHost)
		apiURL = spec.APIHost
	}
	// Create a oauth2 client to Authenticate and get access token
	oauthClient := clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: secretKey,
		TokenURL:     apiURL + spec.TokenResource,
		Scopes:       []string{"read"},
		AuthStyle:    oauth2.AuthStyleAutoDetect,
	}

	ctx := context.Background()
	token, err := oauthClient.Token(ctx)
	if err != nil {
		logger.Printf("GetReportsClientFromCredentials: Error in oauthClient.Token: %+v", err)
		return nil, err
	}
	logger.Print("Client credentials authenticated successfully...")

	// Create a reporting client with access token received after authentication
	reportsClient, err := NewReportsClient(token.AccessToken, apiURL)
	if err != nil {
		logger.Printf("GetReportsClientFromCredentials: NewReportsClient, Reports client initialization failed with Error: %v", err)
		return nil, err
	}
	logger.Print("NewReportsClient creation successful...")

	return reportsClient, nil
}

// GetBearerTokenHeader returns access token by appending Bearer keyword
func GetBearerTokenHeader(accessToken string) string {
	return spec.AuthorizationBearer + " " + accessToken
}
