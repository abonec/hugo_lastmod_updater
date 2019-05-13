package updater

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUpdater_updateAddLastMod(t *testing.T) {
	modTime := time.Now()
	expected := getContent(modTime)
	table := []struct {
		content  string
		expected string
	}{
		{
			content:  getContent(modTime.Add(100 * time.Hour)),
			expected: expected,
		},
		{
			content:  getContent(time.Time{}),
			expected: expected,
		},
	}
	for _, set := range table {
		content, err := updateLastMod([]byte(set.content), modTime)
		assert.NoError(t, err)
		assert.Equal(t, set.expected, string(content))
	}
}

func getContent(modTime time.Time) string {
	if modTime.IsZero() {
		return fmt.Sprintf(pattern, "")
	}
	return fmt.Sprintf(pattern, "\n"+lastModTag+": "+modTime.Format(timeFormat))
}

var pattern = `---%s
title: "title"
date: 2019-05-09T20:14:29+03:00
draft: false
sitemap:
    priority: 1
---
content
`
