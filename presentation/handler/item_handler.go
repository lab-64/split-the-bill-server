package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"split-the-bill-server/core"
	. "split-the-bill-server/domain/service/service_inf"
)

type ItemHandler struct {
	billService IBillService
}

func NewItemHandler(billService *IBillService) *ItemHandler {
	return &ItemHandler{billService: *billService}
}

// GetByID 	 gets item by ID.
//
//	@Summary	Get Item by ID
//	@Tags		Item
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"Item ID"
//	@Success	200	{object}	dto.GeneralResponseDTO{data=dto.ItemOutputDTO}
//	@Router		/api/item/{id} [get]
func (h ItemHandler) GetByID(c *fiber.Ctx) error {
	itemID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return core.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, c.Params("id"), err))
	}

	item, err := h.billService.GetItemByID(itemID)
	if err != nil {
		return core.Error(c, fiber.StatusNotFound, fmt.Sprintf(ErrMsgItemNotFound, err))
	}

	return core.Success(c, fiber.StatusOK, SuccesMsgItemFound, item)
}
