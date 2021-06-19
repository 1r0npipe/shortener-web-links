package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateShortUrl(t *testing.T) {

	tests := []struct {
		longLink  string
		shortLink string
	}{
		{longLink: "www.yandex.ru/super-puper-long-link", 
		shortLink: "QJNdBNdA"},
		{longLink: "www.revenue.ie/this-is%fully-fake-link%no-sense-to-use-it", 
		shortLink: "MJZ9gmvb"},
		{longLink: "www.google.com/i-dont-know%what%to%write%right-here-a?", 
		shortLink: "ZvrgR7LP"},
	}
	for _, tt := range tests {
		shortLink, _ := GenerateShortUrl(tt.longLink)
		assert.Equal(t, shortLink, tt.shortLink)

	}
}
