/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package pkg

import (
	"fmt"
	"strings"
	"text/template"
	sprig "github.com/Masterminds/sprig/v3"
)

// Funcs return available template out of the box and custom functions.
func TemplateFuncs() template.FuncMap {
	return builtinFuncs()
}

func builtinFuncs() template.FuncMap { // nolint:gocyclo
	fns := template.FuncMap{
		"default": func(_default string, value string) string {
			if value != "" {
				return value
			}
			return _default
		},
		"indent": func(extra int, char string) string {
			return GenerateIndentation(2, extra, char)
		},
		"name": func(name string) string {
			return SanitizeName(name, true)
		},
		"sanitize": func(word string) string {
			return Sanitize(word)
		}, 
		"ternary": func(condition interface{}, trueValue string, falseValue string) string {
			var c bool
			switch x := fmt.Sprintf("%T", condition); x {
			case "string":
				c = condition.(string) != ""
			case "int":
				c = condition.(int) != 0
			case "bool":
				c = condition.(bool)
			}
			if c {
				return trueValue
			}
			return falseValue
		},
		"tostring": func(s interface{}) string {
			return fmt.Sprintf("%v", s)
		},

		// trim
		"trim": func(cut string, s string) string {
			if s != "" {
				return strings.Trim(s, cut)
			}
			return s
		},
		"trimLeft": func(cut string, s string) string {
			if s != "" {
				return strings.TrimLeft(s, cut)
			}
			return s
		},
		"trimRight": func(cut string, s string) string {
			if s != "" {
				return strings.TrimRight(s, cut)
			}
			return s
		},
		"trimPrefix": func(prefix string, s string) string {
			if s != "" {
				return strings.TrimPrefix(s, prefix)
			}
			return s
		},
		"trimSuffix": func(suffix string, s string) string {
			if s != "" {
				return strings.TrimSuffix(s, suffix)
			}
			return s
		},

		// anchors
		"anchorNameMarkdown": func(prefix string, value string) string {
			return CreateAnchorMarkdown(prefix, value, true, true)
		},
		"anchorNameAsciidoc": func(prefix string, value string) string {
			return CreateAnchorAsciidoc(prefix, value, true, true)
		},
	}

	for name, fn := range sprig.FuncMap() {
		if _, found := fns[name]; !found {
			fns[name] = fn
		}
	}

	return fns
}

func GenerateIndentation(base int, extra int, char string) string {
	if char == "" {
		return ""
	}
	if base < 1 || base > 5 {
		base = 2
	}
	var indent string
	for i := 0; i < base+extra; i++ {
		indent += char
	}
	return indent
}
