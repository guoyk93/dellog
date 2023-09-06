package main

import (
	"github.com/stretchr/testify/require"
	"path/filepath"
	"testing"
	"time"
)

func TestParseRule(t *testing.T) {
	rds, err := ParseRule([]byte(`
match: /hello/world
days: 3
---
match:
  - /hello/**/*.log
  - /world/**/*.csv
size: 2m
---
match:
  - /hello/**/*.log
  - /world/**/*.csv
size: 10241024
`))
	require.NoError(t, err)
	require.Equal(t, []RuleDoc{
		{Match: SingleOrMulti[string]{"/hello/world"}, Days: 3, Size: 0},
		{Match: SingleOrMulti[string]{"/hello/**/*.log", "/world/**/*.csv"}, Days: 0, Size: 2097152},
		{Match: SingleOrMulti[string]{"/hello/**/*.log", "/world/**/*.csv"}, Days: 0, Size: 10241024}},
		rds,
	)
}

func TestLoadRuleDir(t *testing.T) {
	items, err := LoadRuleDir(filepath.Join("testdata", "conf"))
	require.NoError(t, err)
	require.Equal(t, []Rule{
		{Match: []string{"/hello/world"}, Days: 0, Size: 2097152},
		{Match: []string{"/hello/*.log", "/hello/*/*.log", "/hello/*/*/*.log", "/hello/*/*/*/*.log", "/hello/*/*/*/*/*.log", "/world.log"}, Days: 3, Size: 0},
		{Match: []string{"/world/*.csv", "/world/*/*.csv", "/world/*/*/*.csv", "/world/*/*/*/*.csv", "/world/*/*/*/*/*.csv", "/world.log"}, Days: 0, Size: 3072}},
		items,
	)
}

func TestEvaluateRule(t *testing.T) {
	act := EvaluateRule(filepath.Join("testdata", "logs", "bydate", "test.2023-01-01.log"), Rule{
		Days: 3,
	}, time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local))
	require.Equal(t, ActionSkip, act)

	act = EvaluateRule(filepath.Join("testdata", "logs", "bydate", "test.2023-01-01.log"), Rule{
		Days: 3,
	}, time.Date(2023, 1, 8, 0, 0, 0, 0, time.Local))
	require.Equal(t, ActionRemove, act)

	act = EvaluateRule(filepath.Join("testdata", "logs", "bysize", "world.txt"), Rule{
		Days: 3,
		Size: 4,
	}, time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local))
	require.Equal(t, ActionTruncate, act)

	act = EvaluateRule(filepath.Join("testdata", "logs", "bysize", "world.txt"), Rule{
		Days: 3,
		Size: 100,
	}, time.Date(2023, 1, 8, 0, 0, 0, 0, time.Local))
	require.Equal(t, ActionSkip, act)
}
