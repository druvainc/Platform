package spec

// URLs
const (
	APIHost = "https://apis.druva.com"
)

// URL Prefixes
const (
	ReportingPrefix = "/platform/reportsvc"
)

// Resources
const (
	// TokenResource ...
	TokenResource = "/token"
	// ListReportsResource
	ListReportsResource = "/%s/reports"
	// GetReportResource
	GetReportResource = "/%s/reports/%s"
)

// HTTP request parameters
const (
	// HeaderContentType ...
	HeaderContentType = "Content-Type"
	// HeaderContentTypeJSON ...
	HeaderContentTypeJSON = "application/json"
	// HeaderAuthorization
	HeaderAuthorization = "authorization"

	// AuthorizationBearer ...
	AuthorizationBearer = "Bearer"
)

// ProductIDs for reports
const (
	HybridWorkloadsProductID = 12289
	EndpointsProductID       = 8193
	CyberResilienceProductID = 36880
)
