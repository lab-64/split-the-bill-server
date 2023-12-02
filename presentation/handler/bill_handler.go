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

// GetByID 		gets bill by id.
//
//	@Summary	Get Bill by ID
//	@Tags		Bill
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"Bill ID"
//	@Success	200	{object}	dto.GeneralResponseDTO{data=dto.BillOutputDTO}
//	@Router		/api/bill/{id} [get]
func (h BillHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return core.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, c.Params("id"), err))
	}

	bill, err := h.billService.GetByID(id)
	if err != nil {
		return core.Error(c, fiber.StatusNotFound, fmt.Sprintf(ErrMsgBillNotFound, err))
	}

	return core.Success(c, fiber.StatusOK, SuccessMsgBillFound, bill)
}

// Create 		creates a bill.
//
//	@Summary	Create Bill
//	@Tags		Bill
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.BillInputDTO	true	"Request Body"
//	@Success	201		{object}	dto.GeneralResponseDTO{data=dto.BillOutputDTO}
//	@Router		/api/bill [post]
//
// TODO: How to handle bills without a group? Maybe add a default group which features only the owner? => how to mark such a group?
func (h BillHandler) Create(c *fiber.Ctx) error {

	var request BillInputDTO
	// parse nested bill from request body
	err := c.BodyParser(&request)
	if err != nil {
		return core.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgBillParse, err))
	}

	// create bill
	bill, err := h.billService.Create(request)
	if err != nil {
		return core.Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgBillCreate, err))
	}

	return core.Success(c, fiber.StatusCreated, SuccessMsgBillCreate, bill)
}

// AddItemToBill 		adds item to a bill.
//
//	@Summary	Add Item to Bill
//	@Tags		Bill
//	@Accept		json
//	@Produce	json
//	@Param		billId	path		string				true	"Bill ID"
//	@Param		request	body		dto.ItemCreateDTO	true	"Request Body"
//	@Success	201		{object}	dto.GeneralResponseDTO{data=dto.ItemOutputDTO}
//	@Router		/api/bill/{billId}/item [post]
func (h BillHandler) AddItemToBill(c *fiber.Ctx) error {
	// parse parameters
	billID, err := uuid.Parse(c.Params("billId"))
	if err != nil {
		return core.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, c.Params("billId"), err))
	}

	// parse request
	var request ItemCreateDTO
	if err := c.BodyParser(&request); err != nil {
		return core.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgItemParse, err))
	}

	// create item
	item, err := h.billService.AddItemToBill(billID, request)
	if err != nil {
		return core.Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgItemCreate, err))
	}

	return core.Success(c, fiber.StatusCreated, SuccessMsgItemCreate, item)
}
