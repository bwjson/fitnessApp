package rest

import "github.com/gin-gonic/gin"

type Handler struct {
	//services *service.Service
}

//func NewHandler(services *service.Service) *Handler {
//	return &Handler{services: services}
//}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/register", h.register)
		auth.POST("/login", h.login)
	}

	//api := router.Group("/api", h.userIdentity)
	//{
	//	lists := api.Group("/lists")
	//	{
	//		lists.POST("/", h.createList)
	//		lists.GET("/:id", h.getListById)
	//		lists.GET("/", h.getAllLists)
	//		lists.PUT("/:id", h.updateList)
	//		lists.DELETE("/:id", h.deleteList)
	//
	//		items := lists.Group(":id/items")
	//		{
	//			items.POST("/", h.createItem)
	//			items.GET("/", h.getAllItems)
	//		}
	//	}
	//
	//	items := api.Group("/items")
	//	{
	//		items.GET("/:id", h.getItemById)
	//		items.PUT("/:id", h.updateItem)
	//		items.DELETE("/:id", h.deleteItem)
	//	}
	//}

	return router
}
