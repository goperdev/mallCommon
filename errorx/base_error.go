package errorx

import "fmt"

type CodeErrorResponse struct {
	TraceID   string                   `json:"TraceID"`
	ErrorData CodeErrorResponseContent `json:"Error"`
	Data      interface{}              `json:"Data"`
}

type CodeErrorResponseContent struct {
	Code ErrCode `json:"Code"`
	Msg  string  `json:"Msg"`
}

func (o ErrCode) GenResonse(traceID string, msg string) CodeErrorResponse {
	targetMsg := msg
	if len(msg) < 1 {
		targetMsg = codeMsgMap[o]
	}
	return CodeErrorResponse{
		TraceID: traceID,
		ErrorData: CodeErrorResponseContent{
			Code: o,
			Msg:  targetMsg,
		},
	}
}

func (o ErrCode) GenResonseWithData(traceID string, data interface{}) CodeErrorResponse {
	return CodeErrorResponse{
		TraceID: traceID,
		ErrorData: CodeErrorResponseContent{
			Code: o,
			Msg:  codeMsgMap[o],
		},
		Data: data,
	}
}

func (o ErrCode) GetMsg() string {
	return codeMsgMap[o]
}

// GenError
func (o ErrCode) GenError(msg ...string) *CodeErrorResponseContent {
	resp := CodeErrorResponseContent{
		Code: o,
		Msg:  codeMsgMap[o],
	}
	if len(msg) > 0 {
		resp.Msg = msg[0]
	}

	return &resp
}

// Error 错误打印
func (o *CodeErrorResponseContent) Error() string {
	return fmt.Sprintf("code:%s, msg:%s", string(o.Code), o.Msg)
}

func (o *CodeErrorResponseContent) GenResonse(traceID string) CodeErrorResponse {
	return CodeErrorResponse{
		TraceID:   traceID,
		ErrorData: *o,
	}
}

type ErrCode string

const (
	ErrCodeNone                    ErrCode = "Success"
	ErrCodeDataNotFound            ErrCode = "DataNotFound"            // 数据不存在
	ErrCodeInternal                ErrCode = "InternalError"           // 内部服务错误
	ErrCodeInvalidParamter         ErrCode = "InvalidParamter"         // 参数错误
	ErrCodeAuthorizationTokenValid ErrCode = "AuthorizationTokenValid" // token失效
	ErrCodeUserForzen              ErrCode = "UserForzen"              // 用户冻结
	ErrCodeVirifyCodeInvalid       ErrCode = "InvalidVerifyCode"       // 验证码已失效
)

var (
	codeMsgMap = map[ErrCode]string{
		ErrCodeDataNotFound:            "数据不存在",
		ErrCodeInternal:                "内部服务错误",
		ErrCodeInvalidParamter:         "参数错误",
		ErrCodeAuthorizationTokenValid: "token失效",
		ErrCodeUserForzen:              "用户已冻结",
	}
)
