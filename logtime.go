package logtime

import (
	"fmt"
	"log"
	"regexp"
	"time"
)

func LogLineMoment(text string, matcher *regexp.Regexp, layout string) *time.Time {
	matches := matcher.FindAllString(text, -1)
	if len(matches) > 0 {
		parse, err := time.Parse(layout, matches[0])
		if err != nil {
			log.Printf("Failed to parse input line: '%s', error: %v", text, err)
			return nil
		}
		return &parse
	} else {
		return nil
	}
}

func FromLayoutToPatternMatcher(logTime string) *regexp.Regexp {
	extractedPattern := logTime
	extractedPattern = regexp.QuoteMeta(extractedPattern)
	extractedPattern = regexp.MustCompile(`[0-9]`).ReplaceAllString(extractedPattern, `\d`)
	extractedPattern = regexp.MustCompile(`\+`).ReplaceAllString(extractedPattern, `\\+`)
	return regexp.MustCompile(fmt.Sprintf("^%s", extractedPattern))
}
