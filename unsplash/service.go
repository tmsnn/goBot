package unsplash

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type Service interface {
	GetRandomPhoto() (*Photo, error)
}
type service struct {
	client    *http.Client
	accessKey string
}

func NewService(key string) *service {
	client := &http.Client{Timeout: 30 * time.Second}
	return &service{
		client:    client,
		accessKey: key,
	}
}

func (s *service) GetRandomPhoto() (*Photo, error) {
	url := "https://api.unsplash.com/photos/random?client_id=" + s.accessKey

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	randPhoto := &Photo{}
	response, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(response, randPhoto)
	if err != nil {
		return nil, err
	}

	return randPhoto, nil
}
