package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github/quocdaitrn/golang-gin-poc/entity"
)

const (
	TITLE       = "Video Title"
	DESCRIPTION = "Video Description"
	URL         = "https://youtu.be/JgW-i2QjgHQ"
)

func getVideo() entity.Video {
	return entity.Video{
		Title:       TITLE,
		Description: DESCRIPTION,
		URL:         URL,
	}
}

func TestFindAll(t *testing.T) {
	svc := New()

	video := getVideo()
	svc.Save(&video)

	videos := svc.FindAll()

	firstVideo := videos[0]

	assert.NotNil(t, videos)
	assert.Equal(t, TITLE, firstVideo.Title)
	assert.Equal(t, DESCRIPTION, firstVideo.Description)
	assert.Equal(t, URL, firstVideo.URL)
}
