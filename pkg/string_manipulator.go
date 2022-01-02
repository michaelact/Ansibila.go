// Reference
// https://gist.github.com/julien-noblet/1d9e9a5f4a526337a4fd6f8d4fb03e5c

package pkg

import (
	"fmt"
    "strings"
)

func StringInSlice(word string, list []string) bool {
    for _, wordInSlice := range list {
        if strings.Contains(word, wordInSlice) {
            return true
        }
    }

    return false
}

func TypeOf(v interface{}) string {
    return fmt.Sprintf("%T", v)
}
