package codes

//错误码定义
const (
	SUCCESS       = 200
	ERROR         = 500
	InvalidParams = 400

	ErrExistTag        = 10001
	ErrNotExistTag     = 10002
	ErrNotExistArticle = 10003

	ErrAuthCheckTokenFail    = 20001
	ErrAuthCheckTokenTimeout = 20002
	ErrAuthToken             = 20003
	ErrAuth                  = 20004

	ErrExistUser = 30001

	PageNotFound = 40001
)
