package types

import (
	"fmt"
	"regexp"
)

var namespaceRegExp = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9]{2,31}$`)
var prefixRegExp = regexp.MustCompile(`^[a-zA-Z0-9]{2,31}$`)

func GetDenom(namespaceId string, prefix string, baseDenom string) string {
	return fmt.Sprintf("%s/%s/%s", namespaceId, prefix, baseDenom)
}
