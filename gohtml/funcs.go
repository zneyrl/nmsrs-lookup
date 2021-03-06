package gohtml

import (
	"html/template"

	"github.com/emurmotol/nmsrs/env"
	"github.com/emurmotol/nmsrs/helper"
	"github.com/emurmotol/nmsrs/lang"
)

func Funcs() template.FuncMap {
	return template.FuncMap{
		"BaseUrl":        helper.BaseUrl,
		"config":         env.Map,
		"lang":           lang.Get,
		"dateForHumans":  helper.DateForHumans,
		"parsleyDateNow": helper.ParsleyDateNow,
		"photoPath":      helper.PhotoPath,
	}
}
