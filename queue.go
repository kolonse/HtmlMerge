package main

type Queue []interface{}

func (q *Queue) In(e interface{}) {
	*q = append(*q, e)
}

func (q *Queue) Out() interface{} {
	if q.Empty() {
		return nil
	}
	e := (*q)[0]
	*q = (*q)[1:]
	return e
}

func (q Queue) Empty() bool {
	if q.Size() == 0 {
		return true
	}
	return false
}

func (q *Queue) Clear() {
	if q.Empty() {
		return
	}
	*q = NewQueue()
}
func (q Queue) Size() int {
	return len(q)
}

func NewQueue() Queue {
	return make([]interface{}, 0)
}
