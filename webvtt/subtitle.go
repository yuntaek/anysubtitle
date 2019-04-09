package webvtt

import (
	"io"

	"github.com/anysub/interfaces"
)

type Subtitle struct {
	interfaces.ISubtitles
	Regions    map[string]*Region
	Stylesheet map[string]*Stylesheet
	Cues       []Cue
}

// Parse Webtvvt file and save into webvtt file
func (s *Subtitle) Parse(i io.Reader) error {

}

func (s *Subtitle) isValid() {

}
