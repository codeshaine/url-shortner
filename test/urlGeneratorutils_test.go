package test

import (
	"testing"
	"unicode/utf8"

	utils "github.com/codeshaine/url-shortner/internal/utils"
)

func TestUrlGenerator(t *testing.T) {
	t.Run("testing the legnth", func(t *testing.T) {
		url := utils.GenerateUnqueUrl()
		length := utf8.RuneCountInString(url)

		if length != 10 {
			t.Errorf("expected length %d got %d", 32, length)
		}
	})

	t.Run("testing uniquness", func(t *testing.T) {
		urls := make(map[string]bool)
		for i := 0; i < 10000; i++ {
			url := utils.GenerateUnqueUrl()
			if urls[url] {
				t.Errorf("duplicate url")
			}
			urls[url] = true
		}
	})

}
