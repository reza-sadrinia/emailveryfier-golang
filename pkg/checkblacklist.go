package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

type BlacklistResponse struct {
	IsBlocked bool   `json:"is_blocked"`
	Provider  string `json:"provider"`
}

func CheckBlacklist(domain string) (bool, string) {
	apiKey := os.Getenv("API_KEY")
	apiURL := fmt.Sprintf("https://ipqualityscore.com/api/json/domain/%s/%s", apiKey, domain)

	resp, err := http.Get(apiURL)
	if err != nil {
		logrus.WithError(err).Errorf("Error: could not reach the blacklist API for domain %s", domain)
		return false, ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logrus.Errorf("Error: received non-OK response from the API for domain %s", domain)
		return false, ""
	}

	var response BlacklistResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		logrus.WithError(err).Error("Error: could not decode API response")
		return false, ""
	}

	return response.IsBlocked, response.Provider
}

// func CheckBlacklist(domain string) (bool, string) {
// 	blacklistProviders := []string{
// 		"zen.spamhaus.org",
// 		"bl.spamcop.net",
// 		"dnsbl.sorbs.net",
// 	}

// 	for _, provider := range blacklistProviders {
// 		lookupDomain := domain + "." + provider
// 		_, err := net.LookupIP(lookupDomain)
// 		if err == nil {
// 			return true, provider
// 		}
// 	}
// 	return false, ""
// }
