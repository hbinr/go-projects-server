package jstr

import "strings"

func CompatiblePrefixStr(url string, prefix ...string) string {
	p := "/"
	if len(prefix) > 0 {
		p = prefix[0]
	}
	if !strings.HasPrefix(url, p) {
		return p + url
	}
	return url
}
func CompatibleSuffixStr(url string, prefix ...string) string {
	p := "/"
	if len(prefix) > 0 {
		p = prefix[0]
	}
	if !strings.HasSuffix(url, p) {
		return p + url
	}
	return url
}
