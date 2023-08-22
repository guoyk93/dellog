package main

import (
	"gopkg.in/yaml.v3"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCapacity_MarshalYAML(t *testing.T) {
	buf, err := yaml.Marshal(map[string]any{"c": Capacity(1024)})
	require.NoError(t, err)
	require.Equal(t, "c: 1024\n", string(buf))
}

type testCapacityStruct struct {
	C Capacity `yaml:"c"`
}

func TestCapacity_UnmarshalYAML(t *testing.T) {
	{
		var v testCapacityStruct
		err := yaml.Unmarshal([]byte("c: 1024"), &v)
		require.NoError(t, err)
		require.Equal(t, Capacity(1024), v.C)
	}
	{
		var v testCapacityStruct
		err := yaml.Unmarshal([]byte("c: '1k'"), &v)
		require.NoError(t, err)
		require.Equal(t, Capacity(1024), v.C)
	}
}

func TestParseCapacity(t *testing.T) {
	v, err := ParseCapacity("1k")
	require.NoError(t, err)
	require.Equal(t, Capacity(1024), v)

	v, err = ParseCapacity("1m")
	require.NoError(t, err)
	require.Equal(t, Capacity(1024*1024), v)
}

func TestSingleOrMulti_MarshalYAML(t *testing.T) {
	var v = SingleOrMulti[string]{"hello", "world"}
	buf, err := yaml.Marshal(map[string]any{"c": v})
	require.NoError(t, err)
	require.Equal(t, "c:\n    - hello\n    - world\n", string(buf))
}

type testSingleOrMultiStruct struct {
	C SingleOrMulti[int] `yaml:"c"`
}

func TestSingleOrMulti_UnmarshalYAML(t *testing.T) {
	{
		var v testSingleOrMultiStruct
		err := yaml.Unmarshal([]byte(`c: 1`), &v)
		require.NoError(t, err)
		require.Equal(t, []int{1}, v.C.Unwrap())
	}
	{
		var v testSingleOrMultiStruct
		err := yaml.Unmarshal([]byte(`c: [1, 2]`), &v)
		require.NoError(t, err)
		require.Equal(t, []int{1, 2}, v.C.Unwrap())
	}
}

func TestExpandDoubleAsteriskPattern(t *testing.T) {
	rules := []string{
		"/hello/world.og",
		"/hello/**/*.log",
		"/world/**/*.csv",
	}
	ExpandDoubleAsteriskPattern(&rules)

	require.Equal(t, []string{
		"/hello/world.og",
		"/hello/*.log",
		"/hello/*/*.log",
		"/hello/*/*/*.log",
		"/hello/*/*/*/*.log",
		"/hello/*/*/*/*/*.log",
		"/world/*.csv",
		"/world/*/*.csv",
		"/world/*/*/*.csv",
		"/world/*/*/*/*.csv",
		"/world/*/*/*/*/*.csv",
	}, rules)
}
