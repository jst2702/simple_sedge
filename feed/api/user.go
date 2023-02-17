package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"simplesedge.com/feed/pkg/db"
)

type createUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.CreateUserParams{
		Email:          req.Email,
		HashedPassword: req.Password,
	}
	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type getUserRequest struct {
	Email string `uri:"email" binding:"required,email"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}
