package whatsgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type WhatsGo struct {
	Key         string `json:"-"`      // API Key provided by whatools
	Status      string `json:"status"` // User account status: active, existing, new
	CountryCode string `json:"cc"`     // self explained, Country code
	PhoneNumber string `json:"pn"`     // self explained, phone number
}

type WhatoolsResponse struct {
	Result  WhatsGo `json:"result"`
	Success bool    `json:"success"`
	Error   string  `json:"error"`
}

const(
	Version = "2"
	Endpoint  = "https://api.wha.tools"
	URL = Endpoint + "/v" + Version + "/"
)

func Subscribe(key string) *WhatsGo {
	resp, err := http.Get(URL + "subscribe?key=" + key)
	if err != nil {
		log.Fatal("could not connect with wha.tools api")
	}

	jsonStream, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var wr WhatoolsResponse
	err = json.Unmarshal(jsonStream, &wr)
	if err != nil {
		log.Fatal("could not unmarshall wha.tools JSON response")
	}

	if wr.Success {
		whatsGo := &wr.Result
		whatsGo.Key = key
		return whatsGo
	}
	return nil
}

func (whatsGo *WhatsGo) Unsubscribe() error {
	resp, err := http.Get(URL + "unsubscribe?key=" + whatsGo.Key)
	if err != nil {
		log.Fatal("could not connect with wha.tools api")
	}

	jsonStream, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var wr map[string]interface{}
	err = json.Unmarshal(jsonStream, &wr)
	if err != nil {
		log.Fatal("could not unmarshall wha.tools JSON response")
	}

	val, ok := wr["success"]
	if ok && val.(bool) == false {
		return errors.New(fmt.Sprintf("%v", wr["error"]))
	}

	return nil
}

type Message struct {
	To    string //phone number with country code associated, e.g 573001231232
	Body  string //text to send
	Honor bool   //(true/false) This flag prevents the phone number from being formatted again if it is already in international form.
}

func (whatsGo *WhatsGo) SendMessage(messages ...*Message) {
	for _, message := range messages {
		params := "key=" + whatsGo.Key +
			"&to=" + message.To +
			"&body=" + message.Body +
			"&honor=" + fmt.Sprintf("%v", message.Honor)

		resp, err := http.Post(URL + "message", "application/x-www-form-urlencoded", strings.NewReader(params))

		if err != nil {
			log.Printf("could not send message to %s", message.To)
			continue
		}

		jsonStream, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		var wr map[string]string
		json.Unmarshal(jsonStream, &wr)
		_, ok := wr["success"]
		if ok {
			log.Printf("Message Sent to %s", message.To)
		}
	}
}

func (whatsGo *WhatsGo) GetNickname() string {
	resp, err := http.Get(URL + "nickname?key=" + whatsGo.Key)
	if err != nil {
		log.Printf("could not connect with wha.tools %s", err)
	}

	jsonStream, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var wr map[string]interface{}
	json.Unmarshal(jsonStream, &wr)

	val, ok := wr["success"]
	if ok && val.(bool) {
		return wr["result"].(string)
	}
	return ""
}

func (whatsGo *WhatsGo) SetNickname(nickname string) error {
	params := "key=" + whatsGo.Key + "&nickname=" + nickname
	resp, err := http.Post(URL + "nickname", "application/x-www-form-urlencoded", strings.NewReader(params))

	if err != nil {
		return err
	}

	jsonStream, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var wr map[string]interface{}
	json.Unmarshal(jsonStream, &wr)

	val, ok := wr["success"]
	if ok && val.(bool) {
		return nil
	}
	return errors.New("could not change nickname on whatools")
}

func (whatsGo *WhatsGo) GetStatusMessage() string {
	resp, err := http.Get(URL + "status?key=" + whatsGo.Key)
	if err != nil {
		log.Printf("could not connect with wha.tools %s", err)
	}

	jsonStream, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var wr map[string]interface{}
	json.Unmarshal(jsonStream, &wr)

	val, ok := wr["success"]
	if ok && val.(bool) {
		return fmt.Sprintf("%v", wr["result"])
	}
	return ""
}

func (whatsGo *WhatsGo) SetStatusMessage(message string) error {
	params := "key=" + whatsGo.Key + "&message=" + message
	resp, err := http.Post(URL + "status", "application/x-www-form-urlencoded", strings.NewReader(params))

	if err != nil {
		return err
	}

	if resp.StatusCode == 200 {
		jsonStream, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		var wr map[string]interface{}
		json.Unmarshal(jsonStream, &wr)

		val, ok := wr["success"]
		if ok && val.(bool) {
			return nil
		}
	}
	return errors.New("could not change nickname on whatools")
}
