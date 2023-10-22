package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"enigmacamp.com/enigma-laundry-apps/utils/security"
	"github.com/gin-gonic/gin"
)

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

// Middleware sederhana untuk otentikasi
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Di sini Anda dapat menambahkan logika otentikasi, seperti memeriksa token atau session
		// Jika otentikasi berhasil, lanjutkan ke penanganan rute
		// Jika otentikasi gagal, Anda dapat menolak akses atau mengirimkan respons kesalahan

		// Contoh: Simulasi otentikasi sederhana
		// tokenString := c.GetHeader("Authorization")

		// if tokenString == "" {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not provided"})
		// 	c.Abort()
		// 	return
		// }

		var header authHeader

		if err := c.ShouldBindHeader(&header); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not provided"})
			c.Abort()
			return
		}

		token := strings.Split(header.AuthorizationHeader, " ")[1]

		fmt.Println(token)

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not provided"})
			c.Abort()
			return
		}

		claims, err := security.VerifyAccessToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Simpan data pengguna dalam konteks
		c.Set("username", claims["username"])

		c.Next()

	}
}

func ValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
		} else {
			c.Set("validatedInput", input)
			c.Next()
		}
	}
}

// Middleware untuk memeriksa otorisasi admin
func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")

		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}

		c.Next()
	}
}
