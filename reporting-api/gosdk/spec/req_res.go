package spec

type GetReportListResponse struct {
	Reports []interface{} `json:"reports"`
}

type GetReportRequest struct {
	// The token to access the next page of results. Use the token value received in the previous response's parameter 'nextPageToken'.
	PageToken string `json:"pageToken"`

	// Filters to be applied
	Filters Filter `json:"filters"`
}

// FilterBy ...
type FilterBy struct {
	// Name of the filter
	FieldName string `json:"fieldName"`

	// value of the filter
	Value interface{} `json:"value"`

	// The operators that are supported for the filter are :
	// EQUAL
	// NOTEQUAL
	// CONTAINS
	// LT (less-than)
	// LTEQ (less-than-or-equal)
	// GT (greater-than)
	// GTEQ (greater-than-or-equal)
	Operator string `json:"operator"`
}

// Filter ...
type Filter struct {
	// Specify the maximum number of records that should returned in a response. Maximum limit is 500 records. Default pageSize is 100.
	PageSize *int `json:"pageSize"`

	// Filters for the report
	FilterBy []FilterBy `json:"filterBy"`
}

type GetReportResponse struct {
	// Data in report
	Data []map[string]any `json:"data"`

	// filters applied on the report
	Filters Filter `json:"filters"`

	// Time stamp until which report data will be returned in API
	LastSyncTimestamp string `json:"lastSyncTimestamp"`

	// The token to access the next page of results. This parameter will be empty for the last page of the results. This token is valid for 5 minutes.
	NextPageToken string `json:"nextPageToken"`
}
