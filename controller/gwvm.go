package controller

import (
	// "encoding/json"
	"Glue-API/httputil"
	"Glue-API/model"
	gluevm "Glue-API/utils/gwvm"
	"net/http"

	"github.com/gin-gonic/gin"
	// "os/exec"
)

// VmState godoc
//
//	@Summary		State of Gateway VM
//	@Description	gwvm의 상태를 보여줍니다.
//	@param			hypervisorType	path	string	true	"Hypervisor Type"
//	@Tags			Gwvm
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Success		200	{object}	model.GwvmMgmt
//	@Failure		400	{object}	httputil.HTTP400BadRequest
//	@Failure		404	{object}	httputil.HTTP404NotFound
//	@Failure		500	{object}	httputil.HTTP500InternalServerError
//	@Router			/api/v1/gwvm/{hypervisorType} [get]
func (c *Controller) VmState(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")

	var dat model.GwvmMgmt
	hypervisorType := ctx.Param("hypervisorType")

	message, err := gluevm.VmState(hypervisorType)

	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	dat.Message = message
	ctx.IndentedJSON(http.StatusOK, dat)
}

// VmDetail godoc
//
//	@Summary		Detail of Gateway VM
//	@Description	gwvm의 상세정보 상태를 보여줍니다.
//	@param			hypervisorType	path	string	true	"Hypervisor Type"
//	@Tags			Gwvm
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Success		200	{object}	model.GwvmMgmt
//	@Failure		400	{object}	httputil.HTTP400BadRequest
//	@Failure		404	{object}	httputil.HTTP404NotFound
//	@Failure		500	{object}	httputil.HTTP500InternalServerError
//	@Router			/api/v1/gwvm/detail/{hypervisorType} [get]
func (c *Controller) VmDetail(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")

	var dat model.GwvmMgmt
	hypervisorType := ctx.Param("hypervisorType")

	message, err := gluevm.VmDetail(hypervisorType)

	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	dat.Message = message
	ctx.IndentedJSON(http.StatusOK, dat)
}

// VmSetup godoc
//
//	@Summary		Setup Gateway Vm
//	@Description	gwvm을 생성합니다.
//	@param			hypervisorType			path		string	true	"Hypervisor Type"
//	@param			gwvmMngtNicParent		formData	string	true	"Gwvm Management Nic Parent"
//	@param			gwvmMngtNicIp			formData	string	true	"Gwvm Management Nic Ip"
//	@param			gwvmStorageNicParent	formData	string	true	"Gwvm Storage Nic Parent"
//	@param			gwvmStorageNicIp		formData	string	true	"Gwvm Storage Nic Ip"
//	@Tags			Gwvm
//	@Accept			multipart/form-data
//	@Produce		json
//	@Success		200	{object}	model.GwvmMgmt
//	@Failure		400	{object}	httputil.HTTP400BadRequest
//	@Failure		404	{object}	httputil.HTTP404NotFound
//	@Failure		500	{object}	httputil.HTTP500InternalServerError
//	@Router			/api/v1/gwvm/{hypervisorType} [post]
func (c *Controller) VmSetup(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")

	var dat model.GwvmMgmt

	hypervisorType := ctx.Param("hypervisorType")
	gwvmCpu, _ := ctx.GetPostForm("gwvmCpu")
	gwvmMemory, _ := ctx.GetPostForm("gwvmMemory")
	gwvmMngtNicParent, _ := ctx.GetPostForm("gwvmMngtNicParent")
	gwvmMngtNicIp, _ := ctx.GetPostForm("gwvmMngtNicIp")
	gwvmStorageNicParent, _ := ctx.GetPostForm("gwvmStorageNicParent")
	gwvmStorageNicIp, _ := ctx.GetPostForm("gwvmStorageNicIp")

	message, err := gluevm.VmSetup(hypervisorType, gwvmCpu, gwvmMemory, gwvmMngtNicParent, gwvmMngtNicIp, gwvmStorageNicParent, gwvmStorageNicIp)

	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	dat.Message = message
	ctx.IndentedJSON(http.StatusOK, dat)
}

// VmStart godoc
//
//	@Summary		Start to Gateway VM
//	@Description	Gateway VM을 실행합니다.
//	@param			hypervisorType	path	string	true	"Hypervisor Type"
//	@Tags			Gwvm
//	@Accept			multipart/form-data
//	@Produce		json
//	@Success		200	{object}	model.GwvmMgmt
//	@Failure		400	{object}	httputil.HTTP400BadRequest
//	@Failure		404	{object}	httputil.HTTP404NotFound
//	@Failure		500	{object}	httputil.HTTP500InternalServerError
//	@Router			/api/v1/gwvm/start/{hypervisorType} [put]
func (c *Controller) VmStart(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")

	var dat model.GwvmMgmt

	hypervisorType := ctx.Param("hypervisorType")

	message, err := gluevm.VmStart(hypervisorType)

	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	dat.Message = message
	ctx.IndentedJSON(http.StatusOK, dat)
}

