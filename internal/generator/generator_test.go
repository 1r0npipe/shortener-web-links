package generator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateShortUrl(t *testing.T) {
	tests := []struct {
		longLink  string
		userID    string
		shortLink string
	}{
		{longLink: "www.yandex.ru/super-puper-long-link",
			userID:    "1123-1233-1233",
			shortLink: "G2qypjzSMoa"},
		{longLink: "www.revenue.ie/this-is%fully-fake-link%no-sense-to-use-it",
			userID:    "1123-1233-1233",
			shortLink: "UfBNP3mcjX1"},
		{longLink: "www.google.com/i-dont-know%what%to%write%right-here-a?",
			userID:    "1123-1233-1233",
			shortLink: "TA6eiUvMY5b"},
	}
	for _, tt := range tests {
		shortLink, _ := GenerateShortUrl(tt.longLink, tt.userID)
		assert.Equal(t, shortLink, tt.shortLink)
	}
}
