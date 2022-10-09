package test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/naviud/webpage-analyzer/handlers/http/controllers"
	"github.com/naviud/webpage-analyzer/handlers/http/responses"
	"github.com/naviud/webpage-analyzer/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Analyze_HappyPath(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	report := responses.AnalysisSuccessResponse{
		HtmlVersion: "V5",
		Title:       "Test",
		Headings:    nil,
		Urls:        nil,
		HasLogin:    false,
	}
	url := "url1"
	host := "host1"
	body := "body1"

	mockAnalysis := mocks.NewMockWebPageAnalyzerResponseManager(controller)
	mockBodyExtractor := mocks.NewMockBodyExtractor(controller)
	a1 := mocks.NewMockAnalyzer(controller)
	a2 := mocks.NewMockAnalyzer(controller)

	mockAnalysis.EXPECT().To().Return(report)
	mockBodyExtractor.EXPECT().Extract(gomock.Eq(url)).Return(host, body, nil)
	a1.EXPECT().Analyze(gomock.Any(), gomock.Eq(mockAnalysis))
	a2.EXPECT().Analyze(gomock.Any(), gomock.Eq(mockAnalysis))

	c := controllers.NewWebPageAnalyzerController(mockBodyExtractor, a1, a2)
	r, err := c.Analyze(mockAnalysis, url)

	assert.Equal(t, err, nil)
	assert.Equal(t, r.HtmlVersion, "V5")
}

func Test_Analyze_FailurePath(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	url := "url1"
	err1 := errors.New("test error")

	mockAnalysis := mocks.NewMockWebPageAnalyzerResponseManager(controller)
	mockBodyExtractor := mocks.NewMockBodyExtractor(controller)
	a1 := mocks.NewMockAnalyzer(controller)
	a2 := mocks.NewMockAnalyzer(controller)

	mockBodyExtractor.EXPECT().Extract(gomock.Eq(url)).Return("", "", err1)

	c := controllers.NewWebPageAnalyzerController(mockBodyExtractor, a1, a2)
	r, err := c.Analyze(mockAnalysis, url)

	assert.Equal(t, err.Error(), err1.Error())
	assert.Equal(t, r.HasLogin, false)
}
