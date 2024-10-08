package controller

import (
	"kafka-upload-api-golang/internal/kafka-upload/service"
	"kafka-upload-api-golang/pkg/response"

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
		response.Error(c, err.Error())
		return
	}

	var kafkaUploadService service.KafkaUploadService
	if err := kafkaUploadService.Upload(c, dataUri.Topic, dataBody); err != nil {
		response.Error(c, err.Error())
		return
	}

	response.Success(c, nil)
}
