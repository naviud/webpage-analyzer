package test

import (
	"github.com/golang/mock/gomock"
	"github.com/naviud/webpage-analyzer/analyzers"
	"github.com/naviud/webpage-analyzer/channels"
	"github.com/naviud/webpage-analyzer/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_AnalyzeLink_HappyPath(t *testing.T) {
	htmlData := "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n\t<meta charset='utf-8'>\n\t<meta name='viewport' content='width=device-width,initial-scale=1'>\n\t<title>Test App</title>\n</head>\n<body>\n\t<a href=\"http://link1.l\">a</a>\n\t<a href=\"http://link2.l\">b</a>\n</body>\n</html>"

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAnalyzerInfo := mocks.NewMockAnalyzerInfo(controller)
	mockAnalysis := mocks.NewMockWebPageAnalyzerResponseManager(controller)
	mockUrlExecProvider := mocks.NewMockUrlExecutorProvider(controller)

	var m = channels.NewTestUrlExecutor()

	mockAnalyzerInfo.EXPECT().GetBody().Return(htmlData).MinTimes(1)
	mockAnalyzerInfo.EXPECT().GetHost().Return("link1").MinTimes(1)
	mockUrlExecProvider.EXPECT().Provide().Return(m).MinTimes(1)

	c1 := mockAnalysis.EXPECT().AddUrlInfo(gomock.Eq("http://link1.l"), gomock.Eq(0), gomock.Any(), gomock.Any()).
		Do(func(u1 interface{}, t1 interface{}, s1 interface{}, l1 interface{}) {
			assert.Equal(t, u1, "http://link1.l")
			assert.Equal(t, t1, 0)
		}).AnyTimes()
	c2 := mockAnalysis.EXPECT().AddUrlInfo(gomock.Eq("http://link2.l"), gomock.Eq(1), gomock.Any(), gomock.Any()).
		Do(func(u2 interface{}, t2 interface{}, s2 interface{}, l2 interface{}) {
			assert.Equal(t, u2, "http://link2.l")
			assert.Equal(t, t2, 1)
		}).AnyTimes()

	gomock.InOrder(
		c1,
		c2,
	)

	a := analyzers.NewLinkAnalyzer(mockUrlExecProvider)

	a.Analyze(mockAnalyzerInfo, mockAnalysis)
}
