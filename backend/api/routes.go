package api

import (
	"github.com/gin-contrib/cors"
	"github.com/khalil9022/OJT_Postgre/controller/pencairanreport"
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

	pencairanreportRepo := pencairanreport.NewRepository(s.DB)
	pencairanreportService := pencairanreport.NewService(pencairanreportRepo)
	pencairanreportHandler := pencairanreport.NewHandler(pencairanreportService)

	s.Router.GET("/Customer", stagingCustomerHandler.PencairanKredit)
	s.Router.GET("/Generate", generateSkalaAngsuranHandler.GenerateSkalaAngsuran)

	s.Router.GET("/branch", pencairanreportHandler.GetDataBranch)
	s.Router.GET("/company", pencairanreportHandler.GetDataCompany)
	s.Router.GET("/allcustomer", pencairanreportHandler.GetAllCustomerAs9)
	s.Router.POST("/spesifikcustomer", pencairanreportHandler.GetSpesifikCustomerAs9)
	s.Router.GET("/allcustomerreport", pencairanreportHandler.GetAllCustomerAs01)
	s.Router.POST("/spesifikcustomerreport", pencairanreportHandler.GetSpesifikCustomerAs01)
	s.Router.POST("/updateapproval", pencairanreportHandler.UpdateApprovalStatus)
}
