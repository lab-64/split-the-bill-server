package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
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

func (h InvitationHandler) GetByID(c *fiber.Ctx) error {
	return nil
}

func (h InvitationHandler) GetByUserID(c *fiber.Ctx) error {
	return nil
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
