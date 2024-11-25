package replayer

type Source[T any] interface {
	// returns the next item, and whether more items exist
	next() T
	hasMore() bool
	reset()
}

type ListSource[T any] struct {
	List  []T
	Index int
}

func (l *ListSource[T]) reset() {
	l.Index = 0
}

func (l *ListSource[T]) hasMore() bool {
	return l.Index < len(l.List)
}

func (l *ListSource[T]) next() T {
	defer func() {
		l.Index++
	}()

	return l.List[l.Index]
}
