package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"split-the-bill-server/domain/service"
	. "split-the-bill-server/presentation"
	. "split-the-bill-server/presentation/dto"
	"split-the-bill-server/presentation/middleware"
)

type BillHandler struct {
	billService  service.IBillService
	groupService service.IGroupService
}

func NewBillHandler(billService *service.IBillService, groupService *service.IGroupService) *BillHandler {
	return &BillHandler{billService: *billService, groupService: *groupService}
}

// GetByID 		gets bill by id.
//
//	@Summary	Get Bill by ID
//	@Tags		Bill
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"Bill ID"
//	@Success	200	{object}	dto.GeneralResponseDTO{data=dto.BillDetailedOutputDTO}
//	@Router		/api/bill/{id} [get]
func (h BillHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, uid, err))
	}

	bill, err := h.billService.GetByID(uid)
	if err != nil {
		return Error(c, fiber.StatusNotFound, fmt.Sprintf(ErrMsgBillNotFound, err))
	}

	return Success(c, fiber.StatusOK, SuccessMsgBillFound, bill)
}

// Create 		creates a bill.
//
//	@Summary	Create Bill
//	@Tags		Bill
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.BillInputDTO	true	"Request Body"
//	@Success	201		{object}	dto.GeneralResponseDTO{data=dto.BillDetailedOutputDTO}
//	@Router		/api/bill [post]
//
// TODO: How to handle bills without a group? Maybe add a default group which features only the owner? => how to mark such a group?
func (h BillHandler) Create(c *fiber.Ctx) error {

	var request BillInputDTO
	// parse nested bill from request body
	err := c.BodyParser(&request)
	if err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgBillParse, err))
	}

	// create bill
	bill, err := h.billService.Create(request)
	if err != nil {
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgBillCreate, err))
	}

	return Success(c, fiber.StatusCreated, SuccessMsgBillCreate, bill)
}

// Update updates a bill with the given id.
//
//	@Summary	Update Bill
//	@Tags		Bill
//	@Accept		json
//	@Produce	json
//	@Param		id		path		string				true	"Bill ID"
//	@Param		request	body		dto.BillInputDTO	true	"Request Body"
//	@Success	200		{object}	dto.GeneralResponseDTO{data=dto.BillDetailedOutputDTO}
//
//	@Router		/api/bill/{id} [put]
func (g BillHandler) Update(c *fiber.Ctx) error {
	// parse parameters
	id := c.Params("id")
	if id == "" {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, uid, err))
	}

	// parse request
	var request BillInputDTO
	if err := c.BodyParser(&request); err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgBillParse, err))
	}

	userID := c.Locals(middleware.UserKey).(uuid.UUID)

	// update item
	item, err := g.billService.Update(userID, uid, request)
	if err != nil {
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgBillUpdate, err))
	}

	return Success(c, fiber.StatusOK, SuccessMsgBillUpdate, item)
}

// GetAllByUser 	gets all bills by user id.
//
//	@Summary	Get All Bills by User ID
//	@Tags		Bill
//	@Accept		json
//	@Produce	json
//	@Param		userId	query		string	true	"User ID"
//	@Success	200		{object}	dto.GeneralResponseDTO{data=[]dto.BillDetailedOutputDTO}
//	@Router		/api/bill [get]
func (h BillHandler) GetAllByUser(c *fiber.Ctx) error {
	userID := c.Query("userId")
	if userID == "" {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "userId"))
	}
	uid, err := uuid.Parse(userID)
	if err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, uid, err))
	}

	requesterID := c.Locals(middleware.UserKey).(uuid.UUID)

	bills, err := h.billService.GetAllByUserID(uid, requesterID)
	if err != nil {
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgBillGetAll, err))
	}

	return Success(c, fiber.StatusOK, SuccessMsgBillGetAll, bills)
}

// AddItem 		adds item to a bill.
//
//	@Summary	Add Item to Bill
//	@Tags		Bill
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.ItemInputDTO	true	"Request Body"
//	@Success	201		{object}	dto.GeneralResponseDTO{data=dto.ItemOutputDTO}
//	@Router		/api/bill/item [post]
func (h BillHandler) AddItem(c *fiber.Ctx) error {
	// parse request
	var request ItemInputDTO
	if err := c.BodyParser(&request); err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgItemParse, err))
	}

	// create item
	item, err := h.billService.AddItem(request)
	if err != nil {
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgItemCreate, err))
	}

	return Success(c, fiber.StatusCreated, SuccessMsgItemCreate, item)
}

// GetItemByID 	 gets item by ID.
//
//	@Summary	Get Item by ID
//	@Tags		Bill
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"Item ID"
//	@Success	200	{object}	dto.GeneralResponseDTO{data=dto.ItemOutputDTO}
//	@Router		/api/bill/item/{id} [get]
func (h BillHandler) GetItemByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, uid, err))
	}

	item, err := h.billService.GetItemByID(uid)
	if err != nil {
		return Error(c, fiber.StatusNotFound, fmt.Sprintf(ErrMsgItemNotFound, err))
	}

	return Success(c, fiber.StatusOK, SuccessMsgItemFound, item)
}

// ChangeItem 	changes item.
//
//	@Summary	Change Item
//	@Tags		Bill
//	@Accept		json
//	@Produce	json
//	@Param		id		path		string				true	"Item ID"
//	@Param		request	body		dto.ItemInputDTO	true	"Request Body"
//	@Success	200		{object}	dto.GeneralResponseDTO{data=dto.ItemOutputDTO}
//
//	@Router		/api/bill/item/{id} [put]
func (h BillHandler) ChangeItem(c *fiber.Ctx) error {
	// parse parameters
	id := c.Params("id")
	if id == "" {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, uid, err))
	}

	// parse request
	var request ItemInputDTO
	if err := c.BodyParser(&request); err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgItemParse, err))
	}

	// update item
	item, err := h.billService.ChangeItem(uid, request)
	if err != nil {
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgItemUpdate, err))
	}

	return Success(c, fiber.StatusOK, SuccessMsgItemUpdate, item)
}
