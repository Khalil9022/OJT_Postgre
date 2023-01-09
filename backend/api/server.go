package api

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/khalil9022/OJT_Postgre/controller/procesvalidation"
	"github.com/khalil9022/OJT_Postgre/controller/skalaangsuran"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

type server struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func MakeServer(db *gorm.DB) *server {
	s := &server{
		Router: gin.Default(),
		DB:     db,
	}

	//Auto Run Service
	c := cron.New()
	c.AddFunc("@every 1m", func() {
		stagingCustomer := procesvalidation.NewRepository(s.DB)
		stagingCustomer.PencairanKredit()
	})

	c.AddFunc("@every 2m", func() {
		generateSkalaAngsuran := skalaangsuran.NewRepository(s.DB)
		generateSkalaAngsuran.GenerateSkalaAngsuran()
	})
	c.Start()
	return s
}

func (s *server) RunServer() {
	s.SetupRouter()
	port := os.Getenv("PORT")
	if err := s.Router.Run(":" + port); err != nil {
		panic(err)
	}

}
