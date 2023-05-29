package branch

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

const _replacementCharacter = "-"

var _regexp = regexp.MustCompile("[^A-Za-z0-9.]+")

// GenerateName generates a new branch name off of the identifier and title specified.
// Both the identifier and title must not be blank.
// Leading hyphens (prefix and suffix) are removed from the name returned.
func GenerateName(id string, title string, maxLen int) (string, error) {
	if id == "" || strings.TrimSpace(id) == "" {
		return "", errors.New("id must not be blank")
	}
	if title == "" || strings.TrimSpace(title) == "" {
		return "", errors.New("title must not be blank")
	}
	titleWithoutDotSuffix := strings.TrimSuffix(title, ".")
	if titleWithoutDotSuffix == "" {
		return "", fmt.Errorf("invalid title: %s", title)
	}
	generatedName := strings.ToLower(
		_regexp.ReplaceAllString(
			fmt.Sprintf("%s-%s", id, titleWithoutDotSuffix),
			_replacementCharacter))
	generatedName = strings.TrimSuffix(generatedName, _replacementCharacter)
	generatedName = strings.TrimPrefix(generatedName, _replacementCharacter)
	if maxLen <= 0 {
		return generatedName, nil
	}
	if utf8.RuneCountInString(generatedName) < maxLen {
		return generatedName, nil
	}
	return string([]rune(generatedName)[:maxLen]), nil
}
