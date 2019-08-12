package scraper

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var nameRegexp, _ = regexp.Compile("^\\s*(.+)\\s*\\((\\d+)\\)\\s*$")

// ExtractGameName split the name and the release date from a game name.
func ExtractGameName(s string) (string, int) {

	results := nameRegexp.FindStringSubmatch(strings.TrimSpace(s))

	date, parseError := strconv.Atoi(results[2])

	if parseError != nil {
		fmt.Println("Error: impossible to parse date of", s, "got error: ", parseError)
	}

	return strings.TrimSpace(results[1]), date
}
