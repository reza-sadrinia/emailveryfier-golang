package pkg

import (
	"net"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	hasSPF    bool
	spfRecord string
)

func CheckSpf(domain string) (bool, string) {
	// Check TXT records for SPF
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		logrus.WithError(err).Errorf("Error: could not lookup TXT for domain %s", domain)
	} else {
		for _, record := range txtRecords {
			if strings.HasPrefix(record, "v=spf1") {
				hasSPF = true
				spfRecord = record
				break
			}
		}
	}
	return hasSPF, spfRecord
}
