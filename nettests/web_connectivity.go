package nettests

import (
	"github.com/apex/log"
	"github.com/ooni/probe-cli/internal/database"
	engine "github.com/ooni/probe-engine"
)

func lookupURLs(ctl *Controller, limit int64, categories []string) ([]string, map[int64]int64, error) {
	var urls []string
	urlIDMap := make(map[int64]int64)
	testlist, err := ctl.Ctx.Session.QueryTestListsURLs(&engine.TestListsURLsConfig{
		Limit:      limit,
		Categories: categories,
	})
	if err != nil {
		return nil, nil, err
	}
	for idx, url := range testlist.Result {
		log.Debugf("Going over URL %d", idx)
		urlID, err := database.CreateOrUpdateURL(
			ctl.Ctx.DB, url.URL, url.CategoryCode, url.CountryCode,
		)
		if err != nil {
			log.Error("failed to add to the URL table")
			return nil, nil, err
		}
		log.Debugf("Mapped URL %s to idx %d and urlID %d", url.URL, idx, urlID)
		urlIDMap[int64(idx)] = urlID
		urls = append(urls, url.URL)
	}
	return urls, urlIDMap, nil
}

// WebConnectivity test implementation
type WebConnectivity struct {
}

// Run starts the test
func (n WebConnectivity) Run(ctl *Controller) error {
	log.Debugf("Enabled category codes are the following %v", ctl.Ctx.Config.Nettests.WebsitesEnabledCategoryCodes)
	urls, urlIDMap, err := lookupURLs(ctl, ctl.Ctx.Config.Nettests.WebsitesURLLimit, ctl.Ctx.Config.Nettests.WebsitesEnabledCategoryCodes)
	if err != nil {
		return err
	}
	ctl.SetInputIdxMap(urlIDMap)
	builder, err := ctl.Ctx.Session.NewExperimentBuilder(
		"web_connectivity",
	)
	if err != nil {
		return err
	}
	if err := builder.SetOptionString("LogLevel", "INFO"); err != nil {
		return err
	}
	return ctl.Run(builder, urls)
}

// WebConnectivityTestKeys for the test
type WebConnectivityTestKeys struct {
	Accessible bool   `json:"accessible"`
	Blocking   string `json:"blocking"`
	IsAnomaly  bool   `json:"-"`
}

// GetTestKeys generates a summary for a test run
func (n WebConnectivity) GetTestKeys(tk map[string]interface{}) (interface{}, error) {
	var (
		blocked    bool
		blocking   string
		accessible bool
	)

	// We need to do these complicated type assertions, because some of the fields
	// are "nullable" and/or can be of different types
	switch v := tk["blocking"].(type) {
	case bool:
		blocked = false
		blocking = "none"
	case string:
		blocked = true
		blocking = v
	default:
		blocked = false
		blocking = "none"
	}

	if tk["accessible"] == nil {
		accessible = false
	} else {
		accessible = tk["accessible"].(bool)
	}

	return WebConnectivityTestKeys{
		Accessible: accessible,
		Blocking:   blocking,
		IsAnomaly:  blocked,
	}, nil
}

// LogSummary writes the summary to the standard output
func (n WebConnectivity) LogSummary(s string) error {
	return nil
}
