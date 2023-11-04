package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"split-the-bill-server/dto"
	"split-the-bill-server/http"
	"split-the-bill-server/service"
)

type InvitationHandler struct {
	service.IInvitationService
}

func NewInvitationHandler(invitationService *service.IInvitationService) *InvitationHandler {
	return &InvitationHandler{IInvitationService: *invitationService}
}

// GetByID returns the group invitation with the given ID.
//
//	@Summary	Get Group Invitation By ID
//	@Tags		Invitation
//	@Accept		json
//	@Produce	json
//	@Param		id		path	string	true	"Invitation ID"
//	@Success	200		{object}	dto.GeneralResponseDTO
//	@Router		/api/invitation/{id} [get]
func (h InvitationHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return http.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		return http.Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgParseUUID, id, err))
	}
	invitation, err := h.IInvitationService.GetGroupInvitationByID(uid)
	if err != nil {
		return http.Error(c, fiber.StatusNotFound, fmt.Sprintf(ErrMsgUserNotFound, err))
	}
	return http.Success(c, fiber.StatusOK, SuccessMsgInvitationFound, invitation)
}

// GetAllFromUser returns all group invitations for the given user.
//
//	@Summary	Get All Group Invitations From User
//	@Tags		Invitation
//	@Accept		json
//	@Produce	json
//	@Param		id		path	string	true	"User ID"
//	@Success	200		{object}	dto.GeneralResponseDTO
//	@Router		/api/invitation/user/{id} [get]
func (h InvitationHandler) GetAllFromUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return http.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		return http.Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgParseUUID, id, err))
	}
	invitations, err := h.IInvitationService.GetGroupInvitationsFromUser(uid)
	if err != nil {
		return http.Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgUserNotFound, err))
	}
	return http.Success(c, fiber.StatusOK, SuccessMsgInvitationFound, invitations)
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
		return http.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgInvitationParse, err))
	}
	// create invitation for all invitees
	err := h.IInvitationService.CreateGroupInvitation(request)
	if err != nil {
		return http.Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgInvitationCreate, err))
	}

	return http.Success(c, fiber.StatusOK, SuccessMsgInvitationCreate, nil)
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
		return http.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgInvitationParse, err))
	}
	// TODO: accept invitation
	err := h.IInvitationService.AcceptGroupInvitation(request.InvitationID, request.Issuer)
	return http.Success(c, fiber.StatusOK, SuccessMsgInvitationHandled, err)
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
		return http.Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgInvitationParse, err))
	}
	err := h.IInvitationService.DeclineGroupInvitation(request.InvitationID, request.Issuer)
	return http.Success(c, fiber.StatusOK, SuccessMsgInvitationHandled, err)
}
