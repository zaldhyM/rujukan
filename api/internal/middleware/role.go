package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RequireRole memastikan user yang sedang login memiliki salah satu role yang diizinkan.
// Harus dipasang setelah AuthMiddleware karena bergantung pada context "role".
func RequireRole(roles ...string) gin.HandlerFunc {
	allowed := make(map[string]struct{}, len(roles))
	for _, r := range roles {
		allowed[r] = struct{}{}
	}

	return func(c *gin.Context) {
		roleVal, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Unauthorized"})
			c.Abort()
			return
		}

		role, _ := roleVal.(string)
		if _, ok := allowed[role]; !ok {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Akses ditolak: role '" + role + "' tidak memiliki izin untuk aksi ini",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
