package engine

import (
	"fmt"
	"net/http"
	"path"
	"sync"

	"github.com/gin-gonic/gin"
)

type Engine struct {
	*gin.Engine

	mutex   sync.RWMutex
	Routers map[string]ResourceType
}

type RouterGroupWrapper struct {
	*gin.RouterGroup
	engine *Engine
}

// ResourceScope is the scope of the resource.
type ResourceScope string

// ResourceType used to bind API
type ResourceType struct {
	// Scope is the scope of the resource.
	Scope ResourceScope `json:"scope"`
	// Resource is the resource name.
	Resource string `json:"resource"`
}

func New() *Engine {
	return &Engine{
		Engine:  gin.New(),
		Routers: make(map[string]ResourceType),
	}
}

func FormatRoute(method string, path string) string {
	return fmt.Sprintf("%s %s", method, path)
}

func (e *Engine) Group(relativePath string, handlers ...gin.HandlerFunc) RouterGroupWrapper {
	group := e.Engine.Group(relativePath, handlers...)
	return RouterGroupWrapper{RouterGroup: group, engine: e}
}

func (rg RouterGroupWrapper) Group(relativePath string, handlers ...gin.HandlerFunc) RouterGroupWrapper {
	group := rg.RouterGroup.Group(relativePath, handlers...)
	return RouterGroupWrapper{RouterGroup: group, engine: rg.engine}
}

func (rg RouterGroupWrapper) GET(relativePath string, typ ResourceType, handlers ...gin.HandlerFunc) gin.IRoutes {
	absolutePath := path.Join(rg.BasePath(), relativePath)
	rg.engine.mutex.Lock()
	rg.engine.Routers[FormatRoute(http.MethodGet, absolutePath)] = typ
	rg.engine.mutex.Unlock()

	return rg.RouterGroup.GET(relativePath, handlers...)
}

func (rg RouterGroupWrapper) POST(relativePath string, typ ResourceType, handlers ...gin.HandlerFunc) gin.IRoutes {
	absolutePath := path.Join(rg.BasePath(), relativePath)
	rg.engine.mutex.Lock()
	rg.engine.Routers[FormatRoute(http.MethodPost, absolutePath)] = typ
	rg.engine.mutex.Unlock()

	return rg.RouterGroup.POST(relativePath, handlers...)
}

func (rg RouterGroupWrapper) PATCH(relativePath string, typ ResourceType, handlers ...gin.HandlerFunc) gin.IRoutes {
	absolutePath := path.Join(rg.BasePath(), relativePath)
	rg.engine.mutex.Lock()
	rg.engine.Routers[FormatRoute(http.MethodPatch, absolutePath)] = typ
	rg.engine.mutex.Unlock()

	return rg.RouterGroup.PATCH(relativePath, handlers...)
}

func (rg RouterGroupWrapper) PUT(relativePath string, typ ResourceType, handlers ...gin.HandlerFunc) gin.IRoutes {
	absolutePath := path.Join(rg.BasePath(), relativePath)
	rg.engine.mutex.Lock()
	rg.engine.Routers[FormatRoute(http.MethodPut, absolutePath)] = typ
	rg.engine.mutex.Unlock()

	return rg.RouterGroup.PUT(relativePath, handlers...)
}

func (rg RouterGroupWrapper) DELETE(relativePath string, typ ResourceType, handlers ...gin.HandlerFunc) gin.IRoutes {
	absolutePath := path.Join(rg.BasePath(), relativePath)
	rg.engine.mutex.Lock()
	rg.engine.Routers[FormatRoute(http.MethodDelete, absolutePath)] = typ
	rg.engine.mutex.Unlock()

	return rg.RouterGroup.DELETE(relativePath, handlers...)
}
