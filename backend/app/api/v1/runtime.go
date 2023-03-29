package v1

import (
	"github.com/1Panel-dev/1Panel/backend/app/api/v1/helper"
	"github.com/1Panel-dev/1Panel/backend/app/dto"
	"github.com/1Panel-dev/1Panel/backend/app/dto/request"
	"github.com/1Panel-dev/1Panel/backend/constant"
	"github.com/gin-gonic/gin"
)

// @Tags Runtime
// @Summary List runtimes
// @Description 获取运行环境列表
// @Accept json
// @Param request body request.RuntimeSearch true "request"
// @Success 200
// @Security ApiKeyAuth
// @Router /runtimes/search [post]
func (b *BaseApi) SearchRuntimes(c *gin.Context) {
	var req request.RuntimeSearch
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	total, items, err := runtimeService.Page(req)
	if err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	helper.SuccessWithData(c, dto.PageResult{
		Total: total,
		Items: items,
	})
}