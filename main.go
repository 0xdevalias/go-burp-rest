package main

import (
	"fmt"

	"github.com/0xdevalias/go-burp-rest-api/api"
)

func main() {
	// Note: The best kind of hacky testing is running a bunch of things in serial in main!! You've been warned..

	var err error
	var v *api.Versions

	url := "http://example.com"

	c := api.DefaultClient("http://127.0.0.1:8090")

	fmt.Printf("Removing URL from scope: %s\n", url)
	err = c.TargetScopeExclude(url)
	if err != nil {
		panic(err)
	}

	inScope, err := c.IsInScope(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("isInScope: %v\n", inScope)

	fmt.Printf("Adding URL to scope: %s\n", url)
	err = c.TargetScopeAdd(url)
	if err != nil {
		panic(err)
	}

	inScope2, err := c.IsInScope(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("isInScope: %v\n", inScope2)

	v, err = c.BurpVersion()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Burp Version: %+v\n", v)


	//err = c.SpiderSite(url)
	//if err != nil {
	//	panic(err)
	//}

	proxyHistory, err := c.ProxyHistory()
	if err != nil {
		panic(err)
	}
	fmt.Printf("ProxyHistoryCnt: %v\n", len(proxyHistory))

	fmt.Printf("ActiveScan: %v\n", url)
	err = c.ScannerActiveScan(url)
	if err != nil {
		panic(err)
	}

	issues, err := c.ScannerIssues(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("IssuesCnt: %v\n", len(issues))

	reportHtml, err := c.ReportAsHTML(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ReportHtmlCnt: %v\n", len(reportHtml))

	reportXml, err := c.ReportAsXML(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ReportXmlCnt: %v\n", len(reportXml))
}


