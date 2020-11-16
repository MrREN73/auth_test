package handler

import (
	"auth/src/common/httpserver"
	"auth/src/common/httpserver/middleware"
	"auth/src/common/utils"
	"auth/src/models"
	"auth/src/service"

	"github.com/gin-gonic/gin"
)

type handler struct {
	Srv  *httpserver.HTTPServer
	user *service.Service
}

func Serve(s *service.Service) utils.Stop {
	h := &handler{Srv: httpserver.New(), user: s}

	h.Srv.Router.Use(
		middleware.CorsMiddleware(),
	)

	users := h.Srv.Router.Group("users")

	users.POST("register", h.Register)
	users.GET("confirm")
	users.POST("login")
	users.GET("find")
	users.GET("me")
	users.GET("/:id")
	users.PUT("/:id")
	//Для создания и удаления можно написать handlerFunc для проверки прав доступа и передавать ее в аргументах
	users.POST("entity")
	users.DELETE("/:id")

	go h.Srv.ListenAndServe()

	return h.Srv.Shutdown
}

func (h handler) Register(c *gin.Context) {
	var params models.User

	if err := c.ShouldBindJSON(&params); err != nil {
		//Тут рендерим фейл

		return
	}

	user, err := h.user.CreateUser(params)
	if err != nil {
		//Рендерим ошибку
		//Я бы добавил сюда обвязку, куда будем передавать возвращаемые данные у функции и далее рендерим исходя из этих данных
		return
	}

	//c.Render(http.StatusOK, //тут делаем респонс рендер и отдаем сюда модель юзера user)
}

//остальные функции делаем анлогично, где то парсим json, где то uri.
