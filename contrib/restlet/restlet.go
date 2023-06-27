package restlet

type Handler[T any] struct {
}

func (Handler[T]) onPost(obj T) (T, error) {
	return obj, nil
}
