package modules

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

/** Session Structure */
type LeclercSession struct {
	UUID    string       `json:"uuid"`
	Product string       `json:"product"`
	Client  *http.Client `json:"client"`
}

/** Init Session */
func (s *LeclercSession) InitSession(product string) {
	u, _ := uuid.NewV4()
	s.UUID = u.String()
	s.Product = product
	s.Client = &http.Client{
		Transport: &http.Transport{},
		Timeout:   5 * time.Second,
	}
}

/** Adding To Cart */
func (s *LeclercSession) AddToCart() error {

	/** Payload for ATC */
	// Note : Les données sont en brutes et pourrait être dynamique en fonction du lien
	values := []map[string]interface{}{
		{
			"offerId":                  "2294206",
			"productSku":               "7640305958908",
			"quantity":                 1,
			"slug":                     "bottines-en-cuir-a-lacet-outdor",
			"productLogisticClassCode": "TARIF2",
			"stock":                    8,
			"isAdProduct":              true,
			"offerPrice": map[string]interface{}{
				"price": 5990,
			},
		},
	}

	payloadBytes, _ := json.Marshal(values)

	req, _ := http.NewRequest("POST", "https://www.e.leclerc/api/rest/oms-order-api/cart/compute-local-cart-from-offers", bytes.NewReader(payloadBytes))
	req.Header.Set("authority", "www.e.leclerc")
	req.Header.Set("method", "POST")
	req.Header.Set("path", "/api/rest/oms-order-api/cart/compute-local-cart-from-offers")
	req.Header.Set("scheme", "https")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("origin", "https://www.e.leclerc")
	req.Header.Set("referer", s.Product)
	req.Header.Set("sec-ch-ua", `".Not/A)Brand";v="99", "Google Chrome";v="103", "Chromium";v="103"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "macOS")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	// Send Request
	resp, err := s.Client.Do(req)
	if err != nil {
		return err
	}
	ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	switch resp.StatusCode {
	// Success return null
	case 201:
		return nil

	case 404:
		return errors.New("404")
	}

	return errors.New("Error_Untype")
}

/** Login Account */
func (s *LeclercSession) Login() error {
	return errors.New("Error_Untype")
}
