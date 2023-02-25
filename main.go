package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter email address: ")
	scanner.Scan()
	email := scanner.Text()
	domain := strings.Split(email, "@")[1]
	checkEmail(domain, email)
}

func checkEmail(domain string, email string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string
	
	// Check for MX records
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	// Get all TXT records
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	// Check for SPF record
	for _, record := range txtRecords {
		if strings.Contains(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	// Check for DMARC record
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain) 
	if err != nil {
		log.Printf("Error: %v", err)
	}

	for _, record := range dmarcRecords {
		if strings.Contains(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	
	// Print results
	fmt.Println("Results for email: ", email)
	fmt.Println("Domain: ", domain)
	fmt.Println("Has MX: ", hasMX)
	fmt.Println("MX Records: ", mxRecords)
	fmt.Println("Has SPF: ", hasSPF)
	fmt.Println("SPF Record: ", spfRecord)
	fmt.Println("Has DMARC: ", hasDMARC)
	fmt.Println("DMARC Record: ", dmarcRecord)

	// Create opinionated report
	if hasMX && hasSPF && hasDMARC {
		fmt.Println("Email is valid")
	} else {
		fmt.Println("Email is invalid")
	}
}