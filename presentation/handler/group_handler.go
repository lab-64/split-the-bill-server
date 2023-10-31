package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"split-the-bill-server/core"
	"split-the-bill-server/domain/service/service_inf"
	"split-the-bill-server/presentation/dto"
)

type GroupHandler struct {
	groupService      service_inf.IGroupService
	invitationService service_inf.IInvitationService
}

func NewGroupHandler(GroupService *service_inf.IGroupService, InvitationService *service_inf.IInvitationService) *GroupHandler {
	return &GroupHandler{groupService: *GroupService, invitationService: *InvitationService}
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
		return core.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgGroupParse, err))
	}

	// validate group inputs
	// TODO: if name is empty, generate default name
	err := request.ValidateInput()
	if err != nil {
		return core.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgInputsInvalid, err))
	}

	// create group
	group, err := h.groupService.Create(request)
	if err != nil {
		return core.Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgGroupCreate, err))
	}

	// handle group invitations
	err = h.invitationService.CreateGroupInvitation(request, group.ID)
	if err != nil {
		return core.Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgGroupCreate, err))
	}

	return core.Success(c, fiber.StatusOK, SuccessMsgGroupCreate, group)
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
		return core.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}
	gid, err := uuid.Parse(id)

	if err != nil {
		return core.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, id, err))
	}
	group, err := h.groupService.GetByID(gid)

	if err != nil {
		return core.Error(c, fiber.StatusNotFound, fmt.Sprintf(ErrMsgGroupNotFound, err))
	}

	return core.Success(c, fiber.StatusOK, SuccessMsgGroupFound, group)
}
