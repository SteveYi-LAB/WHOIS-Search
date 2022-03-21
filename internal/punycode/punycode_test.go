package punycode_test

import (
	"fmt"
	"testing"

	valid "github.com/asaskevich/govalidator"
	"golang.org/x/net/idna"
)

var p *idna.Profile

func TestConverter(t *testing.T) {
	p = idna.New()

	fmt.Println(p.ToASCII("中文.tw"))
	fmt.Println(p.ToUnicode("google.tw"))
	if valid.IsDNSName("中文.tw") {
		fmt.Println("ACK")
	} else {
		fmt.Println("Failed")
	}
}
