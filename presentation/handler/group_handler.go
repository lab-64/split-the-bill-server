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
)

type GroupHandler struct {
	groupService service.IGroupService
}

func NewGroupHandler(GroupService *service.IGroupService) *GroupHandler {
	return &GroupHandler{groupService: *GroupService}
}

// Create creates a new group with the owner being the only member.
//
//	@Summary	Create Group
//	@Tags		Group
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.GroupInput	true	"Request Body"
//	@Success	200		{object}	dto.GeneralResponse{data=dto.GroupDetailedOutput}
//	@Router		/api/group [post]
func (h GroupHandler) Create(c *fiber.Ctx) error {
	// parse group from request body
	var request dto.GroupInput
	if err := c.BodyParser(&request); err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgGroupParse, err))
	}
	// validate inputs
	err := request.ValidateInput()
	if err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgInputsInvalid, err))
	}
	// get requesterID from context
	requesterID := c.Locals(middleware.UserKey).(uuid.UUID)
	// create group
	group, err := h.groupService.Create(requesterID, request)
	if err != nil {
		if errors.Is(err, domain.ErrNotAuthorized) {
			return Error(c, fiber.StatusUnauthorized, fmt.Sprintf(ErrMsgGroupCreate, err))
		}
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgGroupCreate, err))
	}

	return Success(c, fiber.StatusCreated, SuccessMsgGroupCreate, group)
}

// Update updates a group with the given id.
//
//	@Summary	Update Group
//	@Tags		Group
//	@Accept		json
//	@Produce	json
//	@Param		id		path		string			true	"Group ID"
//	@Param		request	body		dto.GroupInput	true	"Request Body"
//	@Success	200		{object}	dto.GeneralResponse{data=dto.GroupDetailedOutput}
//
//	@Router		/api/group/{id} [put]
func (g GroupHandler) Update(c *fiber.Ctx) error {
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
	var request dto.GroupInput
	if err = c.BodyParser(&request); err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgGroupParse, err))
	}
	// validate inputs
	err = request.ValidateInput()
	if err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgInputsInvalid, err))
	}
	// get requesterID from context
	requesterID := c.Locals(middleware.UserKey).(uuid.UUID)
	// update group
	group, err := g.groupService.Update(requesterID, uid, request)
	if err != nil {
		if errors.Is(err, domain.ErrNotAuthorized) {
			return Error(c, fiber.StatusUnauthorized, fmt.Sprintf(ErrMsgGroupUpdate, err))
		}
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgGroupUpdate, err))
	}

	return Success(c, fiber.StatusOK, SuccessMsgGroupUpdate, group)
}

// GetByID returns the group with the given ID.
//
//	@Summary	Get Group by ID
//	@Tags		Group
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"Group Id"
//	@Success	200	{object}	dto.GeneralResponse{data=dto.GroupDetailedOutput}
//	@Router		/api/group/{id} [get]
func (h GroupHandler) GetByID(c *fiber.Ctx) error {
	// parse parameter
	id := c.Params("id")
	if id == "" {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}
	gid, err := uuid.Parse(id)
	if err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, id, err))
	}
	// get requesterID from context
	requesterID := c.Locals(middleware.UserKey).(uuid.UUID)
	// get group
	group, err := h.groupService.GetByID(requesterID, gid)
	if err != nil {
		if errors.Is(err, domain.ErrNotAuthorized) {
			return Error(c, fiber.StatusUnauthorized, fmt.Sprintf(ErrMsgGroupNotFound, err))
		}
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgGroupNotFound, err))
	}

	return Success(c, fiber.StatusOK, SuccessMsgGroupFound, group)
}

// GetAll returns all groups with applied filter.
//
//	@Summary	Get Groups by User/Invitation
//	@Tags		Group
//	@Accept		json
//	@Produce	json
//	@Param		userId			query		string	false	"User ID"
//	@Param		invitationId	query		string	false	"Invitation ID"
//	@Success	200				{object}	dto.GeneralResponse{data=dto.GroupDetailedOutput}
//	@Router		/api/group [get]
func (h GroupHandler) GetAll(c *fiber.Ctx) error {
	// parse query parameters
	userID := c.Query("userId")
	invitationID := c.Query("invitationId")
	userUUID := uuid.Nil
	invitationUUID := uuid.Nil
	if userID != "" {
		uid, err := uuid.Parse(userID)
		if err != nil {
			return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, userID, err))
		}
		userUUID = uid
	}
	if invitationID != "" {
		uid, err := uuid.Parse(invitationID)
		if err != nil {
			return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, userID, err))
		}
		invitationUUID = uid
	}
	// get requesterID from context
	requesterID := c.Locals(middleware.UserKey).(uuid.UUID)
	// get groups
	groups, err := h.groupService.GetAll(requesterID, userUUID, invitationUUID)
	if err != nil {
		if errors.Is(err, domain.ErrNotAuthorized) {
			return Error(c, fiber.StatusUnauthorized, fmt.Sprintf(ErrMsgGetUserGroups, err))
		}
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgGetUserGroups, err))
	}

	return Success(c, fiber.StatusOK, SuccessMsgGroupsFound, groups)
}

