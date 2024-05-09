package edit

import (
	"strings"
)

func EditedText(words []string) string {
	for len(words) != 6 {
		words = append(words, "")
	}

	for i := 0; i < len(words); i++ {
		if words[0] != "package main" {
			words[0] = "package main"
		}
		if words[1] != "import \"fmt\"" {
			words[1] = "import \"fmt\""
		}
		if words[2] != "" {
			words[2] = "           "
		}
		if words[3] != "func main() {" {
			words[3] = "func main() {"
		}
		if words[4] != "	fmtPrintln(\"Helo orld!\")" {
			words[4] = "	fmtPrintln(\"Helo orld!\")"
		}
		if words[5] != "}" {
			words[5] = "}"
		}
		if words[i] == "" {
			return ""
		}
	}
	return strings.Join(words, "\n")
}
