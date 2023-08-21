package main

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadRules(t *testing.T) {
	items, err := LoadRules(filepath.Join("testdata", "conf"))
	require.NoError(t, err)
	require.Equal(t, []Rule{
		{
			Pattern: "hello",
			Days:    3,
		}, {
			Pattern: "world",
			Days:    4,
		},
	}, items)
}

func TestExpandRules(t *testing.T) {
	rules := []Rule{
		{
			Pattern: "/hello/**/*.log",
			Days:    3,
		},
	}
	ExpandRules(&rules)

	require.Equal(t, []Rule{
		{
			Pattern: "/hello/*.log",
			Days:    3,
		},
		{
			Pattern: "/hello/*/*.log",
			Days:    3,
		},
		{
			Pattern: "/hello/*/*/*.log",
			Days:    3,
		},
		{
			Pattern: "/hello/*/*/*/*.log",
			Days:    3,
		},
		{
			Pattern: "/hello/*/*/*/*/*.log",
			Days:    3,
		},
	}, rules)

	rules = []Rule{
		{
			Pattern: "/hello/world.og",
			Days:    3,
		},
	}
	ExpandRules(&rules)

	require.Equal(t, []Rule{
		{
			Pattern: "/hello/world.og",
			Days:    3,
		},
	}, rules)
}
