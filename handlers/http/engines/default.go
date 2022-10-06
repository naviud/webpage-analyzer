package engines

import (
	"github.com/gin-gonic/gin"
	"github.com/naviud/webpage-analyzer/handlers/http/controllers"
)

type DefaultEngine struct {
	webPageAnalyzeController *controllers.WebPageAnalyzerController
}

func NewDefaultEngine(ctrl *controllers.WebPageAnalyzerController) *DefaultEngine {
	return &DefaultEngine{webPageAnalyzeController: ctrl}
}

func (d *DefaultEngine) GetDefaultEngine() *gin.Engine {
	engine := gin.New()

	engine.GET("/v1/analyze", d.webPageAnalyzeController.AnalyzeWebPage)

	return engine
}
