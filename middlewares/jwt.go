package middleware

import (
	util "attendance-is/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		token := util.ExtractToken(c)
		if token == "" || util.ValidateToken(token) != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthenticated",
			})
			defer c.AbortWithStatus(http.StatusUnauthorized)
		}
		claim, _ := util.ExtractTokenClaim(token)
		c.Set("user", claim)
		c.Next()
	})
}

type Claim struct {
	Uid    string
	Type   string
	TypeID string
}

func IsStudent() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		val, exist := c.Get("user")
		// id, _ := strconv.Atoi(c.Param("studentid"))
		claim, _ := val.(Claim)

		if !exist {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthenticated",
			})
			defer c.AbortWithStatus(http.StatusUnauthorized)
		}

		fmt.Println(claim.TypeID)
		fmt.Println(claim.Type)
		fmt.Println(claim.Uid)

		if claim.Type != "student" || claim.TypeID != c.Param("studentid") {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Unauthorized",
			})
			defer c.AbortWithStatus(http.StatusForbidden)
		}
	})
}

func IsLecturer() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		val, exist := c.Get("user")
		claim, _ := val.(util.JWTClaim)

		if !exist {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthenticated",
			})
			defer c.AbortWithStatus(http.StatusUnauthorized)
		}

		if claim.Type != "lecturer" {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Unauthorized",
			})
			defer c.AbortWithStatus(http.StatusForbidden)
		}
	})
}

func IsPBM() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		val, exist := c.Get("user")
		claim, _ := val.(util.JWTClaim)

		if !exist {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthenticated",
			})
			defer c.AbortWithStatus(http.StatusUnauthorized)
		}

		if claim.Type != "pbm" {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Unauthorized",
			})
			defer c.AbortWithStatus(http.StatusForbidden)
		}
	})
}
