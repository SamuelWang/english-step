package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DBContextMiddleware(globalDB *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		dbWithContext := globalDB.WithContext(c.Request.Context())
		c.Set("db", dbWithContext)

		c.Next()
	}
}

func GetDBFromContext(c *gin.Context) *gorm.DB {
	db, exists := c.Get("db")
	if !exists {
		log.Println("Error: DB instance not set in Context. Please ensure DBContextMiddleware is registered.")
		return nil
	}
	return db.(*gorm.DB)
}
