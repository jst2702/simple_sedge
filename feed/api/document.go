package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"simplesedge.com/feed/pkg/db"
)

type getDocumentRequest struct {
	Guid string `uri:"guid" binding:"required"`
}

func (server *Server) getDocument(ctx *gin.Context) {
	var req getDocumentRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	document, err := server.store.GetDocument(ctx, req.Guid)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, document)
}

type listDocumentsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=10"`
}

func (server *Server) listDocuments(ctx *gin.Context) {
	var req listDocumentsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListDocumentsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	documents, err := server.store.ListDocuments(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, documents)
}
