package api

import "fmt"

func (c *Client) SpiderSite(url string) (error) {
	var err error
	var inScope bool
	//var scanProgress *ScanProgress
	var siteMap []SiteMapEntry

	inScope, err = c.IsInScope(url)
	if err != nil {
		return err
	}
	fmt.Printf("InScope: %v (%s)\n", inScope, url)

	if !inScope {
		fmt.Printf("Adding URL to scope: %s\n", url)
		err = c.TargetScopeAdd(url)
		if err != nil {
			return err
		}
	}

	fmt.Printf("Starting Spider: %s\n", url)
	err = c.Spider(url)
	if err != nil {
		return err
	}

	//// TODO: This doesn't work for spider status..
	// Ref: https://github.com/vmware/burp-rest-api/issues/35
	//scanProgress, err = c.ScannerStatus()
	//if err != nil {
	//	return err
	//}
	//fmt.Printf("ScanProgress: %v%%\n", scanProgress.ScanPercentage)

	fmt.Println("TODO: Wait till finished spidering somehow")

	siteMap, err = c.TargetSitemap(url)
	if err != nil {
		return err
	}
	fmt.Printf("Sitemap: %+v\n", siteMap)

	return nil
}

func (c *Client) IsInScope(url string) (bool, error) {
	scopeItem, err := c.TargetScopeCheck(url)
	if err != nil {
		return false, err
	}
	return scopeItem.InScope, nil
}

func (c *Client) ReportAsXML(urlPrefix string) (string, error) {
	return c.Report(urlPrefix, "XML")
}

func (c *Client) ReportAsHTML(urlPrefix string) (string, error) {
	return c.Report(urlPrefix, "HTML")
}
