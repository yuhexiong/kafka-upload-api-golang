package controller

import (
	"kafka-upload-api-golang/internal/kafka-upload/service"
	"kafka-upload-api-golang/pkg/response"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func UploadKafka(c *gin.Context) {
	var dataUri struct {
		Topic string `uri:"topic" binding:"required"`
	}
	if err := c.ShouldBindUri(&dataUri); err != nil {
		response.Error(c, err.Error())
		return
	}

	var dataBody map[string]interface{}
	if err := c.ShouldBind(&dataBody); err != nil {
		slog.Error(err.Error())
		response.Error(c, err.Error())
		return
	}

	var kafkaUploadService service.KafkaUploadService
	result := kafkaUploadService.Upload(dataUri.Topic, dataBody)
	response.Success(c, result)
}
