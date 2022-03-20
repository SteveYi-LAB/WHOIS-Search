package whoisSearch

import (
	"fmt"
	"testing"

	"github.com/likexian/whois"
)

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

func checkerr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func TestMain(t *testing.T) {
	fmt.Println(IRR_DB("", "steveyi.net"))
	fmt.Println(IRR_DB("ripencc", "as7480"))
	fmt.Println(IRR_DB("apnic", "as141173"))
}
