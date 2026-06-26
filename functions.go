package gext

import (
	"html/template"
	"strings"
	"time"

	"gitlab.com/metakeule/fmtdate"
)

func getFuncions() map[string]interface{} {
	return template.FuncMap{
		"Now":        now,
		"ToUpper":    strings.ToUpper,
		"ToLower":    strings.ToLower,
		"TimeNow":    timeNow,
		"FormatDate": formatDate,
	}
}

func now() string {
	return time.Now().String()
}

func timeNow() time.Time {
	return time.Now()
}

func formatDate(date time.Time, format string) string {
	return fmtdate.Format(format, date)
}
