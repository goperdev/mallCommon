package mallCommon

import "github.com/zeromicro/go-zero/core/logx"

// CtxMessage context Message
type CtxMessage struct {
	TraceID string `json:"traceID"`
}

// genLogFields 生成logFields
func (o *CtxMessage) genLogFields() []logx.LogField {
	result := make([]logx.LogField, 0)
	if len(o.TraceID) > 0 {
		result = append(result, logx.LogField{
			Key:   "traceID",
			Value: o.TraceID,
		})
	}
	return result
}
