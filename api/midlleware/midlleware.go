package middleware

import (
	"net/http"
	"strings"

	"github.com/Exam4/4th-month-exam-Api-Gateway/api/token"
	"github.com/gin-gonic/gin"
)

func MiddleWare() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		t := ctx.GetHeader("Authorization")

		url := ctx.Request.URL.Path

		if strings.Contains(url, "swagger") {

			ctx.Next()

			return

		} else if _, err := token.ExtractClaim(t); err != nil {

			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{

				"error": err.Error(),
			})

			return

		}

		ctx.Next()

	}

}
