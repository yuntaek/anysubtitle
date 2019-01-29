package interfaces

import (
	"bufio"
	"time"
)

type IColor interface {
	ToString() string
}

type IItem interface {
	ToString() string
}

type IStyleAttributes interface {
	ToString() string
	Propagate()
}

type IRegion interface {
	Parse(s *bufio.Scanner) error
	ToString() string
	Transform(s string) string
}

type ISubtitles interface {
	Open(o IOption) error
	OpenFile(filename string) error
	Add(d time.Duration)
	GetDuration() time.Duration
	alignDuration() (d time.Duration, addDummyItem bool)
	Fragment(f time.Duration) error
	IsEmpty() bool
	Merge() error
	Optimize() error
	removeUnUsedRegionsAndStyles() error
}

type IMetaData interface {
}

type IOption interface {
}
