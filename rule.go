package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Rule struct {
	Pattern string
	Days    int
}

func LoadRules(dir string) (rules []Rule, err error) {
	var entries []os.DirEntry

	if entries, err = os.ReadDir(dir); err != nil {
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		log.Println("loading rules:", entry.Name())

		var content []byte
		if content, err = os.ReadFile(filepath.Join(dir, entry.Name())); err != nil {
			return
		}

		for _, _line := range bytes.Split(content, []byte{'\n'}) {
			line := string(bytes.TrimSpace(_line))
			if len(line) == 0 {
				continue
			}
			if strings.HasPrefix(line, "#") {
				continue
			}
			log.Println("rule:", line)
			splits := strings.SplitN(line, ":", 2)
			if len(splits) != 2 {
				continue
			}
			var days int
			if days, err = strconv.Atoi(splits[1]); err != nil {
				return
			}
			if days < 1 {
				err = fmt.Errorf("invalid days: %d", days)
				return
			}
			rules = append(rules, Rule{
				Pattern: splits[0],
				Days:    days,
			})
		}
	}
	return
}

const DoubleAsteriskPath = "**" + string(filepath.Separator)

func ExpandRules(rules *[]Rule) {
	var expanded []Rule
	for _, rule := range *rules {
		if strings.Contains(rule.Pattern, DoubleAsteriskPath) {
			for i := 0; i < 5; i++ {
				expanded = append(expanded, Rule{
					Pattern: strings.ReplaceAll(
						rule.Pattern,
						DoubleAsteriskPath,
						strings.Repeat("*"+string(filepath.Separator), i),
					),
					Days: rule.Days,
				})
			}
		} else {
			expanded = append(expanded, rule)
		}
	}
	*rules = expanded
}
