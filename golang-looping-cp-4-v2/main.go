package main

import (
	"fmt"
	"strings"
)

func EmailInfo(email string) string {
	atIndex := strings.Index(email, "@")
	dotIndex := strings.Index(email, ".")
	domain := email[atIndex+1 : dotIndex]
	tld := email[dotIndex+1:]
	return fmt.Sprintf("Domain: %s dan TLD: %s", domain, tld)
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(EmailInfo("admin@yahoo.co.id"))
}
