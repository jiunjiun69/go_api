package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/service"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/go-programming-tour-book/blog-service/pkg/convert"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
)

type Fish struct{}

func NewFish() Fish {
	return Fish{}
}

// @Summary 獲取單個魚隻
// @Produce json
// @Param id path int true "魚隻ID"
// @Success 200 {object} model.Fish "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/fishs/{id} [get]
func (a Fish) Get(c *gin.Context) {
	param := service.FishRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	fish, err := svc.GetFish(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetFish err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetFishFail)
		return
	}

	response.ToResponse(fish)
	return
}

// @Summary 獲取多個魚隻
// @Produce json
// @Param name query string false "魚隻名稱"
// @Param tag_id query int false "標籤ID"
// @Param state query int false "狀態"
// @Param page query int false "頁碼"
// @Param page_size query int false "每頁數量"
// @Success 200 {object} model.FishSwagger "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/fishs [get]

func (a Fish) List(c *gin.Context) {
	param := service.FishListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	fishs, totalRows, err := svc.GetFishList(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetFishList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetFishsFail)
		return
	}

	response.ToResponseList(fishs, totalRows)
	return
}

// @Summary 建立魚隻
// @Produce json
// @Param tag_id body string true "標籤ID"
// @Param title body string true "魚隻名稱"
// @Param desc body string false "魚隻介紹"
// @Param cover_image_url body string true "封面圖片地址"
// @Param content body string true "魚隻習性"
// @Param created_by body int true "建立者"
// @Param state body int false "狀態"
// @Success 200 {object} model.Fish "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/fishs [post]
func (a Fish) Create(c *gin.Context) {
	param := service.CreateFishRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateFish(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CreateFish err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateFishFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 更新魚隻
// @Produce json
// @Param tag_id body string false "標籤ID"
// @Param title body string false "魚隻名稱"
// @Param desc body string false "魚隻介紹"
// @Param cover_image_url body string false "封面圖片地址"
// @Param content body string false "魚隻習性"
// @Param modified_by body string true "修改者"
// @Success 200 {object} model.Fish "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/fishs/{id} [put]
func (a Fish) Update(c *gin.Context) {
	param := service.UpdateFishRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateFish(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateFish err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateFishFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 刪除魚隻
// @Produce  json
// @Param id path int true "魚隻ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/fishs/{id} [delete]
func (a Fish) Delete(c *gin.Context) {
	param := service.DeleteFishRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteFish(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.DeleteFish err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteFishFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}
