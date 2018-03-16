package api

import (
	"fmt"

	"github.com/go-resty/resty"
)

type Client struct {
	restyClient *resty.Client
}

type Versions struct {
	BurpVersion string `json:"burpVersion"`
	RestVersion string `json:"extensionVersion"`
}

type ScopeItem struct {
	InScope bool   `json:"inScope"`
	Url     string `json:"url"`
}

type ScanProgress struct {
	ScanPercentage int `json:"scanPercentage"`
}

type ScannerIssues struct {
	Issues []ScannerIssue `json:"issues"`
}

type ScannerIssue struct {
	Confidence            string         `json:"confidence"`
	Host                  string         `json:"host"`
	HttpMessages          []SiteMapEntry `json:"httpMessages"`
	IssueBackground       string         `json:"issueBackground"`
	IssueDetail           string         `json:"issueDetail"`
	IssueName             string         `json:"issueName"`
	IssueType             int            `json:"issueType"`
	Port                  int            `json:"port"`
	Protocol              string         `json:"protocol"`
	RemediationBackground string         `json:"remediationBackground"`
	RemediationDetail     string         `json:"remediationDetail"`
	Severity              string         `json:"severity"`
	Url                   string         `json:"url"`
}

type HttpMessageList struct {
	Messages []SiteMapEntry `json:"messages"`
}

type SiteMapEntry struct {
	Comment    string `json:"comment"`
	Highlight  string `json:"highlight"`
	Host       string `json:"host"`
	Port       int    `json:"port"`
	Protocol   string `json:"protocol"`
	Request    string `json:"request"`
	Response   string `json:"response"`
	StatusCode int    `json:"statusCode"`
	URL        string `json:"url"`
}

func DefaultClient(baseUrl string) *Client {
	c := resty.DefaultClient
	c.SetHostURL(baseUrl)

	return &Client{restyClient: c}
}

// BurpVersion retrieves the version of Burp and the version of the burp-rest-api extension
//   eg. GET http://127.0.0.1:8090/burp/versions
func (c *Client) BurpVersion() (*Versions, error) {
	resp, err := c.restyClient.R().
		SetResult(Versions{}).
		Get("/burp/versions")
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode() {
	case 200:
		return resp.Result().(*Versions), nil
	default:
		return nil, errorResponseFormatter(resp)
	}
}

// TODO: Get Burp Config
// TODO: Put Burp Config

// TargetScopeAdd will add the supplier url prefix to the target scope.
//   eg. PUT http://127.0.0.1:8090/burp/target/scope?url=http%3A%2F%2Fexample.com
func (c *Client) TargetScopeAdd(url string) error {
	resp, err := c.restyClient.R().
		SetQueryParam("url", url).
		Put("/burp/target/scope")
	if err != nil {
		return err
	}

	switch resp.StatusCode() {
	case 200:
		return nil
	default:
		return errorResponseFormatter(resp)
	}
}

// TargetScopeCheck will query whether a specific URL is within the current Suite-wide scope.
//   eg. GET http://127.0.0.1:8090/burp/target/scope?url=http%3A%2F%2Fexample.com
func (c *Client) TargetScopeCheck(url string) (*ScopeItem, error) {
	resp, err := c.restyClient.R().
		SetQueryParam("url", url).
		SetResult(ScopeItem{}).
		Get("/burp/target/scope")
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode() {
	case 200:
		return resp.Result().(*ScopeItem), nil
	default:
		return nil, errorResponseFormatter(resp)
	}
}

// TargetScopeExclude excludes the specified URL from the Suite-wide scope.
//   eg. DELETE http://127.0.0.1:8090/burp/target/scope?url=http%3A%2F%2Fexample.com
func (c *Client) TargetScopeExclude(url string) error {
	resp, err := c.restyClient.R().
		SetQueryParam("url", url).
		Delete("/burp/target/scope")
	if err != nil {
		return err
	}

	switch resp.StatusCode() {
	case 200:
		return nil
	default:
		return errorResponseFormatter(resp)
	}
}

// Spider sends a seed URL to the Burp Spider tool. The baseUrl should be in Suite-wide scope for the Spider to run.
//   eg. POST http://127.0.0.1:8090/burp/spider?baseUrl=http%3A%2F%2Fexample.com%2F
func (c *Client) Spider(baseUrl string) error {
	resp, err := c.restyClient.R().
		SetQueryParam("baseUrl", baseUrl).
		Post("/burp/spider")
	if err != nil {
		return err
	}

	switch resp.StatusCode() {
	case 200:
		return nil
	default:
		return errorResponseFormatter(resp)
	}
}

