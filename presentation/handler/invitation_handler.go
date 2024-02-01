package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	"split-the-bill-server/domain/service"
	. "split-the-bill-server/presentation"
	. "split-the-bill-server/presentation/dto"
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

// GetAllByUser returns all group invitations for the given user.
//
//	@Summary	Get All Group Invitations From User
//	@Tags		Invitation
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"User ID"
//	@Success	200	{object}	dto.GeneralResponseDTO{data=[]GroupInvitationOutputDTO}
//	@Router		/api/invitation/user/{id} [get]
func (h InvitationHandler) GetAllByUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, id, err))
	}

	invitations, err := h.invitationService.GetGroupInvitationsByUser(uid)
	println(invitations)
	if err != nil {
		return Error(c, fiber.StatusNotFound, fmt.Sprintf(ErrMsgUserNotFound, err))
	}
	return Success(c, fiber.StatusOK, SuccessMsgInvitationFound, invitations)
}

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
	invitations, err := h.invitationService.CreateGroupInvitations(request)
	if err != nil {
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgInvitationCreate, err))
	}

	return Success(c, fiber.StatusCreated, SuccessMsgInvitationCreate, invitations)
}

// HandleInvitation handles a group invitation.
//
//	@Summary	Accept or decline Group Invitation
//	@Tags		Invitation
//	@Accept		json
//	@Produce	json
//	@Param		id		path		string							true	"Invitation ID"
//	@Param		request	body		dto.InvitationResponseInputDTO	true	"Request Body"
//	@Success	200		{object}	dto.GeneralResponseDTO
//	@Router		/api/invitation/{id}/response [post]
func (h InvitationHandler) HandleInvitation(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, id, err))
	}

	// parse request
	var request InvitationResponseInputDTO
	if err := c.BodyParser(&request); err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgInvitationParse, err))
	}

	// handle invitation
	if err := h.invitationService.HandleGroupInvitation(uid, request.IsAccept); err != nil {
		return Error(c, fiber.StatusInternalServerError, fmt.Sprintf(ErrMsgInvitationHandle, err))
	}

	return Success(c, fiber.StatusOK, SuccessMsgInvitationHandled, err)
}

// AddFriend adds a friend to the user's friend list.
//
//	@Summary	Add Friend
//	@Tags		Invitation
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.InvitationResponseInputDTO	true	"Request Body"
//	@Success	200		{object}	dto.GeneralResponseDTO
//	@Router		/api/invitation/add-user [post]
func (h InvitationHandler) AddFriend(c *fiber.Ctx) error {
	token := uuid.New().String()
	log.Println(token)

	// Construct the invitation link
	invitationLink := fmt.Sprintf("http://localhost:8080/api/invitation/get-invitation/%s", token)
	log.Println(invitationLink)

	return c.SendString("Invitation sent! Share this link with your friend: " + invitationLink)
}

// GetInvitation returns the invitation with the given ID.
//
//	@Summary	Get Invitation
//	@Tags		Invitation
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"Invitation ID"
//	@Success	200	{object}	dto.GeneralResponseDTO{data=GroupInvitationOutputDTO}
//	@Router		/api/invitation/get-invitation/{id} [get]
func (h InvitationHandler) GetInvitation(c *fiber.Ctx) error {
	log.Println("Getting an invitation...")
	id := c.Params("id")
	if id == "" {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "id"))
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParseUUID, id, err))
	}
	log.Println(uid)
	return nil
}

func (h InvitationHandler) InviteFriend(c *fiber.Ctx) error {
	log.Println("Inviting a friend...")
	// Generate a unique token
	token := uuid.New().String()
	log.Println(token)

	// Construct the invitation link
	invitationLink := fmt.Sprintf("http://localhost:8080/accept-invitation?token=%s", token)
	log.Println(invitationLink)

	return c.SendString("Invitation sent! Share this link with your friend: " + invitationLink)
}

func (h InvitationHandler) AcceptInvitation(c *fiber.Ctx) error {
	log.Println("Accepting an invitation...")
	token := c.Query("token")
	if token == "" {
		return Error(c, fiber.StatusBadRequest, fmt.Sprintf(ErrMsgParameterRequired, "token"))
	}
	return nil
}
