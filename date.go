package main

import (
	"errors"
	"regexp"
	"time"
)

var (
	ErrNoDateFound = errors.New("no date found in filename")
)

var (
	regexpDate      = regexp.MustCompile(`(^|\D)(?P<y>\d{4})(\D)?(?P<m>\d{2})(\D)?(?P<d>\d{2})(\D|$)`)
	regexpDateYear  = regexpDate.SubexpIndex("y")
	regexpDateMonth = regexpDate.SubexpIndex("m")
	regexpDateDay   = regexpDate.SubexpIndex("d")
)

func DateFromFilename(name string) (date time.Time, err error) {
	matched := regexpDate.FindStringSubmatch(name)
	if len(matched) == 0 {
		err = ErrNoDateFound
		return
	}
	date, err = time.ParseInLocation(
		"2006-01-02",
		matched[regexpDateYear]+"-"+matched[regexpDateMonth]+"-"+matched[regexpDateDay],
		time.Local,
	)
	return
}