// ScannerActiveScan scans through Burp Sitemap and sends all HTTP requests with url starting with baseUrl to Burp Scanner for active scan.
//   eg. POST http://127.0.0.1:8090/burp/scanner/scans/active?baseUrl=http%3A%2F%2Fexample.com
func (c *Client) ScannerActiveScan(baseUrl string) error {
	resp, err := c.restyClient.R().
		SetQueryParam("baseUrl", baseUrl).
		Post("/burp/scanner/scans/active")
	if err != nil {
		return err
	}

	switch resp.StatusCode() {
	case 200:
		return nil
	default:
		return errorResponseFormatter(resp)
	}
}

// ScannerClearQueue deletes the scan queue map from memory, not from Burp suite UI.
//   eg. DELETE https://127.0.0.1:8090/burp/scanner/scans/active
func (c *Client) ScannerClearQueue() error {
	resp, err := c.restyClient.R().
		Delete("/burp/scanner/scans/active")
	if err != nil {
		return err
	}

	switch resp.StatusCode() {
	case 200:
		return nil
	default:
		return errorResponseFormatter(resp)
	}
}

// ScannerStatusPercent returns an aggregate of percentage completed for all the scan queue items.
//   eg. GET http://127.0.0.1:8090/burp/scanner/status
func (c *Client) ScannerStatusPercent() (int, error) {
	resp, err := c.restyClient.R().
		SetResult(ScanProgress{}).
		Get("/burp/scanner/status")
	if err != nil {
		return 0, err
	}

	switch resp.StatusCode() {
	case 200:
		return resp.Result().(*ScanProgress).ScanPercentage, nil
	default:
		return 0, errorResponseFormatter(resp)
	}
}

// ScannerIssues returns all of the current scan issues for URLs matching the specified urlPrefix. Performs a simple case-sensitive text match, returning all scan issues whose URL begins with the given urlPrefix. Returns all issues if urlPrefix is null.
//   eg. GET http://127.0.0.1:8090/burp/scanner/issues?urlPrefix=http%3A%2F%2Fexample.com
func (c *Client) ScannerIssues(urlPrefix string) ([]ScannerIssue, error) {
	resp, err := c.restyClient.R().
		SetQueryParam("urlPrefix", urlPrefix).
		SetResult(ScannerIssues{}).
		Get("/burp/scanner/issues")
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode() {
	case 200:
		return resp.Result().(*ScannerIssues).Issues, nil
	default:
		return nil, errorResponseFormatter(resp)
	}
}

// Report returns the scan report with current Scanner issues for URLs matching the specified urlPrefix in the form of a byte array. Report format can be specified as HTML or XML. Report with scan issues of all URLs are returned in HTML format if no urlPrefix and format are specified.
//   eg. GET http://127.0.0.1:8090/burp/report?urlPrefix=http%3A%2F%2Fexample.com&reportType=XML
func (c *Client) Report(urlPrefix string, reportType string) (string, error) {
	resp, err := c.restyClient.R().
		SetQueryParam("urlPrefix", urlPrefix).
		SetQueryParam("reportType", reportType).
		Get("/burp/report")
	if err != nil {
		return "", err
	}

	switch resp.StatusCode() {
	case 200:
		return string(resp.Body()), nil
	default:
		return "", errorResponseFormatter(resp)
	}
}

// TargetSitemap returns details of items in the Burp suite Site map. urlPrefix parameter can be used to specify a URL prefix, in order to extract a specific subset of the site map.
//   eg. GET http://127.0.0.1:8090/burp/target/sitemap?urlPrefix=http%3A%2F%2Fexample.com
func (c *Client) TargetSitemap(urlPrefix string) ([]SiteMapEntry, error) {
	resp, err := c.restyClient.R().
		SetQueryParam("urlPrefix", urlPrefix).
		SetResult(HttpMessageList{}).
		Get("/burp/target/sitemap")
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode() {
	case 200:
		return resp.Result().(*HttpMessageList).Messages, nil
	default:
		return nil, errorResponseFormatter(resp)
	}
}

// ProxyHistory returns details of items in Burp Suite Proxy history.
//   eg. GET http://127.0.0.1:8090/burp/proxy/history
func (c *Client) ProxyHistory() ([]SiteMapEntry, error) {
	resp, err := c.restyClient.R().
		SetResult(HttpMessageList{}).
		Get("/burp/proxy/history")
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode() {
	case 200:
		return resp.Result().(*HttpMessageList).Messages, nil
	default:
		return nil, errorResponseFormatter(resp)
	}
}

func errorResponseFormatter(resp *resty.Response) error {
	return fmt.Errorf("api returned an error response (%v): %s", resp.StatusCode(), string(resp.Body()))
}
