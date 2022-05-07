package branch

import (
	"fmt"
	"regexp"
	"strings"
)

const _replacementCharacter = "-"

var _regexp = regexp.MustCompile("[^A-Za-z0-9.]+")

func GenerateName(index int, title string) string {
	generatedName := strings.ToLower(_regexp.ReplaceAllString(fmt.Sprintf("%d-%s", index, title), _replacementCharacter))
	if len(generatedName) > 2 && strings.HasSuffix(generatedName, _replacementCharacter) {
		generatedName = generatedName[0 : len(generatedName)-2]
	}
	return generatedName
}
