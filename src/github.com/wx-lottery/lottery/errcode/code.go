package errcode

type ErrorCode int

const (
	OK          = 0
	NotFound    = 404
	NoMethod    = 405
	ServerError = 500

	InvalidParam   = 1001
	Unauthorized   = 1002
	AlreadyCheckin = 1003
	AuthorityError = 1004
	EmptyReward    = 1005
)
