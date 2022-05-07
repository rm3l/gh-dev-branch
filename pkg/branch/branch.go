package branch

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const _replacementCharacter = "-"

var _regexp = regexp.MustCompile("[^A-Za-z0-9.]+")

func GenerateName(index int, title string) (string, error) {
	if index <= 0 {
		return "", fmt.Errorf("invalid negative or zero index: %d", index)
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
			fmt.Sprintf("%d-%s", index, titleWithoutDotSuffix),
			_replacementCharacter))
	return strings.TrimSuffix(generatedName, _replacementCharacter), nil
}
