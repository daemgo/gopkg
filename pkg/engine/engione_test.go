package engine

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	// resource types for test
	bookResourceType = ResourceType{
		Scope:    ResourceScope("store"),
		Resource: "book",
	}
	reviewResourceType = ResourceType{
		Scope:    ResourceScope("store"),
		Resource: "review",
	}
	userResourceType = ResourceType{
		Scope:    ResourceScope("management"),
		Resource: "user",
	}
)

func TestNew(t *testing.T) {
	e := New()

	group := e.Group("/api/v1")

	group.GET("/users", userResourceType, func(c *gin.Context) {})
	group.POST("/books", bookResourceType, func(c *gin.Context) {})
	group.PATCH("/books/:book_id", bookResourceType, func(c *gin.Context) {})
	group.PUT("/users/:user_id", userResourceType, func(c *gin.Context) {})
	group.DELETE("/books/:book_id", bookResourceType, func(c *gin.Context) {})

	group2 := group.Group("/group")
	group2.GET("/reviews/:review_id", reviewResourceType, func(c *gin.Context) {})

	assert.Equal(t, e.Routers, map[string]ResourceType{
		"GET /api/v1/users":                    userResourceType,
		"POST /api/v1/books":                   bookResourceType,
		"PATCH /api/v1/books/:book_id":         bookResourceType,
		"PUT /api/v1/users/:user_id":           userResourceType,
		"DELETE /api/v1/books/:book_id":        bookResourceType,
		"GET /api/v1/group/reviews/:review_id": reviewResourceType,
	}, "checking engine.Routers")
}
