/*
Sniperkit-Bot
- Status: analyzed
*/

package utils

import (
	"regexp"
	"strings"
)

var (
	cleanAlphaNum = regexp.MustCompile("[^[:alpha:]\\s\\d_]")
	spaceTrim     = regexp.MustCompile("\\s+")
	manyDash      = regexp.MustCompile("-+")
)

func MakeId(v string) string {
	v = cleanAlphaNum.ReplaceAllString(v, " ")
	v = spaceTrim.ReplaceAllString(v, "-")
	v = manyDash.ReplaceAllString(v, "-")
	v = strings.Trim(v, "-")
	v = strings.ToLower(v)
	return v
}
