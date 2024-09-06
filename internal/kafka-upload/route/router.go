package route

import (
	"kafka-upload-api-golang/internal/kafka-upload/controller"

	"github.com/gin-gonic/gin"
)

func SetupKafkaUploadRoutes(r *gin.Engine) {
	kafkaUploadGroup := r.Group("/kafka-upload")
	{
		kafkaUploadGroup.POST("/:topic", controller.UploadKafka)
	}
}
