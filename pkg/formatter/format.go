package formatter

import (
	"strings"
)

func NameFormatter(names ...string) string {
	return strings.Join(names, ", ")
}
