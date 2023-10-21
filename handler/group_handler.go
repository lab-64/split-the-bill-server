package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"split-the-bill-server/dto"
	"split-the-bill-server/http"
	"split-the-bill-server/service"
)

type GroupHandler struct {
	service.IGroupService
	service.IUserService
}

func NewGroupHandler(UserService *service.IUserService, GroupService *service.IGroupService) *GroupHandler {
	return &GroupHandler{IUserService: *UserService, IGroupService: *GroupService}
}

// Create creates a new group, sets the ownerID to the authenticated user and adds it to the groupStorage.
// Authentication Required
func (h GroupHandler) Create(c *fiber.Ctx) error {

	// TODO: authenticate user
	// parse group from request body
	var request dto.GroupInputDTO
	if err := c.BodyParser(&request); err != nil {
		return http.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgGroupParse, err))
	}

	// validate group inputs
	// TODO: if name is empty, generate default name
	err := request.ValidateInput()
	if err != nil {
		return http.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgInputsInvalid, err))
	}

	group, err := h.IGroupService.Create(request)

	if err != nil {
		return http.Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgGroupCreate, err))
	}

	return http.Success(c, fiber.StatusOK, SuccessMsgGroupCreate, group)
}

// TODO: maybe delete, or add authentication and allow only query of own groups
func (h GroupHandler) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return http.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}
	gid, err := uuid.Parse(id)

	if err != nil {
		return http.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, id, err))
	}
	group, err := h.IGroupService.GetByID(gid)

	if err != nil {
		return http.Error(c, fiber.StatusNotFound, fmt.Sprintf(ErrMsgGroupNotFound, err))
	}

	return http.Success(c, fiber.StatusOK, SuccessMsgGroupFound, group)
}
