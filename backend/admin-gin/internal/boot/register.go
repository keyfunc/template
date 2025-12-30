package boot

import (
	"mall-api/internal/app/admin/iam/auth"
	"mall-api/internal/app/admin/user"
	"mall-api/internal/pkg/cookie"
	"mall-api/internal/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Register(r *gin.Engine, db *gorm.DB, rdb *redis.Client, jt *jwt.JWT, cm *cookie.CookieManager) {
	// admin routes
	adminGroup := r.Group("/admin")
	{
		auth.Register(adminGroup, db, rdb, jt, cm)
		user.Register(adminGroup, db)
	}
}
