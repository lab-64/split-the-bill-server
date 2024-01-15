package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"split-the-bill-server/authentication"
	"split-the-bill-server/core"
	. "split-the-bill-server/domain/service/service_inf"
	. "split-the-bill-server/presentation/dto"
)

type GroupHandler struct {
	groupService      IGroupService
	invitationService IInvitationService
}

func NewGroupHandler(GroupService *IGroupService, InvitationService *IInvitationService) *GroupHandler {
	return &GroupHandler{groupService: *GroupService, invitationService: *InvitationService}
}

// Create creates a new group with the owner being the only member.
//
//	@Summary	Create Group
//	@Tags		Group
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.GroupInputDTO	true	"Request Body"
//	@Success	200		{object}	dto.GeneralResponseDTO{data=dto.GroupDetailedOutputDTO}
//	@Router		/api/group [post]
func (h GroupHandler) Create(c *fiber.Ctx) error {

	// TODO: authenticate user
	// parse group from request body
	var request GroupInputDTO
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

	return core.Success(c, fiber.StatusCreated, SuccessMsgGroupCreate, group)
}

// Update updates a group with the given id.
//
//	@Summary	Update Group
//	@Tags		Group
//	@Accept		json
//	@Produce	json
//	@Param		id		path		string				true	"Group ID"
//	@Param		request	body		dto.GroupInputDTO	true	"Request Body"
//	@Success	200		{object}	dto.GeneralResponseDTO{data=dto.GroupDetailedOutputDTO}
//
//	@Router		/api/group/{id} [put]
func (g GroupHandler) Update(c *fiber.Ctx) error {
	// parse parameters
	id := c.Params("id")
	if id == "" {
		return core.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return core.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, uid, err))
	}

	// parse request
	var request GroupInputDTO
	if err := c.BodyParser(&request); err != nil {
		return core.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgGroupParse, err))
	}

	userID := c.Locals(authentication.UserKey).(uuid.UUID)

	// update item
	item, err := g.groupService.Update(userID, uid, request)
	if err != nil {
		return core.Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgGroupUpdate, err))
	}

	return core.Success(c, fiber.StatusOK, SuccessMsgGroupUpdate, item)
}

// GetByID returns the group with the given ID.
//
//	@Summary	Get Group by ID
//	@Tags		Group
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"Group Id"
//	@Success	200	{object}	dto.GeneralResponseDTO{data=dto.GroupDetailedOutputDTO}
//	@Router		/api/group/{id} [get]
//
// TODO: maybe delete, or add authentication and allow only query of own groups
func (h GroupHandler) GetByID(c *fiber.Ctx) error {
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

// GetAllByUser returns all groups filtered by user.
//
//	@Summary	Get Groups by User
//	@Tags		Group
//	@Accept		json
//	@Produce	json
//	@Param		userId	query		string	true	"User Id"
//	@Success	200		{object}	dto.GeneralResponseDTO{data=dto.GroupDetailedOutputDTO}
//	@Router		/api/group [get]
func (h GroupHandler) GetAllByUser(c *fiber.Ctx) error {
	userID := c.Query("userId")

	if userID == "" {
		return core.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "userId"))
	}

	uid, err := uuid.Parse(userID)
	if err != nil {
		return core.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, userID, err))
	}

	groups, err := h.groupService.GetAllByUser(uid)

	if err != nil {
		return core.Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgGetUserGroups, err))
	}

	return core.Success(c, fiber.StatusOK, SuccessMsgGroupsFound, groups)

}
