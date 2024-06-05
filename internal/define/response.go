package define

type ResponseCode int64

const (
	ResponseCodeOk       = 0
	ResponseCodeBusError = 400
	ResponseCodePanic    = 500
)
