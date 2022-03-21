package tools

import (
	"strconv"
	"strings"

	valid "github.com/asaskevich/govalidator"
)

// Check ASN
func IsASN(target string) bool {

	is_asn := strings.TrimPrefix(strings.ToUpper(target), "AS")
	_, err := strconv.Atoi(is_asn)

	return err == nil
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
