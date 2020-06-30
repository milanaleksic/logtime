package logtime

import (
	"reflect"
	"regexp"
	"testing"
)

func Test_fromLayoutToPatternMatcher(t *testing.T) {
	type args struct {
		logTime string
	}
	tests := []struct {
		name string
		args args
		want *regexp.Regexp
	}{
		{
			name: "2006-01-02 15:04:05",
			args: struct{ logTime string }{logTime: "2006-01-02 15:04:05"},
			want: regexp.MustCompile(`^\d\d\d\d-\d\d-\d\d \d\d:\d\d:\d\d`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromLayoutToPatternMatcher(tt.args.logTime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fromLayoutToPatternMatcher() = %v, want %v", got, tt.want)
			}
		})
	}
}
