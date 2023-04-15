package task_queue

type queue interface {
	pop(data interface{})
	push(data interface{})
}
