package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"split-the-bill-server/dto"
	"split-the-bill-server/service"
)

type BillHandler struct {
	service.IBillService
	service.IGroupService
}

func NewBillHandler(billService *service.IBillService, groupService *service.IGroupService) *BillHandler {
	return &BillHandler{IBillService: *billService, IGroupService: *groupService}
}

func (h BillHandler) Route(api fiber.Router) {
	bill := api.Group("/bill")

	bill.Get("/:id", h.GetByID)
	bill.Post("/create", h.Create)
}

func (h BillHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Parameter id is required", "data": nil})
	}
	bid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Unable to parse uuid: %s, error: %v", id, err), "data": err})
	}
	bill, err := h.IBillService.GetByID(bid)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("BillDTO not found: %v", err), "data": err})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "BillDTO found", "data": bill})
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
	request := dto.BillCreateDTO{
		Items: items,
	}

	// parse bill from request body
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not parse bill: %v", err), "data": err})
	}

	// validate groupID
	_, err = h.IGroupService.GetByID(request.Group)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Group not found: %v", err), "data": err})
	}

	bill, err := h.IBillService.Create(request)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not create bill: %v", err), "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "BillDTO created", "data": bill})
}
