package stringUtils

import (
	"fmt"
)

func AppendPrepend(str string, suffix string, prefix string) string {
	return fmt.Sprintf("%s%s%s", suffix, str, prefix)
}

func CreateString(template string, values ...interface{}) string {
	return fmt.Sprintf(template, values...)
}
