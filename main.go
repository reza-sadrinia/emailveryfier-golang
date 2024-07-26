package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/reza-sadrinia/email-veryfier/pkg"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.WithError(err).Fatalln("Error: could not load.env file")
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your domain: ")

	if scanner.Scan() {
		domain := scanner.Text()
		hasMX, mxRecords := pkg.CheckMX(domain)
		hasSPF, spfRecord := pkg.CheckSpf(domain)
		hasDMARC, dmarcRecord := pkg.CheckDmarc(domain)
		hasBlacklist, blacklistProvider := pkg.CheckBlacklist(domain)

		fmt.Printf("\nDomain: %s\n", domain)
		fmt.Printf("MX Status: %v\n", hasMX)
		fmt.Printf("MX Records: %v\n", mxRecords)
		fmt.Printf("SPF Status: %v\n", hasSPF)
		fmt.Printf("SPF Record: %s\n", spfRecord)
		fmt.Printf("DMARC Status: %v\n", hasDMARC)
		fmt.Printf("DMARC Record: %s\n", dmarcRecord)
		fmt.Printf("Blacklist Status: %v\n", hasBlacklist)
		if hasBlacklist {
			fmt.Printf("Blacklist Provider: %s\n", blacklistProvider)
		}
	}

	if err := scanner.Err(); err != nil {
		logrus.WithError(err).Fatal("Error: could not read from input")
	}
}
