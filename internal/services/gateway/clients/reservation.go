package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/silazemli/lab2-template/internal/services/reservation"
)

type ReservationClient struct {
	client  HTTPClient
	baseURL string
}

func NewReservationClient(client HTTPClient, baseURL string) *ReservationClient {
	return &ReservationClient{
		client:  client,
		baseURL: baseURL,
	}
}

func (reservationClient *ReservationClient) GetAllHotels() ([]reservation.Hotel, error) {
	URL := fmt.Sprintf("%s/%s", reservationClient.baseURL, "hotels")
	request, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return []reservation.Hotel{}, fmt.Errorf("failed to build request: %w", err)
	}
	responce, err := reservationClient.client.Do(request)
	if err != nil {
		return []reservation.Hotel{}, fmt.Errorf("failed to make request: %w", err)
	}
	body, err := io.ReadAll(responce.Body)
	if err != nil {
		return []reservation.Hotel{}, fmt.Errorf("failed to read response body: %w", err)
	}
	defer responce.Body.Close()
	switch responce.StatusCode {
	case http.StatusOK:
		var hotels []reservation.Hotel
		if err := json.Unmarshal(body, &hotels); err != nil {
			return []reservation.Hotel{}, fmt.Errorf("failed to unmarshal response body: %w", err)
		}
		return hotels, nil
	case http.StatusInternalServerError, http.StatusNotFound, http.StatusBadRequest:
		return []reservation.Hotel{}, fmt.Errorf("server error: %w", err)
	default:
		return []reservation.Hotel{}, fmt.Errorf("unknown error: %w", err)
	}
}

func (reservationClient *ReservationClient) GetReservations(username string) ([]reservation.Reservation, error) {
	URL := fmt.Sprintf("%s/%s", reservationClient.baseURL, "reservations")
	request, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return []reservation.Reservation{}, fmt.Errorf("failed to build request: %w", err)
	}
	request.Header.Set("X-User-Name", username)
	responce, err := reservationClient.client.Do(request)
	if err != nil {
		return []reservation.Reservation{}, fmt.Errorf("failed to make request: %w", err)
	}
	body, err := io.ReadAll(responce.Body)
	if err != nil {
		return []reservation.Reservation{}, fmt.Errorf("failed to read response body: %w", err)
	}
	defer responce.Body.Close()
	switch responce.StatusCode {
	case http.StatusOK:
		var reservations []reservation.Reservation
		if err := json.Unmarshal(body, &reservations); err != nil {
			return []reservation.Reservation{}, fmt.Errorf("failed to unmarshal response body: %w", err)
		}
		return reservations, nil
	case http.StatusInternalServerError, http.StatusNotFound, http.StatusBadRequest:
		return []reservation.Reservation{}, fmt.Errorf("server error: %w", err)
	default:
		return []reservation.Reservation{}, fmt.Errorf("unknown error: %w", err)
	}
}

func (reservationClient *ReservationClient) GetReservation(reservationUID string, username string) (reservation.Reservation, error) {
	URL := fmt.Sprintf("%s/%s/%s", reservationClient.baseURL, "reservations", reservationUID)
	request, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return reservation.Reservation{}, fmt.Errorf("failed to build request: %w", err)
	}
	request.Header.Set("X-User-Name", username)
	responce, err := reservationClient.client.Do(request)
	if err != nil {
		return reservation.Reservation{}, fmt.Errorf("failed to make request: %w", err)
	}
	body, err := io.ReadAll(responce.Body)
	if err != nil {
		return reservation.Reservation{}, fmt.Errorf("failed to read response body: %w", err)
	}
	defer responce.Body.Close()
	switch responce.StatusCode {
	case http.StatusOK:
		var theReservation reservation.Reservation
		if err := json.Unmarshal(body, &theReservation); err != nil {
			return reservation.Reservation{}, fmt.Errorf("failed to unmarshal response body: %w", err)
		}
		if theReservation.Username != username {
			return reservation.Reservation{}, fmt.Errorf("access refused")
		}
		return theReservation, nil
	case http.StatusInternalServerError, http.StatusNotFound, http.StatusBadRequest:
		return reservation.Reservation{}, fmt.Errorf("server error: %w", err)
	default:
		return reservation.Reservation{}, fmt.Errorf("unknown error: %w", err)
	}
}

func (reservationClient *ReservationClient) MakeReservation(theReservation reservation.Reservation) error {
	URL := fmt.Sprintf("%s/%s", reservationClient.baseURL, "reservations")
	body, err := json.Marshal(theReservation)
	if err != nil {
		return fmt.Errorf("failed to build request body: %w", err)
	}
	request, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to build request: %w", err)
	}
	request.Header.Set("X-User_Name", theReservation.Username)
	request.Header.Set("Content-Type", "application/json")
	responce, err := reservationClient.client.Do(request)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	switch responce.StatusCode {
	case http.StatusCreated:
		return nil
	case http.StatusInternalServerError, http.StatusNotFound, http.StatusBadRequest:
		return fmt.Errorf("server error: %w", err)
	default:
		return fmt.Errorf("unknown error: %w", err)
	}
}

func (reservationClient *ReservationClient) CancelReservation(reservationUID string) error {
	URL := fmt.Sprintf("%s/%s/%s", reservationClient.baseURL, "reservations", reservationUID)
	request, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return fmt.Errorf("failed to build request: %w", err)
	}
	responce, err := reservationClient.client.Do(request)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	switch responce.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusInternalServerError, http.StatusNotFound, http.StatusBadRequest:
		return fmt.Errorf("server error: %w", err)
	default:
		return fmt.Errorf("unknown error: %w", err)
	}
}
