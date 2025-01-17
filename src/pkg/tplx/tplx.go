package tplx

import (
	"html/template"
	"net/url"
	"regexp"
	"strings"
)

var TemplateFuncMap = template.FuncMap{
	"escape":              url.PathEscape,
	"unescaped":           Unescaped,
	"urlconvert":          Urlconvert,
	"timeformat":          Timeformat,
	"timestamp":           Timestamp,
	"args":                Args,
	"reReplaceAll":        ReReplaceAll,
	"match":               regexp.MatchString,
	"toUpper":             strings.ToUpper,
	"toLower":             strings.ToLower,
	"contains":            strings.Contains,
	"humanize":            Humanize,
	"humanize1024":        Humanize1024,
	"humanizeDuration":    HumanizeDuration,
	"humanizePercentage":  HumanizePercentage,
	"humanizePercentageH": HumanizePercentageH,
}
