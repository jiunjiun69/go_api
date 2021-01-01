package routers

import (
	"net/http"
	"time"

	"github.com/go-programming-tour-book/blog-service/pkg/limiter"

	"github.com/go-programming-tour-book/blog-service/global"

	"github.com/gin-gonic/gin"
	_ "github.com/go-programming-tour-book/blog-service/docs"
	"github.com/go-programming-tour-book/blog-service/internal/middleware"
	"github.com/go-programming-tour-book/blog-service/internal/routers/api"
	"github.com/go-programming-tour-book/blog-service/internal/routers/api/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}

	r.Use(middleware.Tracing())
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	r.Use(middleware.Translations())

	tag := v1.NewTag()
	upload := api.NewUpload()
	fish := v1.NewFish()
	fishpower := v1.NewFishpower()
	r.GET("/debug/vars", api.Expvar)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload/file", upload.UploadFile)
	r.POST("/auth", api.GetAuth)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	apiv1 := r.Group("/api/v1")
	apiv1.Use() //middleware.JWT()
	{
		// 建立標籤
		apiv1.POST("/tags", tag.Create)
		// 刪除指定標籤
		apiv1.DELETE("/tags/:id", tag.Delete)
		// 更新指定標籤
		apiv1.PUT("/tags/:id", tag.Update)
		// 獲取標籤列表
		apiv1.GET("/tags", tag.List)

		// 建立魚隻
		apiv1.POST("/fishs", fish.Create)
		// 刪除指定魚隻
		apiv1.DELETE("/fishs/:id", fish.Delete)
		// 更新指定魚隻
		apiv1.PUT("/fishs/:id", fish.Update)
		// 獲取指定魚隻
		apiv1.GET("/fishs/:id", fish.Get)
		// 獲取魚隻列表
		apiv1.GET("/fishs", fish.List)

		// 建立活動力
		apiv1.POST("/fishpowers", fishpower.Create)
		// 刪除指定活動力
		apiv1.DELETE("/fishpowers/:id", fishpower.Delete)
		// 更新指定活動力
		apiv1.PUT("/fishpowers/:id", fishpower.Update)
		// 獲取指定活動力
		apiv1.GET("/fishpowers/:id", fishpower.Get)
		// 獲取活動力列表
		//apiv1.GET("/fishpowers", fishpower.List)

	}

	return r
}
