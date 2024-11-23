package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/silazemli/lab2-template/internal/services/payment"
)

type PaymentClient struct {
	client  HTTPClient
	baseURL string
}

func NewPaymentClient(client HTTPClient, baseURL string) *PaymentClient {
	return &PaymentClient{
		client:  client,
		baseURL: baseURL,
	}
}

func (paymentClient *PaymentClient) CreatePayment(thePayment payment.Payment) error {
	URL := paymentClient.baseURL
	body, err := json.Marshal(thePayment)
	if err != nil {
		return fmt.Errorf("failed to build request body: %w", err)
	}
	request, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to build request: %w", err)
	}
	response, err := paymentClient.client.Do(request)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	switch response.StatusCode {
	case http.StatusCreated:
		return nil
	case http.StatusBadRequest, http.StatusInternalServerError:
		return fmt.Errorf("server error: %w", err)
	default:
		return fmt.Errorf("unknown error: %w", err)
	}
}

func (paymentClient *PaymentClient) CancelPayment(paymentUID string) error {
	URL := fmt.Sprintf("%s/%s", paymentClient.baseURL, paymentUID)
	request, err := http.NewRequest(http.MethodPatch, URL, nil)
	if err != nil {
		return fmt.Errorf("failed to build request: %w", err)
	}

	response, err := paymentClient.client.Do(request)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	switch response.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusBadRequest, http.StatusInternalServerError:
		return fmt.Errorf("server error: %w", err)
	default:
		return fmt.Errorf("unknown error: %w", err)
	}
}
