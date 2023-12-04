package dto

type GeneralResponseDTO struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
