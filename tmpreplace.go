package tmpreplace

import (
	"fmt"
	"github.com/dongri/phonenumber"
	"strings"
)

func ValidateNumberE164(number string, mobileOnly bool) error {
	if strings.HasPrefix(number, "+") {
		number = strings.Replace(number, "+", "", 1)
	}
	// for validate 883140 number in Incognito project
	if strings.HasPrefix(number, "883140") && len(number) == 15 {
		return nil
	}
	r := phonenumber.GetISO3166ByNumber(number, !mobileOnly)
	if r.CountryName == "" {
		return fmt.Errorf("Not found")
	}
	return nil
}