// Delete deletes the group with the given id.
//
//	@Summary	Delete Group
//	@Tags		Group
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"Group ID"
//	@Success	200	{object}	dto.GeneralResponse
//	@Router		/api/group/{id} [delete]
func (h GroupHandler) Delete(c *fiber.Ctx) error {
	// parse parameter
	id := c.Params("id")
	if id == "" {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, uid, err))
	}
	// get requesterID from context
	requesterID := c.Locals(middleware.UserKey).(uuid.UUID)
	// delete group
	err = h.groupService.Delete(requesterID, uid)
	if err != nil {
		if errors.Is(err, domain.ErrNotAuthorized) {
			return Error(c, fiber.StatusUnauthorized, fmt.Sprintf(ErrMsgGroupDelete, err))
		}
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgGroupDelete, err))
	}
	return Success(c, fiber.StatusOK, SuccessMsgGroupDelete, nil)
}

// AcceptInvitation accepts a group invitation.
//
//	@Summary	Accept Group Invitation
//	@Tags		Group
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"Invitation ID"
//	@Success	200	{object}	dto.GeneralResponse
//	@Router		/api/group/invitation/{id}/accept [post]
func (h GroupHandler) AcceptInvitation(c *fiber.Ctx) error {
	// parse parameter
	id := c.Params("id")
	if id == "" {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}
	invitationID, err := uuid.Parse(id)
	if err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, id, err))
	}
	// get authenticated requesterID from context
	requesterID := c.Locals(middleware.UserKey).(uuid.UUID)
	// accept invitation
	err = h.groupService.AcceptGroupInvitation(invitationID, requesterID)
	if err != nil {
		if errors.Is(err, domain.ErrNotAuthorized) {
			return Error(c, fiber.StatusUnauthorized, fmt.Sprintf(ErrMsgInvitationHandle, err))
		}
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgInvitationHandle, err))
	}

	return Success(c, fiber.StatusOK, SuccessMsgInvitationHandled, nil)
}

// CreateGroupTransaction creates a group transaction and clears the group
//
//	@Summary	Create Group Transaction & Clear Group
//	@Tags		Group
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"Group ID"
//	@Success	200	{object}	dto.GeneralResponse{data=dto.GroupTransactionOutput}
//	@Router		/api/group/{id}/transaction [post]
func (h GroupHandler) CreateGroupTransaction(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}
	groupID, err := uuid.Parse(id)
	if err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, id, err))
	}

	// get authenticated requesterID from context
	requesterID := c.Locals(middleware.UserKey).(uuid.UUID)

	transaction, err := h.groupService.CreateGroupTransaction(requesterID, groupID)

	if err != nil {
		if errors.Is(err, domain.ErrNotAuthorized) {
			return Error(c, fiber.StatusUnauthorized, fmt.Sprintf(ErrMsgGroupTransactionCreate, err))
		}
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgGroupTransactionCreate, err))
	}

	return Success(c, fiber.StatusCreated, SuccessMsgGroupTransactionCreate, transaction)
}

// GetAllGroupTransactions returns all group transactions with applied filter.
//
//	@Summary	Get Group Transactions by User
//	@Tags		Group
//	@Accept		json
//	@Produce	json
//	@Param		userId	query		string	true	"User ID"
//	@Success	200		{object}	dto.GeneralResponse{data=[]dto.GroupTransactionOutput}
//	@Router		/api/group/transaction [get]
func (h GroupHandler) GetAllGroupTransactions(c *fiber.Ctx) error {
	// parse query parameters
	userID := c.Query("userId")
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, userID, err))
	}
	// get authenticated requesterID from context
	requesterID := c.Locals(middleware.UserKey).(uuid.UUID)

	transactions, err := h.groupService.GetAllGroupTransactions(requesterID, userUUID)
	if err != nil {
		if errors.Is(err, domain.ErrNotAuthorized) {
			return Error(c, fiber.StatusUnauthorized, fmt.Sprintf(ErrMsgGetUserTransactions, err))
		}
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgGetUserTransactions, err))
	}

	return Success(c, fiber.StatusOK, SuccessMsgGroupTransactionFound, transactions)
}
