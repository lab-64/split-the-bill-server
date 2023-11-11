package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"split-the-bill-server/core"
	. "split-the-bill-server/domain/service/service_inf"
	"split-the-bill-server/presentation/dto"
)

// TODO: maybe delete invitationService if only one service is used.
// In our handler function we could call methods from the service directly via h.MethodName.
type InvitationHandler struct {
	invitationService IInvitationService
}

func NewInvitationHandler(invitationService *IInvitationService) *InvitationHandler {
	return &InvitationHandler{invitationService: *invitationService}
}

// GetByID returns the group invitation with the given ID.
//
//	@Summary	Get Group Invitation By ID
//	@Tags		Invitation
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"Invitation ID"
//	@Success	200	{object}	dto.GeneralResponseDTO
//	@Router		/api/invitation/{id} [get]
func (h InvitationHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return core.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		return core.Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgParseUUID, id, err))
	}
	invitation, err := h.invitationService.GetGroupInvitationByID(uid)
	if err != nil {
		return core.Error(c, fiber.StatusNotFound, fmt.Sprintf(ErrMsgUserNotFound, err))
	}
	return core.Success(c, fiber.StatusOK, SuccessMsgInvitationFound, invitation)
}

// GetAllFromUser returns all group invitations for the given user.
//
//	@Summary	Get All Group Invitations From User
//	@Tags		Invitation
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"User ID"
//	@Success	200	{object}	dto.GeneralResponseDTO
//	@Router		/api/invitation/user/{id} [get]
func (h InvitationHandler) GetAllFromUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return core.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		return core.Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgParseUUID, id, err))
	}
	invitations, err := h.invitationService.GetGroupInvitationsFromUser(uid)
	if err != nil {
		return core.Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgUserNotFound, err))
	}
	return core.Success(c, fiber.StatusOK, SuccessMsgInvitationFound, invitations)
}

// Create creates a new group invitation.
//
//	@Summary	Create Group Invitation
//	@Tags		Invitation
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.GroupInvitationDTO	true	"Request Body"
//	@Success	200		{object}	dto.GeneralResponseDTO
//	@Router		/api/invitation [post]
func (h InvitationHandler) Create(c *fiber.Ctx) error {
	// TODO: validate inputs
	// parse request
	var request dto.GroupInvitationDTO
	if err := c.BodyParser(&request); err != nil {
		return core.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgInvitationParse, err))
	}
	// create invitation for all invitees
	err := h.invitationService.CreateGroupInvitation(request)
	if err != nil {
		return core.Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgInvitationCreate, err))
	}

	return core.Success(c, fiber.StatusOK, SuccessMsgInvitationCreate, nil)
}

// Accept accepts a group invitation.
//
//	@Summary	Accept Group Invitation
//	@Tags		Invitation
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.HandleInvitationInputDTO	true	"Request Body"
//	@Success	200		{object}	dto.GeneralResponseDTO
//	@Router		/api/invitation/accept [post]
func (h InvitationHandler) Accept(c *fiber.Ctx) error {
	// parse request
	var request dto.HandleInvitationInputDTO
	if err := c.BodyParser(&request); err != nil {
		return core.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgInvitationParse, err))
	}
	// accept invitation
	err := h.invitationService.AcceptGroupInvitation(request.InvitationID, request.Issuer)
	return core.Success(c, fiber.StatusOK, SuccessMsgInvitationHandled, err)
}

// Decline declines a group invitation.
//
//	@Summary	Decline Group Invitation
//	@Tags		Invitation
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.HandleInvitationInputDTO	true	"Request Body"
//	@Success	200		{object}	dto.GeneralResponseDTO
//	@Router		/api/invitation/decline [post]
func (h InvitationHandler) Decline(c *fiber.Ctx) error {
	// parse request
	var request dto.HandleInvitationInputDTO
	if err := c.BodyParser(&request); err != nil {
		return core.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgInvitationParse, err))
	}
	err := h.invitationService.DeclineGroupInvitation(request.InvitationID, request.Issuer)
	return core.Success(c, fiber.StatusOK, SuccessMsgInvitationHandled, err)
}
