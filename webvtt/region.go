package webvtt

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/anysub/common"
	"github.com/pkg/errors"

	"github.com/anysub/interfaces"
)

func isRegion(line []byte) bool {
	return RegExprRegion.Match(line)
}
func NewRegion() *Region {
	return &Region{
		Width:          100,
		Lines:          3,
		Anchor:         common.Point{X: 0.0, Y: 100.0},
		ViewPortAnchor: common.Point{X: 0.0, Y: 100.0},
		Scroll:         "none",
	}
}

type Region struct {
	interfaces.IRegion
	Id             string
	Width          int
	Lines          int
	Anchor         common.Point
	ViewPortAnchor common.Point
	Scroll         string
}

func (r *Region) Parse(s *bufio.Scanner) (err error) {
	for s.Scan() {
		// finding empty line, end of parsing region block
		if RegExprEmptyLine.Match(s.Bytes()) {
			break
		}
		line := s.Text()
		err = r.parseNameValue(line)
		return
	}
	return
}

func (r *Region) ToString() (ret string) {
	return
}

// Parse setting in Region
func (r *Region) parseNameValue(line string) (err error) {
	settings := RegExprSettingDelimiter.Split(line, -1)
	for _, setting := range settings {
		if strings.HasPrefix(setting, ":") || strings.HasSuffix(setting, ":") {
			continue
		}
		element := strings.Split(setting, ":")
		if len(element) > 2 {
			err = errors.Errorf("The setting is not valid name and value format : %q", setting)
			return
		}
		switch element[0] {
		case "id":
			if r.Id != "" {
				err = errors.Errorf("The ID setting is duplicated %q ", setting)
				return
			}
			r.Id = element[1]
		case "width":
			r.Width, err = strconv.Atoi(element[1])
			if err != nil {
				return err
			}
		case "lines":
			r.Lines, err = strconv.Atoi(element[1])
			if err != nil {
				return err
			}
		case "regionanchor":
			r.Anchor, err = r.parseAnchor(element[1])
			if err != nil {
				return err
			}
		case "viewportanchor":
			r.ViewPortAnchor, err = r.parseAnchor(element[1])
			if err != nil {
				return err
			}
		case "scroll":
			if element[1] == "up" {
				r.Scroll = "up"
			}
		}
	}
	return
}

func (r *Region) Transform(typeSubtitle string) (ret string) {
	return
}

func (r *Region) parseAnchor(value string) (point common.Point, err error) {
	anchor := strings.Split(value, ",")
	if len(anchor) != 2 {
		err = errors.Errorf("The value is not valid type of anchnor point: %q", value)
		return
	}
	point.X, err = common.ParsePersentage(anchor[0])
	if err != nil {
		err = errors.Wrap(err, "The X of achor could not be converted")
		return
	}
	point.Y, err = common.ParsePersentage(anchor[1])
	if err != nil {
		err = errors.Wrap(err, "The Y of achor could not be converted")
		return
	}
	return
}
