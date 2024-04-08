package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"

	"github.com/TSF/TFSClient/Business/domain"
	"github.com/TSF/TFSClient/infrastructure/config"
)

const (
	providerName = "namenode"
)

type nameNodeProvider struct {
	baseUrl string
	ip      string
	port    string
}

func NewNameNodeProvider() *nameNodeProvider {
	return &nameNodeProvider{
		baseUrl: config.GetStringPropetyBykey(fmt.Sprintf(config.BaseUrl, providerName)),
		ip:      config.GetStringPropetyBykey(fmt.Sprintf(config.IpConfig, providerName)),
		port:    config.GetStringPropetyBykey(fmt.Sprintf(config.PortConfig, providerName)),
	}
}

func (n *nameNodeProvider) SendMetadata(location string, chunksQuantity int) int {
	url := fmt.Sprintf(n.baseUrl, n.ip, n.port) + config.GetStringPropetyBykey(config.PostMetadata)

	bodyStruct := &domain.Metadata{
		Name:           path.Base(location),
		ChunksQuantity: chunksQuantity,
	}

	jsonData, err := json.Marshal(bodyStruct)
	if err != nil {
		return 400
	}

	body := bytes.NewBuffer(jsonData)

	resp, err := http.Post(url, "application/json", body)
	if err != nil {
		return 400
	}
	defer resp.Body.Close()
	return resp.StatusCode
}

func (n *nameNodeProvider) GetReplicas() ([]domain.Socket, int) {
	url := fmt.Sprintf(n.baseUrl, n.ip, n.port) + config.GetStringPropetyBykey(config.GetReplicas)
	resp, err := http.Get(url)
	if err != nil {
		return nil, 400
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 400
	}
	var response []domain.Socket
	json.Unmarshal(responseBody, &response)

	return response, resp.StatusCode
}

func (n *nameNodeProvider) DeleteMetadata(filename string) int {
	url := fmt.Sprintf(n.baseUrl, n.ip, n.port) + fmt.Sprintf(config.GetStringPropetyBykey(config.DeleteMetadata), filename)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return 400
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 400
	}
	defer resp.Body.Close()

	return resp.StatusCode
}

func (n *nameNodeProvider) GetMetadata(filename string) (*domain.MetadataResponse, int) {
	url := fmt.Sprintf(n.baseUrl, n.ip, n.port) + fmt.Sprintf(config.GetStringPropetyBykey(config.GetMetadata), filename)
	resp, err := http.Get(url)
	if err != nil {
		return nil, 400
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 400
	}
	var response domain.MetadataResponse
	json.Unmarshal(responseBody, &response)

	return &response, resp.StatusCode
}
