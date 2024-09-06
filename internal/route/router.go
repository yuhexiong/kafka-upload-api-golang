package route

import (
	kafkaUploadRoute "kafka-upload-api-golang/internal/kafka-upload/route"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	kafkaUploadRoute.SetupKafkaUploadRoutes(r)

	return r
}
