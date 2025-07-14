package queue

type Queue[T any] struct {
	count       int
	lowestCount int
	items       map[int]T
}

func New[T any](items ...T) Queue[T] {
	i := make(map[int]T, len(items))
	count := 0
	if len(items) > 0 {
		for _, it := range items {
			i[count] = it
			count++
		}
	}

	return Queue[T]{
		count:       count,
		lowestCount: 0,
		items:       i,
	}
}

func (q *Queue[T]) Items() []T {
	if q.IsEmpty() {
		return []T{}
	}

	its := make([]T, 0, q.Len())
	for i := q.lowestCount; i < q.count; i++ {
		its = append(its, q.items[i])
	}

	return its
}

func (q *Queue[T]) Enqueue(item T) {
	q.items[q.count] = item
	q.count++
}

func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T
	if q.IsEmpty() {
		return zero, false
	}

	it := q.items[q.lowestCount]

	delete(q.items, q.lowestCount)
	q.lowestCount++
	if q.IsEmpty() {
		q.count = 0
		q.lowestCount = 0
		q.items = make(map[int]T)
	}
	return it, true
}

func (q *Queue[T]) Peek() (T, bool) {
	var zero T
	if q.IsEmpty() {
		return zero, false
	}
	return q.items[q.lowestCount], true
}

func (q *Queue[T]) Len() int {
	return q.count - q.lowestCount
}

func (q *Queue[T]) IsEmpty() bool {
	return q.count == q.lowestCount
}
