package test

import (
	"github.com/naviud/webpage-analyzer/handlers/http/responses"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SetHtmlVersion_HappyPath(t *testing.T) {
	version := "version 1"
	w := responses.NewWebPageAnalyzerResponseManager()
	w.SetHtmlVersion(version)
	assert.Equal(t, w.To().HtmlVersion, version)
}

func Test_SetTitle_HappyPath(t *testing.T) {
	title := "title 1"
	w := responses.NewWebPageAnalyzerResponseManager()
	w.SetTitle(title)
	assert.Equal(t, w.To().Title, title)
}

func Test_AddHeadingLevel_HappyPath(t *testing.T) {
	tag := "h1"
	level := "l1"
	w := responses.NewWebPageAnalyzerResponseManager()
	w.AddHeadingLevel(tag, level)
	ob := w.To()
	assert.Equal(t, len(ob.Headings), 1)
	assert.Equal(t, ob.Headings[0].TagName, tag)
	assert.Equal(t, ob.Headings[0].Levels[0], level)

	level1 := "l2"

	w.AddHeadingLevel(tag, level1)
	ob1 := w.To()
	assert.Equal(t, len(ob1.Headings), 1)
	assert.Equal(t, ob1.Headings[0].TagName, tag)
	assert.Equal(t, len(ob1.Headings[0].Levels), 2)
	assert.Equal(t, ob1.Headings[0].Levels[1], level1)

	tag2 := "h2"

	w.AddHeadingLevel(tag2, level)
	ob2 := w.To()
	assert.Equal(t, len(ob2.Headings), 2)
	assert.Equal(t, ob2.Headings[1].TagName, tag2)
	assert.Equal(t, len(ob2.Headings[1].Levels), 1)
	assert.Equal(t, ob2.Headings[1].Levels[0], level)
}

func Test_AddUrlInfo_HappyPath(t *testing.T) {
	url := "url1"
	urlType := 1
	status := 200
	latency := 200
	w := responses.NewWebPageAnalyzerResponseManager()
	w.AddUrlInfo(url, urlType, status, int64(latency))
	ob := w.To()
	assert.Equal(t, len(ob.Urls), 1)
	assert.Equal(t, ob.Urls[0].Url, url)
	assert.Equal(t, ob.Urls[0].Type, "External")
	assert.Equal(t, ob.Urls[0].Status, status)
	assert.Equal(t, ob.Urls[0].Latency, int64(latency))

	url1 := "url2"
	urlType1 := 0
	status1 := 500
	latency1 := 501
	w.AddUrlInfo(url1, urlType1, status1, int64(latency1))
	ob1 := w.To()
	assert.Equal(t, len(ob1.Urls), 2)
	assert.Equal(t, ob1.Urls[1].Url, url1)
	assert.Equal(t, ob1.Urls[1].Type, "Internal")
	assert.Equal(t, ob1.Urls[1].Status, status1)
	assert.Equal(t, ob1.Urls[1].Latency, int64(latency1))
}

func Test_SetHasLogin_HappyPath(t *testing.T) {
	hasLogin := true
	w := responses.NewWebPageAnalyzerResponseManager()
	w.SetHasLogin(hasLogin)
	assert.Equal(t, w.To().HasLogin, hasLogin)
}

func Test_ToString_HappyPath(t *testing.T) {
	hasLogin := false
	w := responses.NewWebPageAnalyzerResponseManager()
	w.SetHasLogin(hasLogin)
	s := w.ToString()
	strToCompare := "\"hasLogin\":false}"
	assert.Contains(t, s, strToCompare)
}
