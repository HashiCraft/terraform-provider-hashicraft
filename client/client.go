package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Bot struct {
	host string
	key  string
}

func New(host, apiKey string) *Bot {
	return &Bot{host, apiKey}
}

func (b *Bot) Health() (string, error) {
  d,err := b.doGet("/health")
  if err != nil {
    return "",err
  }

	return string(d), nil
}

// New calls the /bot POST handler to create a new bot 
func (b *Bot) New(c NewRequest) (string,error) {
  d,err := b.doPost("/bot", c)
  if err != nil {
    return "", err
  }

  nr := &NewResponse{}
  json.Unmarshal(d, nr)

  return nr.ID, nil
}

func (b*Bot) Delete(id string) error {
  _,err := b.doDelete("/bot/"+id)
  if err != nil {
    return err
  }

  return nil
}

func (b *Bot) Configure(id string, config ConfigRequest) error {
  _,err := b.doPost("/bot/"+id+"/configure", config)

	return err
}

func (b*Bot) doGet(path string) ([]byte, error) {
	r, err := http.NewRequest(http.MethodGet, b.host+path, nil)
	if err != nil {
		return nil, err
	}

  b.addHeaders(r)

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error getting health, got status %d", r.Response.StatusCode)
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (b *Bot) Start(id string) error {
  _,err := b.doGet("/bot/"+id+"/start")
  if err != nil {
    return err
  }

	return nil
}

func (b*Bot) doPost(path string, body interface{}) ([]byte, error) {
	d, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequest(http.MethodPost, b.host+path, bytes.NewBuffer(d))
	if err != nil {
		return nil, err
	}

  b.addHeaders(r)

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error configuring bot, got status %d", resp.StatusCode)
	}

  defer resp.Body.Close()
  return ioutil.ReadAll(resp.Body)
}

func (b*Bot) doDelete(path string) ([]byte, error) {
	r, err := http.NewRequest(http.MethodDelete, b.host+path, nil)
	if err != nil {
		return nil, err
	}

  b.addHeaders(r)

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNotFound {
		return nil, fmt.Errorf("Error configuring bot, got status %d", resp.StatusCode)
	}

  defer resp.Body.Close()
  return ioutil.ReadAll(resp.Body)
}

func (b*Bot) addHeaders(r *http.Request){
	r.Header.Add("authorization", "bearer "+b.key)
	r.Header.Add("content-type", "application/json")
}

type ConfigRequest struct {
	MineStart string `json:"mine_start"`
	MineEnd   string `json:"mine_end"`
	ToolChest string `json:"tool_chest"`
	DropChest string `json:"drop_chest"`
}

type NewRequest struct {
	Host string `json:"host"`
	Port   int `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewResponse struct {
	ID string `json:"id"`
	Message string `json:"message"`
}
