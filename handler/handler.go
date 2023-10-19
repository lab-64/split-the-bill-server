package handler

import (
	"errors"
	"fmt"
	"github.com/caitlinelfring/nist-password-validator/password"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	"split-the-bill-server/authentication"
	"split-the-bill-server/storage"
	"split-the-bill-server/types"
	"split-the-bill-server/wire"
)

// TODO: Sanitize output / errors
// TODO: Add handler tests

type Handler struct {
	userStorage       storage.UserStorage
	cookieStorage     storage.CookieStorage
	groupStorage      storage.GroupStorage
	billStorage       storage.BillStorage
	PasswordValidator *password.Validator
}

func NewHandler(userStorage storage.UserStorage, cookieStorage storage.CookieStorage, groupStorage storage.GroupStorage, billStorage storage.BillStorage, v *password.Validator) Handler {
	return Handler{userStorage: userStorage, cookieStorage: cookieStorage, groupStorage: groupStorage, billStorage: billStorage, PasswordValidator: v}
}

// RegisterUser parses a types.User from the request body, compares and validates both passwords and adds a new user to the userStorage.
func (h Handler) RegisterUser(c *fiber.Ctx) error {
	var rUser wire.RegisterUser
	if err := c.BodyParser(&rUser); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not parse rUser: %v", err), "data": err})
	}

	err := h.PasswordValidator.ValidatePassword(rUser.Password)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Bad Password: %v", err)})
	}

	user := rUser.ToUser()
	pHash, err := authentication.HashPassword(rUser.Password)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Could not create user"})
	}

	err = h.userStorage.RegisterUser(user, pHash)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not create User: %v", err)})
	}

	return c.Status(200).JSON(fiber.Map{"status": "ok", "message": "User successfully created", "User": user.Username})
}

// Login uses the given login credentials for login and returns an authentication token for the user.
func (h Handler) Login(c *fiber.Ctx) error {
	var userCredentials wire.Credentials
	if err := c.BodyParser(&userCredentials); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not parse user: %v", err), "data": err})
	}
	// Checks if all input fields are filled out
	err := userCredentials.ValidateInputs()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Inputs invalid: %v", err)})
	}
	// Log-in user, get authentication cookie
	user, err := h.userStorage.GetUserByUsername(userCredentials.Username)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not log in: %v", err)})
	}
	creds, err := h.userStorage.GetCredentials(user.ID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not log in: %v", err)})
	}
	err = authentication.ComparePassword(creds, userCredentials.Password)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not log in: %v", err)})
	}

	sc := authentication.GenerateSessionCookie(user.ID)

	h.cookieStorage.AddAuthenticationCookie(sc)

	// Create response cookie
	// TODO: add Secure flag after development (cookie will only be sent over HTTPS)
	c.Cookie(&fiber.Cookie{
		Name:     "session",
		Value:    sc.Token.String(),
		Expires:  sc.ValidBefore,
		HTTPOnly: true,
		//Secure:   true,
	})
	return c.Status(200).JSON(fiber.Map{"status": "ok"})
}

// CreateUser parses a types.User from the request body and adds it to the userStorage.
func (h Handler) CreateUser(c *fiber.Ctx) error {
	log.Printf("CreateUser")
	// Store the body in the user and return error if encountered
	var user types.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not parse user: %v", err), "data": err})
	}
	user.ID = uuid.New()
	// Add user to userStorage.
	err := h.userStorage.AddUser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not create user: %v", err), "data": err})
	}
	// Return the created user
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User has been created", "data": user})
}

func (h Handler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.userStorage.GetAllUsers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not get users: %v", err), "data": err})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "users Found", "data": users})
}

// GetUserByID from db
func (h Handler) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Parameter id is required", "data": nil})
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Unable to parse uuid: %s, error: %v", id, err), "data": err})
	}
	user, err := h.userStorage.GetUserByID(uid)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("User not found: %v", err), "data": err})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User found", "data": user})
}

func (h Handler) DeleteUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "parameter id is required", "data": nil})
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Unable to parse uuid: %s, error: %v", id, err), "data": err})
	}
	err = h.userStorage.DeleteUser(uid)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Unable to delete user: %v", err), "data": err})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User deleted"})
}

func (h Handler) GetUserByUsername(c *fiber.Ctx) error {
	username := c.Params("username")
	if username == "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Parameter username is required", "data": nil})
	}
	user, err := h.userStorage.GetUserByUsername(username)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("User not found: %v", err), "data": err})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User found", "data": user})
}

