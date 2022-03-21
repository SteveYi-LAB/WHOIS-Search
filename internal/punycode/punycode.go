package punycode

import (
	"log"

	"golang.org/x/net/idna"
)

var p *idna.Profile

// Converts a domain name to ASCII.
func ConvertertoASCII(domain string) string {
	p = idna.New()
	return_data, err := p.ToASCII(domain)
	if err != nil {
		log.Println(err)
	}
	return return_data
}
