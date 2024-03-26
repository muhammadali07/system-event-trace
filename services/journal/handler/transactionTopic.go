package handler

import (
	"github.com/segmentio/kafka-go"
)

type messageHandler func(message kafka.Message)

func (h *HandlerKafka) RouteTopic(topic string) (response any, err error) {
	// Map yang memetakan nama topik ke fungsi yang akan menanganinya
	topicHandlers := map[string]messageHandler{
		"cash_deposit": h.handleCashDepositoTrx,
		// "cash_withdraw":    handleCashWithDrawTrx,
		// "transfer_kliring": handleTransferKliringTrx,
		// "get_list_transaction" : h.app.GetListTransaction(data),
		// "default_topic":    handleDefaultMessage,
	}

	if handler, ok := topicHandlers[topic]; ok {
		return handler, nil
	}
	return
}

func (h *HandlerKafka) handleCashDepositoTrx(message kafka.Message) {
	err := h.app.HandleCashDeposito(message)
	if err != nil {
		panic(err.Error)
	}
	// Tambahkan logika penyimpanan data ke dalam database di sini
}
