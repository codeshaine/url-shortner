package test

import (
	"testing"

	"github.com/codeshaine/url-shortner/internal/utils"
)

func TestRegexFunc(t *testing.T) {
	t.Run("should return true", func(t *testing.T) {
		got := utils.IsValidUrl("https://www.google-dev_hex.com/main?value1=test+value2=test2")
		want := true

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("should return false cause of invalid character", func(t *testing.T) {
		got := utils.IsValidUrl("https://www.fb'.com/main?value1=test+value2/test2")
		want := false
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
