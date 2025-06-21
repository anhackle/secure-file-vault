package response

const (
	ErrCodeSuccess              = 20000
	ErrCodeExternal             = 40000
	ErrCodeNotAuthorize         = 40001
	ErrCodeLoginFail            = 40002
	ErrCodeUserHasExists        = 40003
	ErrCodeExtensionNotAllowed  = 40004
	ErrCodeContentTypeNotAllowd = 40005
	ErrCodeInternal             = 50000
)

// message
var msg = map[int]string{
	ErrCodeSuccess:              "Success",
	ErrCodeLoginFail:            "Username or Password invalid",
	ErrCodeInternal:             "Internal server error",
	ErrCodeExternal:             "Bad request",
	ErrCodeUserHasExists:        "User existed",
	ErrCodeNotAuthorize:         "Not authorized",
	ErrCodeExtensionNotAllowed:  "File extension not allowed",
	ErrCodeContentTypeNotAllowd: "Content type not allowed",
}
