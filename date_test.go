package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDateFromFilename(t *testing.T) {
	d, err := DateFromFilename("2023-01-02.log")
	require.NoError(t, err)
	require.Equal(t, 2023, d.Year())
	require.Equal(t, 1, int(d.Month()))
	require.Equal(t, 2, d.Day())
	d, err = DateFromFilename("hello.2023-0102.log")
	require.NoError(t, err)
	require.Equal(t, 2023, d.Year())
	require.Equal(t, 1, int(d.Month()))
	require.Equal(t, 2, d.Day())
	d, err = DateFromFilename("world.202301-02.log")
	require.NoError(t, err)
	require.Equal(t, 2023, d.Year())
	require.Equal(t, 1, int(d.Month()))
	require.Equal(t, 2, d.Day())
	_, err = DateFromFilename("world.202391-02.log")
	require.Error(t, err)
	_, err = DateFromFilename("world.02.log")
	require.Error(t, err)
	require.Equal(t, ErrNoDateFound, err)
}
