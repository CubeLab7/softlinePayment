package softlinePayment

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Service struct {
	config *Config
}

const (
	auth = "/v1/login_check"
)

func New(config *Config) *Service {
	return &Service{
		config: config,
	}
}

func (s *Service) Auth(ctx context.Context) (response *AuthResp, err error) {
	response = new(AuthResp)

	inputs := SendParams{
		Path:       auth,
		HttpMethod: http.MethodPost,
		Response:   response,
	}

	if _, err = sendRequest(s.config, inputs); err != nil {
		return
	}

	return
}

func sendRequest(config *Config, inputs SendParams) (respBody []byte, err error) {
	baseURL, err := url.Parse(config.URI)
	if err != nil {
		return respBody, fmt.Errorf("can't parse URI from config: %w", err)
	}

	// Добавляем путь из inputs.Path к базовому URL
	baseURL.Path += inputs.Path

	// Устанавливаем параметры запроса из queryParams
	query := baseURL.Query()
	for key, value := range inputs.QueryParams {
		query.Set(key, value)
	}
	baseURL.RawQuery = query.Encode()

	finalUrl := baseURL.String()

	req, err := http.NewRequest(inputs.HttpMethod, finalUrl, inputs.Body)
	if err != nil {
		return respBody, fmt.Errorf("can't create request for Softline payment system! Err: %s", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json")

	httpClient := http.Client{
		Transport: &http.Transport{
			IdleConnTimeout: time.Second * time.Duration(config.IdleConnTimeoutSec),
		},
		Timeout: time.Second * time.Duration(config.RequestTimeoutSec),
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return respBody, fmt.Errorf("can't do request! Err: %s", err)
	}
	defer resp.Body.Close()

	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return respBody, fmt.Errorf("can't read response body! Err: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return respBody, fmt.Errorf("error: %v", string(respBody))
	}

	if err = json.Unmarshal(respBody, &inputs.Response); err != nil {
		return respBody, fmt.Errorf("can't unmarshall SomPayments resp: '%v'. Err: %w", string(respBody), err)
	}

	return
}
