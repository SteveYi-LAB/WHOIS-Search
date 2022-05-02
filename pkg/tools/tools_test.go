package tools_test

import (
	"strconv"
	"strings"
	"testing"

	valid "github.com/asaskevich/govalidator"
)

// Check ASN
func IsASN(target string) bool {

	is_asn := strings.TrimPrefix(strings.ToUpper(target), "AS")
	asn, err := strconv.Atoi(is_asn)
	if err != nil {
		return false
	}
	if asn > 0 && asn < 4294967295 {
		return true
	} else {
		return false
	}
}

// Check target type and return it
func CheckType(target string) string {
	if valid.IsIP(target) {
		if valid.IsIPv4(target) {
			return "IPv4"
		}
		if valid.IsIPv6(target) {
			return "IPv6"
		}
	}
	if IsASN(target) {
		return "ASN"
	}
	if valid.IsDNSName(target) {
		return "Domain"
	}
	return "Not Found"
}

// Test CheckType
func TestMain(t *testing.T) {
	IsASN("AS9810238091283091830989")
}
