package handler

const (
	// Generic error messages
	ErrMsgParameterRequired = "Parameter %s is required"
	ErrMsgParseUUID         = "Could not parse uuid: %s, error: %v"
	ErrMsgInputsInvalid     = "Inputs invalid: %v"

	// Group - ERROR
	ErrMsgGroupParse    = "Could not parse group: %v"
	ErrMsgGroupCreate   = "Could not create group: %v"
	ErrMsgGroupUpdate   = "Could not update group: %v"
	ErrMsgGroupNotFound = "Group not found: %v"
	ErrMsgGetUserGroups = "Could not load user groups: %v"
	ErrMsgGroupDelete   = "Could not delete group: %v"

	// Group - SUCCESS
	SuccessMsgGroupFound  = "Group found"
	SuccessMsgGroupsFound = "Groups found"
	SuccessMsgGroupCreate = "Group created"
	SuccessMsgGroupUpdate = "Group updated"
	SuccessMsgGroupDelete = "Group deleted"

	// Group transaction - ERROR
	ErrMsgGroupTransactionCreate = "Could not create group transaction: %v"
	ErrMsgGetUserTransactions    = "Could not get group transaction: %v"

	// Group transaction - SUCCESS
	SuccessMsgGroupTransactionCreate = "Group transaction created"
	SuccessMsgGroupTransactionFound  = "Group transaction found"

	// Bill - ERROR
	ErrMsgBillParse    = "Could not parse bill: %v"
	ErrMsgBillCreate   = "Could not create bill: %v"
	ErrMsgBillUpdate   = "Could not update bill: %v"
	ErrMsgBillNotFound = "Bill not found: %v"
	ErrMsgBillGetAll   = "Could not load bills: %v"
	ErrMsgBillDelete   = "Could not delete bill: %v"

	// Bill - SUCCESS
	SuccessMsgBillFound  = "Bill found"
	SuccessMsgBillCreate = "Bill created"
	SuccessMsgBillUpdate = "Bill updated"
	SuccessMsgBillGetAll = "Bills found"
	SuccessMsgBillDelete = "Bill deleted"

	// User - ERROR
	ErrMsgUserParse            = "Could not parse user: %v"
	ErrMsgUserCreate           = "Could not create user: %v"
	ErrMsgUserDelete           = "Could not delete user: %v"
	ErrMsgUserUpdate           = "Could not update user: %v"
	ErrMsgUserNotFound         = "User not found: %v"
	ErrMsgUsersNotFound        = "Users not found: %v"
	ErrMsgUserLogin            = "Could not log in: %v"
	ErrMsgUserCredentialsParse = "Could not parse credentials: %v"
	ErrMsgBadPassword          = "Bad Password: %v"
	ErrMsgUserLogout           = "Could not log out: %v"

	// User - SUCCESS
	SuccessMsgUserLogin  = "User logged in"
	SuccessMsgUserFound  = "User found"
	SuccessMsgUsersFound = "Users found"
	SuccessMsgUserCreate = "User created"
	SuccessMsgUserDelete = "User deleted"
	SuccessMsgUserUpdate = "User updated"
	SuccessMsgUserLogout = "User logged out"

	// Invitation - ERROR
	ErrMsgInvitationHandle = "Could not handle invitation: %v"

	// Invitation - SUCCESS
	SuccessMsgInvitationHandled = "Invitation handled"

	ErrMsgUserImageUpload = "Could not upload user image: %v"
)