func (c *Controller) VmStartOptions(ctx *gin.Context) {
	SetOptionHeader(ctx)
	ctx.IndentedJSON(http.StatusOK, nil)
}

// VmStop godoc
//
//	@Summary		Stop to Gateway VM
//	@Description	Gateway VM을 정지합니다.
//	@param			hypervisorType	path	string	true	"Hypervisor Type"
//	@Tags			Gwvm
//	@Accept			multipart/form-data
//	@Produce		json
//	@Success		200	{object}	model.GwvmMgmt
//	@Failure		400	{object}	httputil.HTTP400BadRequest
//	@Failure		404	{object}	httputil.HTTP404NotFound
//	@Failure		500	{object}	httputil.HTTP500InternalServerError
//	@Router			/api/v1/gwvm/stop/{hypervisorType} [put]
func (c *Controller) VmStop(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")

	var dat model.GwvmMgmt

	hypervisorType := ctx.Param("hypervisorType")

	message, err := gluevm.VmStop(hypervisorType)

	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	dat.Message = message
	ctx.IndentedJSON(http.StatusOK, dat)
}

func (c *Controller) VmStopOptions(ctx *gin.Context) {
	SetOptionHeader(ctx)
	ctx.IndentedJSON(http.StatusOK, nil)
}

// VmDelete godoc
//
//	@Summary		Delete to Gateway VM
//	@Description	Gateway VM을 삭제합니다.
//	@param			hypervisorType	path	string	true	"Hypervisor Type"
//	@Tags			Gwvm
//	@Accept			multipart/form-data
//	@Produce		json
//	@Success		200	{object}	model.GwvmMgmt
//	@Failure		400	{object}	httputil.HTTP400BadRequest
//	@Failure		404	{object}	httputil.HTTP404NotFound
//	@Failure		500	{object}	httputil.HTTP500InternalServerError
//	@Router			/api/v1/gwvm/delete/{hypervisorType} [delete]
func (c *Controller) VmDelete(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")

	var dat model.GwvmMgmt

	hypervisorType := ctx.Param("hypervisorType")

	message, err := gluevm.VmDelete(hypervisorType)

	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	dat.Message = message
	ctx.IndentedJSON(http.StatusOK, dat)
}

func (c *Controller) VmDeleteOptions(ctx *gin.Context) {
	SetOptionHeader(ctx)
	ctx.IndentedJSON(http.StatusOK, nil)
}

// VmCleanup godoc
//
//	@Summary		Cleanup to Gateway VM
//	@Description	Gateway VM Pcs cluster를 Cleanup 합니다.
//	@param			hypervisorType	path	string	true	"Hypervisor Type"
//	@Tags			Gwvm
//	@Accept			multipart/form-data
//	@Produce		json
//	@Success		200	{object}	model.GwvmMgmt
//	@Failure		400	{object}	httputil.HTTP400BadRequest
//	@Failure		404	{object}	httputil.HTTP404NotFound
//	@Failure		500	{object}	httputil.HTTP500InternalServerError
//	@Router			/api/v1/gwvm/cleanup/{hypervisorType} [put]
func (c *Controller) VmCleanup(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")

	var dat = struct {
		Message string
	}{}
	hypervisorType := ctx.Param("hypervisorType")

	message, err := gluevm.VmCleanup(hypervisorType)

	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	dat.Message = message
	ctx.IndentedJSON(http.StatusOK, dat)
}

func (c *Controller) VmCleanupOptions(ctx *gin.Context) {
	SetOptionHeader(ctx)
	ctx.IndentedJSON(http.StatusOK, nil)
}

// VmCleanup godoc
//
//	@Summary		VmMigrate to Gateway VM
//	@Description	Gateway VM을 Pcs cluster내 다른 호스트로 마이그레이션 합니다.
//	@param			hypervisorType	path	string	true	"Hypervisor Type"
//	@param			target		formData	string	true	"Migration Target Host"
//	@Tags			Gwvm
//	@Accept			multipart/form-data
//	@Produce		json
//	@Success		200	{object}	model.GwvmMgmt
//	@Failure		400	{object}	httputil.HTTP400BadRequest
//	@Failure		404	{object}	httputil.HTTP404NotFound
//	@Failure		500	{object}	httputil.HTTP500InternalServerError
//	@Router			/api/v1/gwvm/migrate/{hypervisorType} [put]
func (c *Controller) VmMigrate(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")

	var dat model.GwvmMgmt

	hypervisorType := ctx.Param("hypervisorType")
	target, _ := ctx.GetPostForm("target")

	message, err := gluevm.VmMigrate(hypervisorType, target)

	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	dat.Message = message
	ctx.IndentedJSON(http.StatusOK, dat)
}

func (c *Controller) VmMigrateOptions(ctx *gin.Context) {
	SetOptionHeader(ctx)
	ctx.IndentedJSON(http.StatusOK, nil)
}
