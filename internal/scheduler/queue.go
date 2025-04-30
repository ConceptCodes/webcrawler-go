package scheduler

type SimpleQueue struct {
	Items []string
}

func New() *SimpleQueue {
	return &SimpleQueue{
		Items: make([]string, 0),
	}
}

func (q *SimpleQueue) Enqueue(item string) {
	q.Items = append(q.Items, item)
}

func (q *SimpleQueue) Dequeue() string {
	if len(q.Items) == 0 {
		return ""
	}
	item := q.Items[0]
	q.Items = q.Items[1:]
	return item
}

func (q *SimpleQueue) IsEmpty() bool {
	return len(q.Items) == 0
}

func (q *SimpleQueue) Len() int {
	return len(q.Items)
}
