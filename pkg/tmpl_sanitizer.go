/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package pkg

import (
	"strings"
)

// SanitizeName escapes underscore character which have special meaning in
// Markdown.
func SanitizeName(name string, escape bool) string {
	if escape {
		// Escape underscore
		name = strings.ReplaceAll(name, "_", "\\_")
	}
	return name
}

// Sanitize cleans a word in markdown document.
func Sanitize(word string) string {
	word = strings.ReplaceAll(word, `|`, `\|`)
	return word
}
