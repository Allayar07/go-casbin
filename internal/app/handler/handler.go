package handler

import (
	"casbin-go_gin/internal/services"
	"github.com/casbin/casbin-pg-adapter"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Services *services.Service
	adapter  *pgadapter.Adapter
}

func NewHandler(service *services.Service, adapter *pgadapter.Adapter) *Handler {
	return &Handler{
		Services: service,
		adapter:  adapter,
	}
}
func (h *Handler) InitRoutes() *gin.Engine {
	//Use Casbin authentication middleware.
	//auth, err := gcasbin.NewCasbinMiddleware("path/RBAC_model.conf", "path/policy.csv", h.SubjectFromJWT)
	//if err != nil {
	//	log.Fatal(err)
	//}
	app := gin.Default()
	app.POST("/log-in", h.LogIn)
	app.POST("/create_user", h.CreateUser)
	app.GET("/read", h.Authorize("admin", "get", h.adapter), h.ReadBook)

	app.POST("/book", h.Authorize("organizations", "post", h.adapter), h.ReadAndWriteAndSoOn)
	return app
}
