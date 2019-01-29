package webvtt

import (
	"bytes"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/anysub/common"
	"github.com/pkg/errors"
)

type Cue struct {
	ID                string
	PauseOneExit      bool
	Region            *Region
	RegionId          string
	WritingDirection  string
	SnapToLines       bool
	Line              string
	LineAlignment     string
	Position          string
	PositionAlignment string
	Size              string
	TextAlignment     string
	Text              string
	StrartedAt        time.Duration
	FinishedAt        time.Duration
}

func NewCue() *Cue {
	return &Cue{
		PauseOneExit:      false,
		WritingDirection:  "horizontal",
		Region:            nil,
		SnapToLines:       true,
		Line:              "auto",
		LineAlignment:     "start",
		Position:          "auto",
		PositionAlignment: "auto",
		Size:              "100",
		TextAlignment:     "center",
	}
}

//Validate if the block is cue
func isCue(line []byte) bool {
	return bytes.IndexAny(line, "-->") > 0
}

//Parse cue
func (c *Cue) Parse(line []byte) (err error) {
	next, err := c.ParseTimings(line)
	next = bytes.TrimLeft(next, " \t")
	if len(next) > 0 {
		err = c.ParseNameValue(string(next))
	}
	//TODO Parse cue text
	return
}

//Parse strat time and end time
func (c *Cue) ParseTimings(line []byte) (next []byte, err error) {
	timestamp := bytes.Split(line, []byte("-->"))
	if len(timestamp) != 2 {
		err = errors.Errorf("wrong format of cue timings %q", string(line))
		return
	}
	c.StrartedAt, err = c.ParseTimestamp(timestamp[0])
	if err != nil {
		err = errors.Wrap(err, "start time")
		return
	}

	nextOffset := bytes.IndexAny(timestamp[1], " \t")
	next = timestamp[1][nextOffset:]
	c.FinishedAt, err = c.ParseTimestamp(timestamp[1][:nextOffset-1])
	if err != nil {
		err = errors.Wrap(err, "finish time")
	}
	return
}

//ParseTimestamp parsing time stamp
func (c *Cue) ParseTimestamp(timestamp []byte) (dur time.Duration, err error) {
	raw := bytes.Trim(timestamp, (" \t"))
	units := bytes.Split(raw, []byte(":"))
	switch len(units) {
	//Most siginificant units is hours
	case 3:
		hour, _ := strconv.Atoi(string(units[0]))
		minute, _ := strconv.Atoi(string(units[1]))
		seconds := RegExprSeconds.FindSubmatch(units[2])
		s, _ := strconv.Atoi(string(seconds[1]))
		ms, _ := strconv.Atoi(string(seconds[2]))
		t := (hour*60*60+minute*60+s)*1000 + ms
		t = t * 1000000
		dur = time.Duration(t)

		//Most significant units is minutes
	case 2:
		minute, _ := strconv.Atoi(string(units[0]))
		seconds := RegExprSeconds.FindSubmatch(units[1])
		if len(seconds[0]) == 0 {
			err = errors.Errorf("It is wrong seconds format in timestamp %f", seconds)
			return 0, err
		}
		s, _ := strconv.Atoi(string(seconds[1]))
		ms, _ := strconv.Atoi(string(seconds[2]))
		t := (minute*60+s)*1000 + ms
		t = t * 1000000
		dur = time.Duration(t)
	default:
		err = errors.Errorf("It is wrong format timestamp %q")
		return
	}
	return
}

// Parse setting in cue
func (c *Cue) ParseNameValue(line string) (err error) {
	settings := RegExprSettingDelimiter.Split(line, -1)
	for _, setting := range settings {
		if strings.HasPrefix(setting, ":") || strings.HasSuffix(setting, ":") {
			continue
		}
		element := strings.Split(setting, ":")
		if len(element) > 2 {
			err = errors.Errorf("The cue setting is not valid name and value format : %q", setting)
			log.Printf("%v", err)
			continue
		}
		switch element[0] {

		case "region":
			if c.RegionId != "" {
				err = errors.Errorf("The region setting is duplicated in cue %q ", setting)
				log.Printf("%v", err)
				continue
			}
			c.RegionId = element[1]
		case "vertical":
			c.WritingDirection = element[1]
			c.RegionId = ""
		case "line":
			var linePos string
			var lineAlign string
			// not contain any ascii digit
			if RegExprNumber.MatchString(setting) == false {
				log.Printf("The line setting value is not number %q", setting)
				continue
			}
			value := strings.Split(element[1], ",")
			switch len(value) {
			case 2:
				linePos = value[0]
				lineAlign = value[1]
			default:
				linePos = element[1]
			}
			if strings.Index(linePos, "%") > 0 {
				c.Line, err = common.ParsePersentageToString(linePos, true)
				if err != nil {
					log.Printf("The line setting value in cue is wrong format : %v(%s) ", err, linePos)
					continue
				}
				c.SnapToLines = false
			} else {
				// check if the line position value float number
				if RegExprFloatNumber.MatchString(linePos) == false {
					log.Printf("The line setting value in cue is wrong format : %s", linePos)
					continue
					c.SnapToLines = true
				}
			}
			// line alignment validation
			if lineAlign == "start" || lineAlign == "end" || lineAlign == "center" || lineAlign == "" {
				c.LineAlignment = lineAlign
			} else {
				log.Printf("The line alignment setting value in cue is not valid : %s", lineAlign)
				continue
			}
		case "position":
			var colPos string
			var colAlign string

			value := strings.Split(element[1], ",")
			switch len(value) {
			case 2:
				colPos = value[0]
				colAlign = value[1]
			default:
				colPos = element[1]
			}
			c.Position, err = common.ParsePersentageToString(colPos, false)
			if err != nil {
				log.Printf("The position setting value in cue is wrong format : %v(%s) ", err, colPos)
				continue
			}
			// position alignment validation
			if colAlign == "line-left" || colAlign == "center" || colAlign == "line-right" || colAlign == "" {
				c.PositionAlignment = colAlign
			} else {
				log.Printf("The line alignment setting value in cue is not valid : %q", colAlign)
				continue
			}
		case "size":
			c.Size, err = common.ParsePersentageToString(element[1], false)
			if err != nil {
				log.Printf("The size setting value in cue is not valid %q", element[1])
				continue
			}
		case "align":
			switch element[1] {
			case "start", "center", "end", "left", "right":
				c.TextAlignment = element[1]
			default:
				log.Print("The align ins not valid vlaue %q", element[1])
				continue
			}
		}
	}
	return
}