// CreateBill creates a new bill and adds it to the billStorage.
// Authentication Required
// TODO: How to handle bills without a group? Maybe add a default group which features only the owner? => how to mark such a group?
func (h Handler) CreateBill(c *fiber.Ctx) error {
	// TODO: authenticate user
	/*user, err := h.getAuthenticatedUserFromHeader(c.GetReqHeaders())
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Authentication declined: %v", err)})
	}
	*/
	// TODO: delete if authentication is used
	userID := uuid.MustParse("7f1b2ed5-1201-4443-b997-56877fe31991")
	// create nested bill struct
	var items []wire.Item
	rBill := wire.Bill{
		Items: &items,
	}
	// parse bill from request body
	err := c.BodyParser(&rBill)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not parse bill: %v", err), "data": err})
	}
	// TODO: delete
	// get group from groupStorage
	/*group, err := h.groupStorage.GetGroupByID(rBill.Group)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not get group: %v", err), "data": err})
	}
	*/

	// create types.bill
	bill, err := rBill.ToBill(userID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not create bill: %v", err), "data": err})
	}
	// validate groupID
	_, err = h.groupStorage.GetGroupByID(rBill.Group)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Group not found: %v", err), "data": err})
	}
	// Inputs valid:
	// store bill in billStorage
	err = h.billStorage.AddBill(bill)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not create bill: %v", err), "data": err})
	}
	// add bill to group
	err = h.groupStorage.AddBillToGroup(&bill, rBill.Group)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not add bill to group: %v", err), "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Bill created", "data": bill})
}

func (h Handler) GetBillByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Parameter id is required", "data": nil})
	}
	bid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Unable to parse uuid: %s, error: %v", id, err), "data": err})
	}
	bill, err := h.billStorage.GetBillByID(bid)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Bill not found: %v", err), "data": err})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Bill found", "data": bill})
}

// CreateGroup creates a new group, sets the ownerID to the authenticated user and adds it to the groupStorage.
// Authentication Required
// TODO: Generalize error messages
func (h Handler) CreateGroup(c *fiber.Ctx) error {
	// TODO: authenticate user
	// parse group from request body
	var rGroup wire.Group
	if err := c.BodyParser(&rGroup); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not parse group: %v", err), "data": err})
	}
	// validate group inputs
	// TODO: if name is empty, generate default name
	err := rGroup.ValidateInput()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Inputs invalid: %v", err)})
	}
	// TODO: get user id from authenticated user
	// TODO: delete, just for testing
	userID, _ := uuid.Parse("7f1b2ed5-1201-4443-b997-56877fe31991")
	// create group with the only member being the owner
	group := rGroup.ToGroup(userID, []uuid.UUID{userID})
	// Create a group invitation for each invited user
	for _, member := range rGroup.Invites {
		groupInvitation := types.CreateGroupInvitation(&group)
		// store group invitation for user
		err := h.userStorage.AddGroupInvitationToUser(groupInvitation, member)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not create invitation: %v", err), "data": err})
		}
	}

	// store group in groupStorage
	err = h.groupStorage.AddGroup(group)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not create group: %v", err), "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Group created", "data": group})
}

// TODO: Check if id belongs to pending invitation
func (h Handler) HandleInvitation(c *fiber.Ctx) error {
	// get authenticated user
	user, err := h.getAuthenticatedUserFromHeader(c.GetReqHeaders())
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Authentication declined: %v", err)})
	}
	// parse invitation reply
	var rInvitation wire.InvitationReply
	if err := c.BodyParser(&rInvitation); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not parse invitation: %v", err), "data": err})
	}
	// handle invitation
	err = h.userStorage.HandleInvitation(rInvitation.Type, user.ID, rInvitation.ID, rInvitation.Accept)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Could not handle invitation: %v", err), "data": err})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Invitation handled"})
}

// TODO: maybe delete, or add authentication and allow only query of own groups
func (h Handler) GetGroup(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Parameter id is required", "data": nil})
	}
	gid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Unable to parse uuid: %s, error: %v", id, err), "data": err})
	}
	group, err := h.groupStorage.GetGroupByID(gid)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("Group not found: %v", err), "data": err})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Group found", "data": group})
}

// getAuthenticatedUserFromHeader tries to return the user associated with the given authentication token in the request header.
// If the token is invalid, an error will be returned.
// TODO: Generalize error messages
func (h Handler) getAuthenticatedUserFromHeader(reqHeader map[string]string) (types.User, error) {
	// get authentication cookie from header
	token := reqHeader["Cookie"]
	// check if cookie is present
	if token == "" {
		return types.User{}, errors.New("authentication cookie is missing")
	}
	// try to parse token
	tokenUUID, err := uuid.Parse(token)
	if err != nil {
		return types.User{}, errors.New("authentication cookie is invalid")
	}
	// get auth cookie from storage
	cookie, err := h.cookieStorage.GetCookieFromToken(tokenUUID)
	if err != nil {
		return types.User{}, err
	}
	// check if cookie is valid
	err = authentication.IsSessionCookieValid(cookie)
	if err != nil {
		return types.User{}, err
	}
	// get user from cookie
	user, err := h.userStorage.GetUserByID(cookie.UserID)
	if err != nil {
		return types.User{}, err
	}
	return user, nil
}
