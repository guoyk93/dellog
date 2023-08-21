package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestDaysFromFilename(t *testing.T) {
	date := time.Now().AddDate(0, 0, -5)
	d, ok := DaysFromFilename(fmt.Sprintf("hello.%04d-%02d-%02d.log", date.Year(), date.Month(), date.Day()))
	require.True(t, ok)
	require.Equal(t, 5, d)
	d, ok = DaysFromFilename(fmt.Sprintf("hello.%04d%02d%02d.log", date.Year(), date.Month(), date.Day()))
	require.True(t, ok)
	require.Equal(t, 5, d)
	d, ok = DaysFromFilename(fmt.Sprintf("hello.%04d_%02d_%02d", date.Year(), date.Month(), date.Day()))
	require.True(t, ok)
	require.Equal(t, 5, d)
}

func TestEvaluateDays(t *testing.T) {
	date := time.Now().AddDate(0, 0, -5)

	a := EvaluateDays(fmt.Sprintf("hello.%04d-%02d-%02d.log", date.Year(), date.Month(), date.Day()), 1)
	require.Equal(t, ActionRemove, a)

	a = EvaluateDays(fmt.Sprintf("hello.%04d-%02d-%02d.log", date.Year(), date.Month(), date.Day()), 7)
	require.Equal(t, ActionSkip, a)

	a = EvaluateDays(fmt.Sprintf("hello.%02d-%02d.log", date.Month(), date.Day()), 7)
	require.Equal(t, ActionUnknownDate, a)
}
