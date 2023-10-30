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
	service.IInvitationService
}

func NewGroupHandler(GroupService *service.IGroupService, InvitationService *service.IInvitationService) *GroupHandler {
	return &GroupHandler{IGroupService: *GroupService, IInvitationService: *InvitationService}
}

// Create 		func create a new group & set the ownerID to the authenticated user
//
//	@Summary	Create Group
//	@Tags		Group
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.GroupInputDTO	true	"Request Body"
//	@Success	200		{object}	dto.GeneralResponseDTO{data=dto.GroupOutputDTO}
//	@Router		/api/group [post]
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

	// create group
	group, err := h.IGroupService.Create(request)
	if err != nil {
		return http.Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgGroupCreate, err))
	}

	// handle group invitations
	err = h.IInvitationService.CreateGroupInvitation(request, group.ID)
	if err != nil {
		return http.Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgGroupCreate, err))
	}

	return http.Success(c, fiber.StatusOK, SuccessMsgGroupCreate, group)
}

// Get 			func get group
//
//	@Summary	Get Group
//	@Tags		Group
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"Group Id"
//	@Success	200	{object}	dto.GeneralResponseDTO{data=dto.GroupOutputDTO}
//	@Router		/api/group/{id} [get]
//
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
