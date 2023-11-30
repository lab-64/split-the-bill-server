package handler

const (
	// Generic error messages
	ErrMsgParameterRequired = "Parameter %s is required"
	ErrMsgParseUUID         = "Could not parse uuid: %s, error: %v"
	ErrMsgInputsInvalid     = "Inputs invalid: %v"

	// Group - ERROR
	ErrMsgGroupParse    = "Could not parse group: %v"
	ErrMsgGroupCreate   = "Could not create group: %v"
	ErrMsgGroupNotFound = "Group not found: %v"

	// Group - SUCCESS
	SuccessMsgGroupFound  = "Group found"
	SuccessMsgGroupCreate = "Group created"

	// Bill - ERROR
	ErrMsgBillParse    = "Could not parse bill: %v"
	ErrMsgBillCreate   = "Could not create bill: %v"
	ErrMsgBillNotFound = "Bill not found: %v"

	// Bill - SUCCESS
	SuccessMsgBillFound  = "Bill found"
	SuccessMsgBillCreate = "Bill created"

	// Item - ERROR
	ErrMsgItemParse            = "Could not parse item: %v"
	ErrMsgItemCreate           = "Could not create item: %v"
	ErrMsgUpdateContributor    = "Could not update contributors: %v"
	ErrMsgItemContributorParse = "Could not parse contributor request: %v"
	ErrMsgItemNotFound         = "Item not found: %v"

	// Item - SUCCESS
	SuccessMsgItemCreate        = "Item created"
	SuccessMsgContributorUpdate = "Contributor updated"
	SuccesMsgItemFound          = "Item found"

	// User - ERROR
	ErrMsgUserParse            = "Could not parse user: %v"
	ErrMsgUserCreate           = "Could not create user: %v"
	ErrMsgUserDelete           = "Could not delete user: %v"
	ErrMsgUserNotFound         = "User not found: %v"
	ErrMsgUsersNotFound        = "Users not found: %v"
	ErrMsgUserLogin            = "Could not log in: %v"
	ErrMsgUserCredentialsParse = "Could not parse credentials: %v"
	ErrMsgInvitationHandle     = "Could not handle invitation: %v"
	ErrMsgBadPassword          = "Bad Password: %v"

	// User - SUCCESS
	SuccessMsgUserLogin  = "User logged in"
	SuccessMsgUserFound  = "User found"
	SuccessMsgUsersFound = "Users found"
	SuccessMsgUserCreate = "User created"
	SuccessMsgUserDelete = "User deleted"

	// Invitation - ERROR
	ErrMsgInvitationParse  = "Could not parse invitation: %v"
	ErrMsgInvitationCreate = "Could not create invitation: %v"

	// Invitation - SUCCESS
	SuccessMsgInvitationCreate  = "Invitation created"
	SuccessMsgInvitationFound   = "Invitation found"
	SuccessMsgInvitationHandled = "Invitation handled"
)
