package branch

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const _replacementCharacter = "-"

var _regexp = regexp.MustCompile("[^A-Za-z0-9.]+")

func GenerateName(id string, title string) (string, error) {
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
	return strings.TrimPrefix(strings.TrimSuffix(generatedName, _replacementCharacter), _replacementCharacter), nil
}
