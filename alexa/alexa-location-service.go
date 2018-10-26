package alexa

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type DeviceService struct {
	URL string
	Id string
	Token string
	Endpoint string
	Client http.Client
}

func (s *DeviceService) GetDeviceLocation() (*DeviceLocationResponse, error) {
	endp := fmt.Sprintf(s.Endpoint, s.Id)
	reqUrl := fmt.Sprintf("%s%s", s.URL, endp)
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, err
	}
	bearer := "Bearer " + s.Token
	req.Header.Add("Authorization", bearer)

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	deviceLocationResponse := DeviceLocationResponse{}

	if err := json.NewDecoder(resp.Body).Decode(&deviceLocationResponse); err != nil {
		return nil, err
	}
	return &deviceLocationResponse, nil
}