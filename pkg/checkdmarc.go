package pkg

import (
	"net"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	hasDMARC    bool
	dmarcRecord string
)

func CheckDmarc(domain string) (bool, string) {
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		logrus.WithError(err).Errorf("Error: could not lookup DMARC for domain %s", domain)
	} else {
		for _, record := range dmarcRecords {
			if strings.HasPrefix(record, "v=DMARC1") {
				hasDMARC = true
				dmarcRecord = record
				break
			}
		}
	}
	return hasDMARC, dmarcRecord
}
