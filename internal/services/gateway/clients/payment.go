package clients

import (
	"fmt"
	"net/http"
	"strconv"
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

func (paymentClient *PaymentClient) CreatePayment(price int) error {
	URL := fmt.Sprintf("%s/%s", paymentClient.baseURL, strconv.Itoa(price))
	request, err := http.NewRequest(http.MethodPost, URL, nil)
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
