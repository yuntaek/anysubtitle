package webvtt

import "github.com/anysub/interfaces"

type Stylesheet struct {
	interfaces.IStyleAttributes
	location      string
	parent        *Stylesheet
	ownerNode     string
	ownerCSSRuld  string
	media         string
	title         string
	alternateFlag string
	originClean   string
}

func NewStyleSheet() *Stylesheet {
	return &Stylesheet{
		location:      "null",
		alternateFlag: "unset",
		originClean:   "set",
	}
}
