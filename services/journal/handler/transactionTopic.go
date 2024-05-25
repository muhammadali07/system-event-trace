package handler

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type messageHandler func(message kafka.Message) error

func (h *HandlerKafka) ConvertPayload(payload interface{}) (data map[string]interface{}, err error) {
	var convertedPayload map[string]interface{}
	if payload != nil {
		payloadBytes, ok := payload.([]byte)
		if !ok {
			err = fmt.Errorf("payload harus dalam format []byte")
			log.Error(err)
			return nil, err
		}
		if err := json.Unmarshal(payloadBytes, &convertedPayload); err != nil {
			err = fmt.Errorf("error saat mengkonversi payload: %v", err)
			log.Error(err)
			return nil, err
		}

	}
	return convertedPayload, nil
}

func (h *HandlerKafka) RouteTopic(topic string, payload any) (response any, err error) {
	converData, err := h.ConvertPayload(payload)
	if err != nil {
		h.log.Error(logrus.Fields{"err": err.Error()}, nil, "")
		return
	}

	// Map yang memetakan nama topik ke fungsi yang akan menanganinya
	topicHandlers := map[string]messageHandler{
		"cash_deposit": func(msg kafka.Message) error {
			return h.handleCashDepositoTrx(converData)
		},
		"cash_withdraw": func(msg kafka.Message) error {
			return h.handleCashWithDrawTrx(converData)
		},
		"transfer_kliring": func(msg kafka.Message) error {
			return h.handleTransferKliringTrx(converData)
		},
		"mutation": func(msg kafka.Message) error {
			return h.handleGetMutationTrx(converData)
		},
	}

	if handler, ok := topicHandlers[topic]; ok {
		// Memanggil fungsi handler dan mengembalikan hasilnya
		err := handler(kafka.Message{})
		if err != nil {
			fmt.Println("Error saat menjalankan handler:", err)
			return nil, err
		}
		return "Handler executed successfully", nil
	}
	return
}

func (h *HandlerKafka) handleCashDepositoTrx(data map[string]interface{}) (err error) {
	h.log.Info(logrus.Fields{"data": data}, nil, "payload")
	if data == nil {
		return fmt.Errorf("data tidak boleh nil")
	}

	err = h.app.HandleCashDeposito(data)
	if err != nil {
		return fmt.Errorf("error saat handle cash deposito: %s", err.Error())
	}
	return nil
}

func (h *HandlerKafka) handleTransferKliringTrx(data map[string]interface{}) (err error) {
	h.log.Info(logrus.Fields{"data": data}, nil, "payload")
	if data == nil {
		return fmt.Errorf("data tidak boleh nil")
	}

	err = h.app.HandleCashDeposito(data)
	if err != nil {
		return fmt.Errorf("error saat handle cash deposito: %s", err.Error())
	}
	return nil
}
func (h *HandlerKafka) handleGetMutationTrx(data map[string]interface{}) (err error) {
	h.log.Info(logrus.Fields{"data": data}, nil, "payload")
	if data == nil {
		return fmt.Errorf("data tidak boleh nil")
	}

	err = h.app.HandleCashDeposito(data)
	if err != nil {
		return fmt.Errorf("error saat handle cash deposito: %s", err.Error())
	}
	return nil
}
func (h *HandlerKafka) handleCashWithDrawTrx(data map[string]interface{}) error {
	err := h.app.HandleCashWithDraw(data)
	if err != nil {
		panic(err.Error())
	}
	return nil
}
