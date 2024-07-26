package pkg

import (
	"net"

	"github.com/sirupsen/logrus"
)

var (
	hasMX     bool
	mxRecords []string
)

func CheckMX(domain string) (bool, []string) {
	// Check MX records
	mxRecordsList, err := net.LookupMX(domain)
	if err != nil {
		logrus.WithError(err).Errorf("Error: could not lookup MX for domain %s", domain)
	} else if len(mxRecordsList) > 0 {
		hasMX = true
		for _, mx := range mxRecordsList {
			mxRecords = append(mxRecords, mx.Host)
		}
	}
	return hasMX, mxRecords
}
