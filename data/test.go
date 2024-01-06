package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"split-the-bill-server/presentation/dto"
	"split-the-bill-server/storage"
	"strings"
)

func main() {
	test_login()
	test_login_unsuccessfull()
}

func test_login() {

	posturl := "http://localhost:8080/api/user/register"

	body := []byte(`{
		"email": "test@mail.com",
		"password": "alek1337"
	}`)

	// fiber client
	agent := fiber.Post(posturl)
	agent.ContentType("application/json")
	agent.Body(body)

	var res dto.GeneralResponseDTO
	statusCode, resBody, _ := agent.Struct(&res)
	log.Println(statusCode)
	err := json.Unmarshal(resBody, &res)
	if err != nil {
		return
	}

	log.Infof("Start Test: %s\n", "test_login")
	if strings.Contains(res.Message, storage.UserAlreadyExistsError.Error()) {
		log.Infof("Test passed. Result: %s\n", res.Message)
	} else {
		log.Errorf("Test failed. Result: %s, Expected Result: %s\n", res.Message, storage.UserAlreadyExistsError.Error())
	}
}

func test_login_unsuccessfull() {

	posturl := "http://localhost:8080/api/user/register"

	body := []byte(`{
		"email": "testy@mail.com",
		"password": "alek1337"
	}`)

	// fiber client
	agent := fiber.Post(posturl)
	agent.ContentType("application/json")
	agent.Body(body)

	var res dto.GeneralResponseDTO
	statusCode, resBody, _ := agent.Struct(&res)
	log.Println(statusCode)
	err := json.Unmarshal(resBody, &res)
	if err != nil {
		return
	}

	log.Infof("Start Test: %s\n", "test_login_unsuccessfull")
	if strings.Contains(res.Message, storage.UserAlreadyExistsError.Error()) {
		log.Infof("Test passed. Result: %s\n", res.Message)
	} else {
		log.Errorf("Test failed. Result: %s, Expected Result: %s\n", res.Message, storage.UserAlreadyExistsError.Error())
	}
}
