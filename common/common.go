package common

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

var (
	ColorBlack   = &Color{}
	ColorBlue    = &Color{Blue: 255}
	ColorCyan    = &Color{Blue: 255, Green: 255}
	ColorGray    = &Color{Blue: 128, Green: 128, Red: 128}
	ColorGreen   = &Color{Green: 128}
	ColorLime    = &Color{Green: 255}
	ColorMagenta = &Color{Blue: 255, Red: 255}
	ColorMaroon  = &Color{Red: 128}
	ColorNavy    = &Color{Blue: 128}
	ColorOlive   = &Color{Green: 128, Red: 128}
	ColorPurple  = &Color{Blue: 128, Red: 128}
	ColorRed     = &Color{Red: 255}
	ColorSilver  = &Color{Blue: 192, Green: 192, Red: 192}
	ColorTeal    = &Color{Blue: 128, Green: 128}
	ColorYellow  = &Color{Green: 255, Red: 255}
	ColorWhite   = &Color{Blue: 255, Green: 255, Red: 255}
)

type Color struct {
	Alpha, Blue, Green, Red uint8
}

type Region struct {
	Id string
}
type Item struct {
	Commnets []string
	StartAt  time.Duration
	EndAt    time.Duration
	Lines    []Line
}

type Line struct {
	Items     []LineItem
	VoiceName string
}

type LineItem struct {
	Text string
}

type Point struct {
	X float64
	Y float64
}

func ParsePersentageToString(value string, includedUnit bool) (retString string, err error) {
	var ret float64
	if includedUnit {
		var sign = strings.Index(value, "%")
		if sign < 1 {
			return "", errors.Errorf("In parsing percentage : %q", value)
		}
		ret, err = strconv.ParseFloat(value[:sign-1], 32)
	} else {
		ret, err = strconv.ParseFloat(value, 32)
	}

	if err != nil {
		return "", errors.Wrap(err, "In parssing persentage")
	}
	if ret < 0 || ret > 100 {
		return "", errors.Errorf("In parsing perstnage : the value(%q) is not in the bounday(0-100)", ret)
	}
	retString = fmt.Sprintf("%f", ret)
	return retString, nil
}

func ParsePersentage(value string) (ret float64, err error) {
	var sign = strings.Index(value, "%")
	if sign < 1 {
		return 0, errors.Errorf("In parsing percentage : %q", value)
	}
	ret, err = strconv.ParseFloat(value[:sign-1], 32)

	if err != nil {
		return 0, errors.Wrap(err, "In parssing persentage")
	}
	if ret < 0 || ret > 100 {
		return 0, errors.Errorf("In parsing perstnage : the value(%q) is not in the bounday(0-100)", ret)
	}
	return
}
