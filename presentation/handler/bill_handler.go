package handler

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"split-the-bill-server/domain"
	"split-the-bill-server/domain/service"
	. "split-the-bill-server/presentation"
	"split-the-bill-server/presentation/dto"
	"split-the-bill-server/presentation/middleware"
	"strconv"
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
//	@Success	200	{object}	dto.GeneralResponse{data=dto.BillDetailedOutput}
//	@Router		/api/bill/{id} [get]
func (h BillHandler) GetByID(c *fiber.Ctx) error {
	// parse parameter
	id := c.Params("id")
	if id == "" {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, uid, err))
	}
	// get authenticated requester from context
	requesterID := c.Locals(middleware.UserKey).(uuid.UUID)
	// get bill
	bill, err := h.billService.GetByID(requesterID, uid)
	if err != nil {
		if errors.Is(err, domain.ErrNotAuthorized) {
			return Error(c, fiber.StatusUnauthorized, fmt.Sprintf(ErrMsgBillNotFound, err))
		}
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgBillNotFound, err))
	}

	return Success(c, fiber.StatusOK, SuccessMsgBillFound, bill)
}

// Create 		creates a bill.
//
//	@Summary	Create Bill
//	@Tags		Bill
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.BillCreate	true	"Request Body"
//	@Success	201		{object}	dto.GeneralResponse{data=dto.BillDetailedOutput}
//	@Router		/api/bill [post]
func (h BillHandler) Create(c *fiber.Ctx) error {
	// parse bill from request
	var request dto.BillCreate
	err := c.BodyParser(&request)
	if err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgBillParse, err))
	}
	// validate inputs
	if err = request.ValidateInputs(); err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgInputsInvalid, err))
	}
	// get authenticated requester from context
	requesterID := c.Locals(middleware.UserKey).(uuid.UUID)
	// create bill
	bill, err := h.billService.Create(requesterID, request)
	if err != nil {
		if errors.Is(err, domain.ErrNotAuthorized) {
			return Error(c, fiber.StatusUnauthorized, fmt.Sprintf(ErrMsgBillCreate, err))
		}
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
//	@Param		id		path		string			true	"Bill ID"
//	@Param		request	body		dto.BillUpdate	true	"Request Body"
//	@Success	200		{object}	dto.GeneralResponse{data=dto.BillDetailedOutput}
//	@Router		/api/bill/{id} [put]
func (h BillHandler) Update(c *fiber.Ctx) error {
	// parse parameter
	id := c.Params("id")
	if id == "" {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, uid, err))
	}
	// parse request
	var request dto.BillUpdate
	if err = c.BodyParser(&request); err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgBillParse, err))
	}
	// get authenticated requester from context
	requesterID := c.Locals(middleware.UserKey).(uuid.UUID)
	// update item
	item, err := h.billService.Update(requesterID, uid, request)
	if err != nil {
		if errors.Is(err, domain.ErrNotAuthorized) {
			return Error(c, fiber.StatusUnauthorized, fmt.Sprintf(ErrMsgBillUpdate, err))
		}
		if errors.Is(err, domain.ErrConcurrentModification) {
			return Error(c, fiber.StatusConflict, fmt.Sprintf(ErrMsgBillUpdate, err))
		}
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgBillUpdate, err))
	}

	return Success(c, fiber.StatusOK, SuccessMsgBillUpdate, item)
}

// Delete 		deletes a bill with the given id.
//
//	@Summary	Delete Bill
//	@Tags		Bill
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"Bill ID"
//	@Success	200	{object}	dto.GeneralResponse
//	@Router		/api/bill/{id} [delete]
func (h BillHandler) Delete(c *fiber.Ctx) error {
	// parse parameter
	id := c.Params("id")
	if id == "" {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, uid, err))
	}
	// get authenticated requester from context
	requesterID := c.Locals(middleware.UserKey).(uuid.UUID)
	// delete bill
	err = h.billService.Delete(requesterID, uid)
	if err != nil {
		if errors.Is(err, domain.ErrNotAuthorized) {
			return Error(c, fiber.StatusUnauthorized, fmt.Sprintf(ErrMsgBillDelete, err))
		}
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgBillDelete, err))
	}

	return Success(c, fiber.StatusOK, SuccessMsgBillDelete, nil)
}

// GetAllByUser 	gets all bills by user.
//
//	@Summary	Get All Bills by User
//	@Tags		Bill
//	@Accept		json
//	@Produce	json
//	@Param		userId		query		string	true	"User ID"
//	@Param		isUnseen	query		bool	false	"Is Unseen"
//	@Param		isOwner		query		bool	false	"Is Owner"
//	@Success	200			{object}	dto.GeneralResponse{data=[]dto.BillDetailedOutput}
//	@Router		/api/bill [get]
func (h BillHandler) GetAllByUser(c *fiber.Ctx) error {
	// parse query parameters
	userID := c.Query("userId")
	if userID == "" {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "userId"))
	}
	uid, err := uuid.Parse(userID)
	if err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, uid, err))
	}

	var isUnseen *bool
	if isUnseenParam, err := strconv.ParseBool(c.Query("isUnseen")); err == nil {
		isUnseen = &isUnseenParam
	}

	var isOwner *bool
	if isOwnerParam, err := strconv.ParseBool(c.Query("isOwner")); err == nil {
		isOwner = &isOwnerParam
	}
	// get authenticated requester from context
	requesterID := c.Locals(middleware.UserKey).(uuid.UUID)
	// get bills according to filter
	bills, err := h.billService.GetAllByUserID(requesterID, uid, isUnseen, isOwner)
	if err != nil {
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgBillGetAll, err))
	}

	return Success(c, fiber.StatusOK, SuccessMsgBillGetAll, bills)
}

// UpdateContribution 	updates the item contribution of the requester to the bill with the given id.
//
//	@Summary	Update Item Contribution
//	@Tags		Bill
//	@Accept		json
//	@Produce	json
//	@Param		id		path		string					true	"Bill ID"
//	@Param		request	body		dto.ContributionInput	true	"Request Body"
//	@Success	200		{object}	dto.GeneralResponse
//	@Router		/api/bill/{id}/contribution [put]
func (h BillHandler) UpdateContribution(c *fiber.Ctx) error {
	// parse parameter
	id := c.Params("id")
	if id == "" {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, uid, err))
	}
	// parse request
	var request dto.ContributionInput
	if err = c.BodyParser(&request); err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgBillParse, err))
	}
	// get authenticated requester from context
	requesterID := c.Locals(middleware.UserKey).(uuid.UUID)
	// update item contribution
	err = h.billService.HandleContribution(requesterID, uid, request)
	if err != nil {
		if errors.Is(err, domain.ErrNotAuthorized) {
			return Error(c, fiber.StatusUnauthorized, fmt.Sprintf(ErrMsgBillUpdate, err))
		}
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgBillUpdate, err))
	}
	return Success(c, fiber.StatusOK, SuccessMsgBillUpdate, nil)
}
