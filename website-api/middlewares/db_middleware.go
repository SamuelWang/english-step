package middlewares

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ctxKey string

const dbCtxKey ctxKey = "db"

func DBContextMiddleware(globalDB *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// attach DB bound to the current request context
		dbWithContext := globalDB.WithContext(c.Request.Context())

		// keep Gin-stored DB for handlers that call GetDBFromContext
		c.Set("db", dbWithContext)

		// store DB in the request.Context using a typed key so repos/services can read it
		ctx := context.WithValue(c.Request.Context(), dbCtxKey, dbWithContext)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

// GetDBFromContext returns *gorm.DB stored in gin.Context (for handlers).
func GetDBFromContext(c *gin.Context) *gorm.DB {
	db, exists := c.Get("db")
	if !exists {
		log.Println("Error: DB instance not set in Context. Please ensure DBContextMiddleware is registered.")
		return nil
	}
	return db.(*gorm.DB)
}

// DBFromContext returns a *gorm.DB from a plain context.Context.
// It also attaches the provided ctx to the returned DB to ensure queries
// use the caller's context (deadlines/cancelation). If no DB is present it returns nil.
func DBFromContext(ctx context.Context) *gorm.DB {
    v := ctx.Value(dbCtxKey)
    if v == nil {
        return nil
    }
    db := v.(*gorm.DB)
    // attach the caller's context (idempotent if already set to the same ctx)
    return db.WithContext(ctx)
}
