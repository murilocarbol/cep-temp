package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/murilocarbol/cep-temp/application/client/response"
)

type ViaCepClient struct {
}

func NewViaCepClient() *ViaCepClient {
	return &ViaCepClient{}
}

type ViaCepClientInterface interface {
	GetEndereco(cep string) (string, error)
}

func (v ViaCepClient) GetEndereco(cep string) (string, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var endereco response.Endereco
	err = json.Unmarshal(body, &endereco)
	if err != nil {
		return "", err
	}

	if endereco.Erro != "" {
		return "", fmt.Errorf(endereco.Erro)
	}

	return endereco.Localidade, nil
}
