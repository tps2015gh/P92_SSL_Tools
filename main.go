package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "ssl-check [domain1] [domain2] ...",
		Short: "Check SSL certificate expiration, start date, company, and issuer",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for _, domain := range args {
				checkSSL(domain)
			}
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func checkSSL(domain string) {
	// Remove https:// if present
	domain = strings.TrimPrefix(domain, "https://")
	domain = strings.TrimPrefix(domain, "http://")
	domain = strings.TrimSuffix(domain, "/")

	if !strings.Contains(domain, ":") {
		domain = domain + ":443"
	}

	conn, err := tls.DialWithDialer(&net.Dialer{
		Timeout: 10 * time.Second,
	}, "tcp", domain, &tls.Config{
		InsecureSkipVerify: false,
	})

	if err != nil {
		fmt.Printf("Error: Failed to connect to %s: %v\n", domain, err)
		return
	}
	defer conn.Close()

	certs := conn.ConnectionState().PeerCertificates
	if len(certs) == 0 {
		fmt.Printf("Error: No certificates found for %s.\n", domain)
		return
	}

	cert := certs[0] // Leaf certificate

	fmt.Printf("--------------------------------------------------\n")
	fmt.Printf(" Domain:      %s\n", domain)
	fmt.Printf(" Common Name: %s\n", cert.Subject.CommonName)
	fmt.Printf(" Company:     %s\n", strings.Join(cert.Subject.Organization, ", "))
	fmt.Printf(" Issuer (CA): %s\n", strings.Join(cert.Issuer.Organization, ", "))
	fmt.Printf(" Start Date:  %s\n", cert.NotBefore.Format("2006-01-02 15:04:05"))
	fmt.Printf(" Expiry Date: %s\n", cert.NotAfter.Format("2006-01-02 15:04:05"))

	daysLeft := int(time.Until(cert.NotAfter).Hours() / 24)
	fmt.Printf(" Days Left:   %d days\n", daysLeft)

	if daysLeft < 0 {
		fmt.Printf(" Status:      EXPIRED\n")
	} else if daysLeft < 30 {
		fmt.Printf(" Status:      EXPIRING SOON\n")
	} else {
		fmt.Printf(" Status:      VALID\n")
	}
	fmt.Printf("--------------------------------------------------\n")
}
