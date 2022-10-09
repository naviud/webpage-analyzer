package test

import (
	"github.com/golang/mock/gomock"
	"github.com/naviud/webpage-analyzer/analyzers"
	"github.com/naviud/webpage-analyzer/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_AnalyzeLoginForm_HappyPath(t *testing.T) {
	htmlData := "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n\t<meta charset='utf-8'>\n\t<meta name='viewport' content='width=device-width,initial-scale=1'>\n\t<title>Test App</title>\n</head>\n<body>\n\t<input type=\"password\" />\n</body>\n</html>"

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAnalyzerInfo := mocks.NewMockAnalyzerInfo(controller)
	mockAnalysis := mocks.NewMockWebPageAnalyzerResponseManager(controller)

	mockAnalyzerInfo.EXPECT().GetBody().Return(htmlData).MinTimes(1)
	mockAnalysis.EXPECT().SetHasLogin(gomock.Any()).
		Do(func(v interface{}) {
			assert.Equal(t, v, true)
		})

	a := analyzers.NewLoginFormAnalyzer()
	a.Analyze(mockAnalyzerInfo, mockAnalysis)
}
