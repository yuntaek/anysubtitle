package webvtt

import (
	"regexp"

	"github.com/anysub/interfaces"
)

const (
	BlockNameComment = "NOTE"
	BlockNameRegion  = "REGION"
	BlockNameStyle   = "STYLE"
	BlockNameCue     = "-->"
)

// Regular expression patterns
var (
	RegExprComment          = regexp.MustCompile(`^NOTE\b\s*`)
	RegExprRegion           = regexp.MustCompile(`^REGION\s*$`)
	RegExprStyle            = regexp.MustCompile(`^STYLE\s*$`)
	RegExprSettingDelimiter = regexp.MustCompile(`\s*`)
	RegExprEmptyLine        = regexp.MustCompile(`^\s*$`)
	RegExprSignature        = regexp.MustCompile(`^WEBVTT\s*`)
	RegExprSeconds          = regexp.MustCompile(`^([0-5][0-9]).([0-9]{3})$`)
	RegExprNumber           = regexp.MustCompile(`[[:digit:]]+`)
	RegExprFloatNumber      = regexp.MustCompile(`^-?[[:digit:]]+(\.{1}[[:digit:]]+)?$`)
)

type Metadata struct {
	interfaces.IMetaData
}
