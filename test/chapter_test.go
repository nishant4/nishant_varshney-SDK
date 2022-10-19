package main

import (
	"lotr"
	"math"
	"testing"
)

const (
	TOTAL_CHAPTERS = 62
)

func TestChapters(t *testing.T) {
	app := lotr.New(ACCESS_CODE)

	chapterList, err := app.ListChapters()
	if err != nil {
		t.Fatal("Error ListChapters ", err)
	}

	if chapterList.Total != TOTAL_CHAPTERS {
		t.Errorf("Wrong Total Chapters: %d != %d (Should be)", chapterList.Total, TOTAL_CHAPTERS)
	}

	chapterId := chapterList.Chapters[0].Id
	chapter, err := app.GetChapter(chapterId)
	if err != nil {
		t.Fatal("GetChapter : Error: ", err)
	}

	if chapter.Chapters[0] != chapterList.Chapters[0] {
		t.Errorf("GetChapter: chapter should match: %v != %v", chapter.Chapters[0], chapterList.Chapters[0])
	}

	// options
	limit := 10
	maxPage := int(math.Ceil(float64(TOTAL_CHAPTERS) / float64(limit)))
	options := lotr.NewGetOptionsPageLimited(maxPage, limit)
	chapterList, err = app.ListChaptersOptions(options)
	if err != nil {
		t.Fatal("Error ListChaptersOptions ", err)
	}

	if chapterList.Total != TOTAL_CHAPTERS {
		t.Errorf("Wrong Total Chapters: %d != %d (Should be)", chapterList.Total, TOTAL_CHAPTERS)
	}

	if len(chapterList.Chapters) != (TOTAL_CHAPTERS % limit) {
		t.Errorf("Chapters Count for last page: %d != %d(Should be)", len(chapterList.Chapters), (TOTAL_CHAPTERS % limit))
	}
}
