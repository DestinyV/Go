package errcode

import "net/http"

type ServiceError struct {
	Status int         `json:"-"`
	Code   ErrorCode   `json:"code"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data,omitempty"`
}

func (err *ServiceError) Error() string {
	return err.Msg
}

func (err *ServiceError) WithData(data interface{}) *ServiceError {
	newErr := &ServiceError{
		Status: err.Status,
		Code:   err.Code,
		Msg:    err.Msg,
		Data:   data,
	}
	return newErr
}

func (err *ServiceError) WithError(e error) *ServiceError {
	newErr := &ServiceError{
		Status: err.Status,
		Code:   err.Code,
		Msg:    err.Msg,
		Data:   err.Msg + "(" + e.Error() + ")",
	}
	return newErr
}

var (
	ErrorOK             = &ServiceError{http.StatusOK, OK, "成功", nil}
	ErrorNotFound       = &ServiceError{http.StatusNotFound, NotFound, "请求路径有误", nil}
	ErrorNoMethod       = &ServiceError{http.StatusMethodNotAllowed, NoMethod, "请求方式错误", nil}
	ErrorServerErr      = &ServiceError{http.StatusInternalServerError, ServerError, "服务器异常", nil}
	ErrorInvalidParam   = &ServiceError{http.StatusBadRequest, InvalidParam, "参数错误", nil}
	ErrorUnauthorized   = &ServiceError{http.StatusUnauthorized, Unauthorized, "未授权", nil}
	ErrorAlreadyCheckin = &ServiceError{http.StatusBadRequest, AlreadyCheckin, "已经签到过了", nil}
	ErrorAuthority      = &ServiceError{http.StatusUnauthorized, AuthorityError, "权限不足", nil}
	ErrorEmptyReward    = &ServiceError{http.StatusBadRequest, EmptyReward, "奖池为空", nil}
)
