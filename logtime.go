package logtime

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
	"sort"
	"time"
)

type LogTime struct {
	matcher *regexp.Regexp
	layout  string
}

func NewLogTime(layout string) *LogTime {
	return &LogTime{
		layout:  layout,
		matcher: FromLayoutToPatternMatcher(layout),
	}
}

type Moment struct {
	t time.Time
	l string
	d time.Duration
}

func (lt *LogTime) logLineMoment(text string) *time.Time {
	matches := lt.matcher.FindAllString(text, -1)
	if len(matches) > 0 {
		parse, err := time.Parse(lt.layout, matches[0])
		if err != nil {
			log.Printf("Failed to parse input line: '%s', error: %v", text, err)
			return nil
		}
		return &parse
	} else {
		return nil
	}
}

func (lt *LogTime) ReadStreamOfLogLines(scanner *bufio.Scanner) {
	var moments = make([]Moment, 0)
	for scanner.Scan() {
		l := scanner.Text()
		t := lt.logLineMoment(l)
		if t != nil {
			newMoment := NewMoment(*t, l)
			var previousMoment *Moment
			lenMoments := len(moments)
			if lenMoments > 0 {
				previousMoment = &moments[lenMoments-1]
			}
			if previousMoment != nil {
				if previousMoment.t == *t {
					moments[lenMoments-1] = newMoment
				} else {
					moments = append(moments, newMoment)
					previousMoment.d = newMoment.t.Sub(previousMoment.t)
				}
			} else {
				moments = append(moments, newMoment)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Detected %v distinct moments", len(moments))
	sort.SliceStable(moments, func(i, j int) bool {
		return moments[i].d > moments[j].d
	})
	for _, moment := range moments {
		fmt.Printf("%f\t%s", moment.d.Seconds(), moment.l)
	}
}

func NewMoment(t time.Time, l string) Moment {
	return Moment{
		t: t,
		l: l,
		d: 0,
	}
}

func FromLayoutToPatternMatcher(logTime string) *regexp.Regexp {
	extractedPattern := logTime
	extractedPattern = regexp.QuoteMeta(extractedPattern)
	extractedPattern = regexp.MustCompile(`[0-9]`).ReplaceAllString(extractedPattern, `\d`)
	extractedPattern = regexp.MustCompile(`\+`).ReplaceAllString(extractedPattern, `\\+`)
	return regexp.MustCompile(fmt.Sprintf("^%s", extractedPattern))
}
