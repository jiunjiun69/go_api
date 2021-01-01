package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/service"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/go-programming-tour-book/blog-service/pkg/convert"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
)

type Fishpower struct{}

func NewFishpower() Fishpower {
	return Fishpower{}
}

// @Summary 獲取單個魚隻活動力
// @Produce json
// @Param id path int true "魚隻活動力ID"
// @Success 200 {object} model.Fishpower "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/fishpowers/{id} [get]
func (a Fishpower) Get(c *gin.Context) {
	param := service.FishpowerRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	fishpower, err := svc.GetFishpower(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetFishpower err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetFishpowerFail)
		return
	}

	response.ToResponse(fishpower)
	return
}

// @Summary 獲取多個魚隻活動力
// @Produce json
// @Param name query string false "活動力狀態"
// @Param tag_id query int false "標籤ID"
// @Param state query int false "狀態"
// @Param page query int false "頁碼"
// @Param page_size query int false "每頁數量"
// @Success 200 {object} model.FishpowerSwagger "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/fishpowers [get]
/*
func (a Fishpower) List(c *gin.Context) {
	param := service.FishpowerListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	fishpowers, totalRows, err := svc.GetFishpowerList(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetFishpowerList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetFishpowersFail)
		return
	}

	response.ToResponseList(fishpowers, totalRows)
	return
}
*/
// @Summary 建立魚隻活動力
// @Produce json
// @Param tag_id body string true "標籤ID"
// @Param title body string true "活動力狀態"
// @Param power body string true "活動力數值"
// @Param created_by body int true "建立者"
// @Param state body int false "狀態"
// @Success 200 {object} model.Fishpower "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/fishpowers [post]
func (a Fishpower) Create(c *gin.Context) {
	param := service.CreateFishpowerRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateFishpower(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CreateFishpower err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateFishpowerFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 更新魚隻活動力
// @Produce json
// @Param tag_id body string false "標籤ID"
// @Param title body string false "活動力狀態"
// @Param power body string false "活動力數值"
// @Param modified_by body string true "修改者"
// @Success 200 {object} model.Fishpower "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/fishpowers/{id} [put]
func (a Fishpower) Update(c *gin.Context) {
	param := service.UpdateFishpowerRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateFishpower(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateFishpower err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateFishpowerFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 刪除魚隻活動力
// @Produce  json
// @Param id path int true "魚隻活動力ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/fishpowers/{id} [delete]
func (a Fishpower) Delete(c *gin.Context) {
	param := service.DeleteFishpowerRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteFishpower(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.DeleteFishpower err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteFishpowerFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}
