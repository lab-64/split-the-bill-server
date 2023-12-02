package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"split-the-bill-server/core"
	. "split-the-bill-server/domain/service/service_inf"
	. "split-the-bill-server/presentation/dto"
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

// ChangeItem 	changes item.
//
//	@Summary	Change Item
//	@Tags		Item
//	@Accept		json
//	@Produce	json
//	@Param		id		path		string			true	"Item ID"
//	@Param		request	body		dto.ItemEditDTO	true	"Request Body"
//	@Success	200		{object}	dto.GeneralResponseDTO{data=dto.ItemOutputDTO}
//
//	@Router		/api/item/{id} [put]
func (h ItemHandler) ChangeItem(c *fiber.Ctx) error {
	// parse parameters
	itemID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return core.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, c.Params("id"), err))
	}

	// parse request
	var request ItemEditDTO
	if err := c.BodyParser(&request); err != nil {
		return core.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgItemContributorParse, err))
	}

	if request.ID != uuid.Nil && itemID != request.ID {
		return core.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParamMismatchFormat, "ID"))
	}

	// update item
	item, err := h.billService.ChangeItem(request)
	if err != nil {
		return core.Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgUpdateContributor, err))
	}

	return core.Success(c, fiber.StatusOK, SuccessMsgContributorUpdate, item)
}
