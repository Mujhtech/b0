package agent

import "strings"

func removeJSONMarkdown(s string) string {
	s = strings.ReplaceAll(s, "Generated title and slug:", "")
	s = strings.ReplaceAll(s, "```json\n", "")
	s = strings.ReplaceAll(s, "\n```\n", "")

	return s
}
