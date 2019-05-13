package updater

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

const (
	lastModTag         = "lastmod"
	frontMatterDividor = "---\n"
	timeFormat         = "2006-01-02T15:04:05-07:00"
)

var (
	replaceModRegex = regexp.MustCompile(fmt.Sprintf("%s:.*\n", lastModTag))
	addModRegex     = regexp.MustCompile("\\A\\s*---\n")
)

func updateLastMod(input []byte, lastMod time.Time) ([]byte, bool, error) {
	content := string(input)
	err := validate(content)
	if err != nil {
		return nil, false, err
	}
	idx := strings.Index(content, lastModTag)
	if idx == - 1 {
		return addLastMod(input, lastMod), false, nil
	}
	updatedContent, updated := changeLastMod(input, lastMod)
	return updatedContent, updated, nil
}

func changeLastMod(content []byte, lastMod time.Time) ([]byte, bool) {
	replace := fmt.Sprintf("%s: %s\n", lastModTag, lastMod.Format(timeFormat))
	result := replaceModRegex.ReplaceAll(content, []byte(replace))

	return result, isDifferent(content, result)
}

func isDifferent(first, second []byte) bool {
	if len(first) != len(second) {
		return true
	}
	for i := range first {
		if first[i] != second[i] {
			return true
		}
	}
	return false
}

func addLastMod(content []byte, lastMod time.Time) []byte {
	replace := fmt.Sprintf("---\n%s: %s\n", lastModTag, lastMod.Format(timeFormat))
	return addModRegex.ReplaceAllLiteral(content, []byte(replace))
}

var WithoutFrontMatter = errors.New("file not contains front matter section")

func validate(content string) error {
	if strings.Count(content, frontMatterDividor) < 2 {
		return WithoutFrontMatter
	}
	return nil
}
