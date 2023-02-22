package api

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"simplesedge.com/feed/pkg/db"
	"simplesedge.com/gokit/util"
)

type createUserResponse struct {
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

type createUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,password"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if !server.validEmail(ctx, req.Email) {
		return
	}

	hashedPassword, err := util.HashPassword(util.RandomString(6))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Email:          req.Email,
		HashedPassword: hashedPassword,
	}
	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				log.Println(pqErr.Code.Name())
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := createUserResponse{
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.PasswordChangedAt,
	}
	ctx.JSON(http.StatusOK, rsp)
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

// this is one way to do input validation, but in this case, it would be better
// to use a custom validator. see password validation for example.
func (server *Server) validEmail(ctx *gin.Context, email string) bool {
	if len(email) < 10 {
		err := fmt.Errorf("email is too short, desu, %d", 10)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	} else if strings.Contains(email, "yahoo.com") {
		err := errors.New("yahoo isn't a real email")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	} else {
		return true
	}
}
