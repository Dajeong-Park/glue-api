package controller

import (
	"Glue-API/httputil"
	"Glue-API/model"
	"Glue-API/utils"
	"Glue-API/utils/fs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (c *Controller) FsOption(ctx *gin.Context) {
	SetOptionHeader(ctx)
	ctx.IndentedJSON(http.StatusOK, nil)
}

// FsStatus godoc
//
//	@Summary		Show Status and List of Glue FS
//	@Description	GlueFS의 상태값과 리스트를 보여줍니다.
//	@Tags			GlueFS
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Success		200	{object}	FsStatus
//	@Failure		400	{object}	httputil.HTTP400BadRequest
//	@Failure		404	{object}	httputil.HTTP404NotFound
//	@Failure		500	{object}	httputil.HTTP500InternalServerError
//	@Router			/api/v1/gluefs [get]
func (c *Controller) FsStatus(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")

	dat, err := fs.FsStatus()
	if err != nil {
		utils.FancyHandleError(err)
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}
	dat2, err := fs.FsList()
	value := model.FsSum{
		FsStatus: dat,
		FsList:   dat2,
	}
	if err != nil {
		utils.FancyHandleError(err)
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}
	// Print the output
	ctx.IndentedJSON(http.StatusOK, value)
}

// FsCreate godoc
//
//	@Summary		Create of Glue FS
//	@Description	GlueFS를 생성합니다.
//	@param			fs_name 	path	string	true	"Glue FS Name"
//	@param			hosts 	formData	[]string	true	"Glue FS Service Host Name" collectionFormat(multi)
//	@Tags			GlueFS
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Success		200	{string}	string	"Success"
//	@Failure		400	{object}	httputil.HTTP400BadRequest
//	@Failure		404	{object}	httputil.HTTP404NotFound
//	@Failure		500	{object}	httputil.HTTP500InternalServerError
//	@Router			/api/v1/gluefs/{fs_name} [post]
func (c *Controller) FsCreate(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")

	fs_name := ctx.Param("fs_name")
	hosts, _ := ctx.GetPostFormArray("hosts")

	hosts_str := strings.Join(hosts, ",")
	dat, err := fs.FsCreate(fs_name, hosts_str)
	if err != nil {
		utils.FancyHandleError(err)
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}
	// Print the output
	ctx.IndentedJSON(http.StatusOK, dat)
}

// FsUpdate godoc
//
//	@Summary		Update of Glue FS
//	@Description	GlueFS를 수정합니다.
//	@param			old_name 	formData	string	true	"Glue FS Old Name"
//	@param			new_name 	formData	string	true	"Glue FS New Name"
//	@param			hosts 	formData	[]string	true	"Glue FS Service Host Name" collectionFormat(multi)
//	@Tags			GlueFS
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Success		200	{string}	string	"Success"
//	@Failure		400	{object}	httputil.HTTP400BadRequest
//	@Failure		404	{object}	httputil.HTTP404NotFound
//	@Failure		500	{object}	httputil.HTTP500InternalServerError
//	@Router			/api/v1/gluefs [put]
func (c *Controller) FsUpdate(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")

	old_name, _ := ctx.GetPostForm("old_name")
	new_name, _ := ctx.GetPostForm("new_name")
	hosts, _ := ctx.GetPostFormArray("hosts")

	hosts_str := strings.Join(hosts, ",")
	dat, err := fs.FsUpdate(old_name, new_name, hosts_str)
	if err != nil {
		utils.FancyHandleError(err)
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}
	// Print the output
	ctx.IndentedJSON(http.StatusOK, dat)
}

// FsDelete godoc
//
//	@Summary		Delete of Glue FS
//	@Description	GlueFS를 삭제합니다.
//	@param			fs_name 	path	string	true	"Glue FS Name"
//	@Tags			GlueFS
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Success		200	{string}	string	"Success"
//	@Failure		400	{object}	httputil.HTTP400BadRequest
//	@Failure		404	{object}	httputil.HTTP404NotFound
//	@Failure		500	{object}	httputil.HTTP500InternalServerError
//	@Router			/api/v1/gluefs/{fs_name} [delete]
func (c *Controller) FsDelete(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")

	fs_name := ctx.Param("fs_name")
	list, err := fs.SubVolumeGroupLs(fs_name)
	if err != nil {
		utils.FancyHandleError(err)
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}
	if len(list) != 0 {
		ctx.IndentedJSON(http.StatusOK, "Please Subvolume Group Check")
	} else {
		dat, err := fs.FsDelete(fs_name)
		if err != nil {
			utils.FancyHandleError(err)
			httputil.NewError(ctx, http.StatusInternalServerError, err)
			return
		}
		ctx.IndentedJSON(http.StatusOK, dat)
	}

}

// FsGetInfo godoc
//
//	@Summary		Detail Info of Glue FS
//	@Description	GlueFS의 상세 정보를 보여줍니다.
//	@param			fs_name 	path	string	true	"Glue FS Name"
//	@Tags			GlueFS
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Success		200	{object}	FsGetInfo
//	@Failure		400	{object}	httputil.HTTP400BadRequest
//	@Failure		404	{object}	httputil.HTTP404NotFound
//	@Failure		500	{object}	httputil.HTTP500InternalServerError
//	@Router			/api/v1/gluefs/info/{fs_name} [get]
func (c *Controller) FsGetInfo(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")

	fs_name := ctx.Param("fs_name")
	dat, err := fs.FsGetInfo(fs_name)
	if err != nil {
		utils.FancyHandleError(err)
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}
	// Print the output
	ctx.IndentedJSON(http.StatusOK, dat)
}
