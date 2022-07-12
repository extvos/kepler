package servlet

// TaskProc defines type of task function.
type TaskProc func(ctx TaskContext) error

// EventProc defines type of event handler function.
type EventProc func(ctx TaskContext, event interface{}) error

// MessageProc defines type of message handler function.
type MessageProc func(ctx TaskContext, topic string, msg interface{}, channel ...string) error
