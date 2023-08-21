package main

import (
	"regexp"
	"time"
)

var (
	regexpDate      = regexp.MustCompile(`(^|\D)(?P<y>\d{4})(\D)?(?P<m>\d{2})(\D)?(?P<d>\d{2})(\D|$)`)
	regexpDateYear  = regexpDate.SubexpIndex("y")
	regexpDateMonth = regexpDate.SubexpIndex("m")
	regexpDateDay   = regexpDate.SubexpIndex("d")
)

func DaysFromFilename(name string) (days int, ok bool) {
	matched := regexpDate.FindStringSubmatch(name)
	if len(matched) == 0 {
		return
	}
	date, err := time.ParseInLocation(
		"2006-01-02",
		matched[regexpDateYear]+"-"+matched[regexpDateMonth]+"-"+matched[regexpDateDay],
		time.Local,
	)
	if err != nil {
		return
	}
	days = int(time.Since(date) / (time.Hour * 24))
	ok = true
	return
}

type Action int

const (
	ActionUnknownDate Action = iota
	ActionSkip
	ActionRemove
)

func EvaluateDays(name string, threshold int) Action {
	days, ok := DaysFromFilename(name)
	if !ok {
		return ActionUnknownDate
	}
	if days <= threshold {
		return ActionSkip
	}
	return ActionRemove
}
