package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// SearchIssues queries GitHub issue tracker
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " ")) // We have to escape the query
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
