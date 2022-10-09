package test

import (
	"github.com/golang/mock/gomock"
	"github.com/naviud/webpage-analyzer/analyzers"
	"github.com/naviud/webpage-analyzer/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_AnalyzeHeading_HappyPath(t *testing.T) {
	htmlData := "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n\t<meta charset='utf-8'>\n\t<meta name='viewport' content='width=device-width,initial-scale=1'>\n\t<title>Test App</title>\n</head>\n<body>\n\t<h1>head1</h1>\n\t<h2><span>head2</span></h2>\n</body>\n</html>"

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAnalyzerInfo := mocks.NewMockAnalyzerInfo(controller)
	mockAnalysis := mocks.NewMockWebPageAnalyzerResponseManager(controller)

	mockAnalyzerInfo.EXPECT().GetBody().Return(htmlData).MinTimes(1)

	c1 := mockAnalysis.EXPECT().AddHeadingLevel(gomock.Eq("h1"), gomock.Eq("head1")).
		Do(func(k1 interface{}, v1 interface{}) {
			assert.Equal(t, k1, "h1")
			assert.Equal(t, v1, "head1")
		}).AnyTimes()
	c2 := mockAnalysis.EXPECT().AddHeadingLevel(gomock.Eq("h2"), gomock.Eq("head2")).
		Do(func(k2 interface{}, v2 interface{}) {
			assert.Equal(t, k2, "h2")
			assert.Equal(t, v2, "head2")
		}).AnyTimes()

	gomock.InOrder(
		c1,
		c2,
	)

	a := analyzers.NewHeadingAnalyzer()
	a.Analyze(mockAnalyzerInfo, mockAnalysis)
}
