package api

import (
	"github.com/gin-contrib/cors"
	"github.com/khalil9022/OJT_Postgre/controller"
)

func (s *server) SetupRouter() {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "DELETE", "PUT", "GET"},
	}))

	stagingCustomerRepo := controller.NewRepository(s.DB)
	stagingCustomerService := controller.NewService(stagingCustomerRepo)
	stagingCustomerHandler := controller.NewHandler(stagingCustomerService)

	s.Router.GET("/", stagingCustomerHandler.GetDataCustomer)
}
