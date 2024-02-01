package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"split-the-bill-server/domain/service"
	. "split-the-bill-server/presentation"
	. "split-the-bill-server/presentation/dto"
	"split-the-bill-server/presentation/middleware"
)

type InvitationHandler struct {
	invitationService service.IInvitationService
}

func NewInvitationHandler(invitationService *service.IInvitationService) *InvitationHandler {
	return &InvitationHandler{invitationService: *invitationService}
}

// GetByID returns the group invitation with the given ID.
//
//	@Summary	Get Group Invitation By ID
//	@Tags		Invitation
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"Invitation ID"
//	@Success	200	{object}	dto.GeneralResponseDTO{data=GroupInvitationOutputDTO}
//	@Router		/api/invitation/{id} [get]
func (h InvitationHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, id, err))
	}

	invitation, err := h.invitationService.GetGroupInvitationByID(uid)
	if err != nil {
		return Error(c, fiber.StatusNotFound, fmt.Sprintf(ErrMsgUserNotFound, err))
	}
	return Success(c, fiber.StatusOK, SuccessMsgInvitationFound, invitation)
}

// TODO: delete
// Create creates a new group invitation.
//
//	@Summary	Create Group Invitation
//	@Tags		Invitation
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.GroupInvitationInputDTO	true	"Request Body"
//	@Success	200		{object}	dto.GeneralResponseDTO{data=GroupInvitationOutputDTO}
//	@Router		/api/invitation [post]
func (h InvitationHandler) Create(c *fiber.Ctx) error {
	// TODO: validate inputs
	// parse request
	var request GroupInvitationInputDTO
	if err := c.BodyParser(&request); err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgInvitationParse, err))
	}
	// create invitation for all invitees
	invitations, err := h.invitationService.CreateGroupInvitations(request.GroupID)
	if err != nil {
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgInvitationCreate, err))
	}
	return Success(c, fiber.StatusCreated, SuccessMsgInvitationCreate, invitations)
}

// AcceptInvitation accepts a group invitation.
//
//	@Summary	Accept Group Invitation
//	@Tags		Invitation
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"Invitation ID"
//	@Success	200	{object}	dto.GeneralResponseDTO
//	@Router		/api/invitation/{id}/accept [post]
func (h InvitationHandler) AcceptInvitation(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}
	// parse invitationID from path
	invitationID, err := uuid.Parse(id)
	if err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, id, err))
	}
	// get authenticated requesterID from context
	requesterID := c.Locals(middleware.UserKey).(uuid.UUID)
	// accept invitation
	err = h.invitationService.AcceptGroupInvitation(invitationID, requesterID)
	if err != nil {
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgInvitationParse, err))
	}
	return Success(c, fiber.StatusOK, SuccessMsgInvitationHandled, nil)
}
