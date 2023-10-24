package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"split-the-bill-server/dto"
	"split-the-bill-server/http"
	"split-the-bill-server/service"
)

type BillHandler struct {
	service.IBillService
	service.IGroupService
}

func NewBillHandler(billService *service.IBillService, groupService *service.IGroupService) *BillHandler {
	return &BillHandler{IBillService: *billService, IGroupService: *groupService}
}

func (h BillHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return http.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}
	bid, err := uuid.Parse(id)
	if err != nil {
		return http.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, id, err))
	}
	bill, err := h.IBillService.GetByID(bid)
	if err != nil {
		return http.Error(c, fiber.StatusNotFound, fmt.Sprintf(ErrMsgBillNotFound, err))
	}

	return http.Success(c, fiber.StatusOK, SuccessMsgBillFound, bill)
}

// Create creates a new bill.
// Authentication Required
// TODO: How to handle bills without a group? Maybe add a default group which features only the owner? => how to mark such a group?
func (h BillHandler) Create(c *fiber.Ctx) error {
	// TODO: authenticate user
	/*user, err := h.getAuthenticatedUserFromHeader(c.GetReqHeaders())
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Authentication declined: %v", err)})
	}
	*/

	// create nested bill struct
	var items []dto.ItemDTO
	request := dto.BillInputDTO{
		Items: items,
	}

	// parse bill from request body
	err := c.BodyParser(&request)
	if err != nil {
		return http.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgBillParse, err))
	}

	// validate groupID
	_, err = h.IGroupService.GetByID(request.Group)
	if err != nil {
		return http.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgGroupNotFound, err))
	}

	bill, err := h.IBillService.Create(request)

	if err != nil {
		return http.Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgBillCreate, err))
	}

	return http.Success(c, fiber.StatusOK, SuccessMsgBillCreate, bill)
}
