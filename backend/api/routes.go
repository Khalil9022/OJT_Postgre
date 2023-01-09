package api

import (
	"github.com/gin-contrib/cors"
	"github.com/khalil9022/OJT_Postgre/controller/procesvalidation"
	"github.com/khalil9022/OJT_Postgre/controller/skalaangsuran"
)

func (s *server) SetupRouter() {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "DELETE", "PUT", "GET"},
	}))

	stagingCustomerRepo := procesvalidation.NewRepository(s.DB)
	stagingCustomerService := procesvalidation.NewService(stagingCustomerRepo)
	stagingCustomerHandler := procesvalidation.NewHandler(stagingCustomerService)

	generateSkalaAngsuranRepo := skalaangsuran.NewRepository(s.DB)
	generateSkalaAngsuranService := skalaangsuran.NewService(generateSkalaAngsuranRepo)
	generateSkalaAngsuranHandler := skalaangsuran.NewHandler(generateSkalaAngsuranService)

	s.Router.GET("/Customer", stagingCustomerHandler.PencairanKredit)
	s.Router.GET("/Generate", generateSkalaAngsuranHandler.GenerateSkalaAngsuran)
}
