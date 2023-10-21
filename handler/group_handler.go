package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"split-the-bill-server/dto"
	"split-the-bill-server/service"
)

type GroupHandler struct {
	service.IGroupService
	service.IUserService
}

func NewGroupHandler(UserService *service.IUserService, GroupService *service.IGroupService) *GroupHandler {
	return &GroupHandler{IUserService: *UserService, IGroupService: *GroupService}
}

func (h GroupHandler) Route(api fiber.Router) {

	group := api.Group("/group")

	group.Get("/:id", h.Get)
	group.Post("/create", h.Create)
}

// Create creates a new group, sets the ownerID to the authenticated user and adds it to the groupStorage.
// Authentication Required
// TODO: Generalize error messages
func (h GroupHandler) Create(c *fiber.Ctx) error {

	// TODO: authenticate user
	// parse group from request body
	var request dto.GroupCreateDTO
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not parse group: %v", err), "data": err})
	}

	// validate group inputs
	// TODO: if name is empty, generate default name
	err := request.ValidateInput()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Inputs invalid: %v", err)})
	}

	group, err := h.IGroupService.Create(request)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not create group: %v", err), "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Group created", "data": group})
}

// TODO: maybe delete, or add authentication and allow only query of own groups
func (h GroupHandler) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Parameter id is required", "data": nil})
	}
	gid, err := uuid.Parse(id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Unable to parse uuid: %s, error: %v", id, err), "data": err})
	}
	group, err := h.IGroupService.GetByID(gid)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("GroupCreateDTO not found: %v", err), "data": err})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "GroupCreateDTO found", "data": group})
}
