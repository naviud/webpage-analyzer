package engines

import (
	"github.com/gin-gonic/gin"
	"github.com/naviud/webpage-analyzer/handlers/http/controllers"
	"net/http"
	"path"
	"path/filepath"
)

type DefaultEngine struct {
	webPageAnalyzeController *controllers.WebPageAnalyzerController
}

func NewDefaultEngine(ctrl *controllers.WebPageAnalyzerController) *DefaultEngine {
	return &DefaultEngine{webPageAnalyzeController: ctrl}
}

// GetDefaultEngine function is responsible to create
// http routes.
func (d *DefaultEngine) GetDefaultEngine() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(CORSMiddleware())

	// This route serves the web page analysis for the
	// given URL via a query param.
	engine.GET("/v1/analyze", d.webPageAnalyzeController.AnalyzeWebPage)

	// This route serves the static content which is
	//the Angular frontend.
	engine.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("./frontend/webpage-analyzer-client/index.html")
		} else {
			c.File("./frontend/webpage-analyzer-client/" + path.Join(dir, file))
		}
	})

	return engine
}

// CORSMiddleware function is responsible to handle CORS
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, OPTIONS, POST, PUT")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
