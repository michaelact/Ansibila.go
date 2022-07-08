package jinja

var (
	// Available symbols in jinja (within filters/tag)
	Symbols = []string{
		// 3-Char symbols
		"{{-", "-}}", "{%-", "-%}",

		// 2-Char symbols
		"==", ">=", "<=", "&&", "||", "{{", "}}", "{%", "%}", "!=", "<>",

		// 1-Char symbol
		"(", ")", "+", "-", "*", "<", ">", "/", "^", ",", ".", "!", "|", ":", "=", "%", "[", "]",
	}

	// Available keywords in jinja
	Keywords = []string{"for", "in", "and", "or", "not", "true", "false", "as", "export"}
)
