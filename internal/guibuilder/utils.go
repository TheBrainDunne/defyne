package guibuilder

import "strings"

func encodeDoubleQuote(inStr string) (outStr string) {
	outStr = strings.ReplaceAll(inStr, "\"", "\\\"")
	return
}
