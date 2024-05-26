package handler

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/sirupsen/logrus"
)

type messageHandler func(data map[string]interface{}) error

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
	// Map yang memetakan nama topik ke fungsi yang akan menanganinya
	topicHandlers := map[string]messageHandler{
		"cash_deposit": func(data map[string]interface{}) error {
			return h.handleCashDepositoTrx(data)
		},
		"cash_withdraw": func(data map[string]interface{}) error {
			return h.handleCashWithDrawTrx(data)
		},
		"transfer_kliring": func(data map[string]interface{}) error {
			return h.handleTransferKliringTrx(data)
		},
		"mutation": func(data map[string]interface{}) error {
			return h.handleGetMutationTrx(data)
		},
	}

	if handler, ok := topicHandlers[topic]; ok {
		// Konversi payload hanya jika topik valid
		converData, err := h.ConvertPayload(payload)
		if err != nil {
			h.log.Error(logrus.Fields{"err": err.Error()}, nil, "")
			return nil, err
		}

		// Memanggil fungsi handler dan mengembalikan hasilnya
		err = handler(converData)
		if err != nil {
			fmt.Println("Error saat menjalankan handler:", err)
			return nil, err
		}
		return "Handler executed successfully", nil
	} else {
		return nil, fmt.Errorf("topik tidak dikenali: %s", topic)
	}
}

func (h *HandlerKafka) handleCashDepositoTrx(data map[string]interface{}) (err error) {
	h.log.Info(logrus.Fields{"data": data}, nil, "payload")
	if data == nil {
		return fmt.Errorf("data should not be valid")
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
		return fmt.Errorf("data should not be valid")
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
		return fmt.Errorf("data should not be valid")
	}

	err = h.app.HandleCashDeposito(data)
	if err != nil {
		return fmt.Errorf("error saat handle cash deposito: %s", err.Error())
	}
	return nil
}
func (h *HandlerKafka) handleCashWithDrawTrx(data map[string]interface{}) error {
	h.log.Info(logrus.Fields{"data": data}, nil, "payload")
	if data == nil {
		return fmt.Errorf("data should not be valid")
	}

	err := h.app.HandleCashWithDraw(data)
	if err != nil {
		panic(err.Error())
	}
	return nil
}
