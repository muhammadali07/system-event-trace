package delivery

import "github.com/muhammadali07/service-grap-go-api/services/kafka/internal/app/usecase"

type KafkaConsumerHandler struct {
	kafkaUsecase *usecase.KafkaConsumerUsecase
}

func NewKafkaConsumerHandler(kafkaUsecase *usecase.KafkaConsumerUsecase) *KafkaConsumerHandler {
	return &KafkaConsumerHandler{kafkaUsecase: kafkaUsecase}
}

func (h *KafkaConsumerHandler) StartConsuming() {
	// Your code to start consuming messages
}
