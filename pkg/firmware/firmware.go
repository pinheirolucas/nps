package firmware

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrBadFormat = errors.New("bad firmware format")
)

type Firmware struct {
	major int
	minor int
}

func Parse(s string) (*Firmware, error) {
	segments := strings.Split(s, ".")
	if len(segments) != 2 {
		return nil, ErrBadFormat
	}

	majorStr, minorStr := segments[0], segments[1]

	major, err := strconv.Atoi(majorStr)
	if err != nil {
		return nil, ErrBadFormat
	}

	minor, err := strconv.Atoi(minorStr)
	if err != nil {
		return nil, ErrBadFormat
	}

	return &Firmware{
		major: major,
		minor: minor,
	}, nil
}

func ParsePS3(s string) (*Firmware, error) {
	return nil, nil
}

func ParsePSV(s string) (*Firmware, error) {
	return nil, nil
}

func ParsePSP(s string) (*Firmware, error) {
	return nil, nil
}

func (f *Firmware) String() string {
	return fmt.Sprintf("%d.%02d", f.major, f.minor)
}
