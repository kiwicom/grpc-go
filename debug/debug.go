package debug

import (
	"sync/atomic"
)

type Handler func(rcvd, limit, n, pendingData, pendingUpdate, delta uint32)

type handler struct {
	fn Handler
}

var globalHandler atomic.Pointer[handler]

// SetHandler sets the handler for flow control debugging.
func SetHandler(fn Handler) {
	globalHandler.Store(&handler{fn: fn})
}

// HandleError calls the handler if it was set.
func HandleError(rcvd, limit, n, pendingData, pendingUpdate, delta uint32) {
	if h := globalHandler.Load(); h != nil && h.fn != nil {
		h.fn(rcvd, limit, n, pendingData, pendingUpdate, delta)
	}
}

type LogHandler func(msg string)

type logHandler struct {
	fn LogHandler
}

var globalLogHandler atomic.Pointer[logHandler]

func SetLogHandler(fn LogHandler) {
	globalLogHandler.Store(&logHandler{fn: fn})
}

func Log(msg string) {
	if h := globalLogHandler.Load(); h != nil && h.fn != nil {
		h.fn(msg)
	}
}
