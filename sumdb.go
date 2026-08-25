package git

import (
	"fmt"
	"net/http"
	"net/url"
)

// HTTPSumDB is the real SumDBClient, backed by sum.golang.org's lookup
// endpoint (https://go.dev/design/25530-sumdb#lookup).
type HTTPSumDB struct {
	Client *http.Client // nil = http.DefaultClient
}

func (s *HTTPSumDB) Lookup(modulePath, version string) (bool, error) {
	client := s.Client
	if client == nil {
		client = http.DefaultClient
	}
	u := fmt.Sprintf("https://sum.golang.org/lookup/%s@%s",
		url.PathEscape(modulePath), url.PathEscape(version))
	resp, err := client.Get(u)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return false, nil // never indexed — free to use
	}
	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("sumdb lookup %s@%s: unexpected status %d", modulePath, version, resp.StatusCode)
	}
	return true, nil
}
