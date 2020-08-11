package servlet

type TaskProc func(ctx TaskContext) error

type EventProc func(ctx TaskContext, event interface{}) error

type MessageProc func(ctx TaskContext, topic string, msg interface{}, channel ...string) error
