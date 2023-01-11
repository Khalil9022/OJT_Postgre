package api

import (
	"github.com/gin-contrib/cors"
	"github.com/khalil9022/OJT_Postgre/controller/checklistpencairan"
	"github.com/khalil9022/OJT_Postgre/controller/procesvalidation"
	"github.com/khalil9022/OJT_Postgre/controller/skalaangsuran"
)

func (s *server) SetupRouter() {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "DELETE", "PUT", "GET"},
		AllowHeaders: []string{"Origin", "Accept", "Content-Type", "Authorization", "Access-Control-Allow-Origin"},
	}))

	stagingCustomerRepo := procesvalidation.NewRepository(s.DB)
	stagingCustomerService := procesvalidation.NewService(stagingCustomerRepo)
	stagingCustomerHandler := procesvalidation.NewHandler(stagingCustomerService)

	generateSkalaAngsuranRepo := skalaangsuran.NewRepository(s.DB)
	generateSkalaAngsuranService := skalaangsuran.NewService(generateSkalaAngsuranRepo)
	generateSkalaAngsuranHandler := skalaangsuran.NewHandler(generateSkalaAngsuranService)

	checklistpencairanRepo := checklistpencairan.NewRepository(s.DB)
	checklistpencairanService := checklistpencairan.NewService(checklistpencairanRepo)
	checklistpencairanHandler := checklistpencairan.NewHandler(checklistpencairanService)

	s.Router.GET("/Customer", stagingCustomerHandler.PencairanKredit)
	s.Router.GET("/Generate", generateSkalaAngsuranHandler.GenerateSkalaAngsuran)

	s.Router.GET("/branch", checklistpencairanHandler.GetDataBranch)
	s.Router.GET("/company", checklistpencairanHandler.GetDataCompany)
	s.Router.GET("/allcustomer", checklistpencairanHandler.GetAllCustomerAs9)
	s.Router.POST("/spesifikcustomer", checklistpencairanHandler.GetSpesifikCustomerAs9)
}
