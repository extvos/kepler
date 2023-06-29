package servlet

// TaskHandler defines type of task function.
type TaskHandler = func(ctx TaskContext) error

// EventHandler defines type of event handler function.
type EventHandler = func(ctx TaskContext, event interface{}) error

// MessageHandler defines type of message handler function.
type MessageHandler = func(ctx TaskContext, topic string, msg interface{}, channel ...string) error
