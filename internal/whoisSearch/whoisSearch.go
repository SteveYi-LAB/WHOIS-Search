package whoisSearch

import (
	"fmt"

	"github.com/likexian/whois"
)

// Choose IRR DB
func IRR_DB(IRR, target string) string {
	var return_date string
	var err error
	switch IRR {
	default:
		return_date, err = whois.Whois(target)
		checkerr(err)
	case "afrinic":
		return_date, err = whois.Whois(target, "whois.afrinic.net")
		checkerr(err)
	case "apnic":
		return_date, err = whois.Whois(target, "whois.apnic.net")
		checkerr(err)
	case "arin":
		return_date, err = whois.Whois(target, "whois.arin.net")
		checkerr(err)
	case "lacnic":
		return_date, err = whois.Whois(target, "whois.lacnic.net")
		checkerr(err)
	case "ripencc":
		return_date, err = whois.Whois(target, "whois.ripe.net")
		checkerr(err)
	case "RADB":
		return_date, err = whois.Whois(target, "whois.radb.net")
		checkerr(err)
	}
	return return_date
}

// Check error
func checkerr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
