package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"split-the-bill-server/core"
	. "split-the-bill-server/domain/service/service_inf"
	. "split-the-bill-server/presentation/dto"
)

type BillHandler struct {
	billService  IBillService
	groupService IGroupService
}

func NewBillHandler(billService *IBillService, groupService *IGroupService) *BillHandler {
	return &BillHandler{billService: *billService, groupService: *groupService}
}

// GetByID 		func get bill by id
//
//	@Summary	Get Bill by ID
//	@Tags		Bill
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"Bill Id"
//	@Success	200	{object}	dto.GeneralResponseDTO{data=dto.BillOutputDTO}
//	@Router		/api/bill/{id} [get]
func (h BillHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return core.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}
	bid, err := uuid.Parse(id)
	if err != nil {
		return core.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, id, err))
	}
	bill, err := h.billService.GetByID(bid)
	if err != nil {
		return core.Error(c, fiber.StatusNotFound, fmt.Sprintf(ErrMsgBillNotFound, err))
	}

	return core.Success(c, fiber.StatusOK, SuccessMsgBillFound, bill)
}

// TODO: delete
/*func (h BillHandler) Create(c *fiber.Ctx) error {

	// create nested bill struct
	var items []ItemDTO
	request := BillInputDTO{
		Items: items,
	}

	// parse bill from request body
	err := c.BodyParser(&request)
	if err != nil {
		return core.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgBillParse, err))
	}

	// validate groupID
	_, err = h.groupService.GetByID(request.Group)
	if err != nil {
		return core.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgGroupNotFound, err))
	}

	bill, err := h.billService.Create(request)

	if err != nil {
		return core.Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgBillCreate, err))
	}

	return core.Success(c, fiber.StatusOK, SuccessMsgBillCreate, bill)
}*/

// Create 		func create bill
//
//	@Summary	Create Bill
//	@Tags		Bill
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.BillInputDTO	true	"Request Body"
//	@Success	200		{object}	dto.GeneralResponseDTO{data=dto.BillOutputDTO}
//	@Router		/api/bill [post]
//
// TODO: How to handle bills without a group? Maybe add a default group which features only the owner? => how to mark such a group?
// TODO: Separate bill and item handler
func (h BillHandler) Create(c *fiber.Ctx) error {

	// parse bill from request body
	var request BillInputDTO
	err := c.BodyParser(&request)
	if err != nil {
		return core.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgBillParse, err))
	}

	// validate groupID
	_, err = h.groupService.GetByID(request.Group)
	if err != nil {
		return core.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgGroupNotFound, err))
	}

	// create bill
	bill, err := h.billService.Create(request)
	if err != nil {
		return core.Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgBillCreate, err))
	}

	return core.Success(c, fiber.StatusOK, SuccessMsgBillCreate, bill)
}
