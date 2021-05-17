package SynchronousQueue

import (
	"gopkg.in/eapache/queue.v1"
	"sync"
)

const minQueueLen = 32

type SynchronousQueue struct {
	lock    sync.Mutex
	popable *sync.Cond
	buffer  *queue.Queue
	closed  bool
}

func NewSyncQueue() *SynchronousQueue {
	ch := &SynchronousQueue{
		buffer: queue.New(),
	}
	ch.popable = sync.NewCond(&ch.lock)
	return ch
}
func (q *SynchronousQueue) Pop() (v interface{}) {
	c := q.popable
	buffer := q.buffer

	q.lock.Lock()
	for buffer.Length() == 0 && !q.closed {
		c.Wait()
	}

	if buffer.Length() > 0 {
		v = buffer.Peek()
		buffer.Remove()
	}

	q.lock.Unlock()
	return
}
func (q *SynchronousQueue) Push(v interface{}) {
	q.lock.Lock()
	if !q.closed {
		q.buffer.Add(v)
		q.popable.Signal()
	}
	q.lock.Unlock()
}

type Collection interface {
	Size() int
	IsEmpty() bool
	Contains(o interface{}) bool
	ToArray() []interface{}
	ToArrayWithContain(T interface{}) []interface{}
	Add(e interface{}) bool
	Remove(o interface{}) bool
	ContainAll()
}

//func New() *SynchronousQueue {
//
//}

//void	clear()
func (q *SynchronousQueue) Clear() {
}

//boolean	contains(Object o)
func (q *SynchronousQueue) Contains(o interface{}) bool {
	return false
}

//boolean	containsAll(Collection<?> c)
func (q *SynchronousQueue) ContainsAll(c Collection) bool {
	return c.IsEmpty()
}

//int	drainTo(Collection<? super E> c)

//int	drainTo(Collection<? super E> c, int maxElements)

//boolean	isEmpty()
func (q *SynchronousQueue) IsEmpty() bool {
	return true
}

//Iterator<E>	iterator()

//boolean	offer(E e)

//boolean	offer(E e, long timeout, TimeUnit unit)

//E	peek()
func (q *SynchronousQueue) Peek() interface{} {
	return nil
}

//E	poll()

//E	poll(long timeout, TimeUnit unit)

//void	put(E e)

//int	remainingCapacity()
func (q *SynchronousQueue) RemainingCapacity() int {
	return 0
}

//boolean	remove(Object o)
func (q *SynchronousQueue) Remove(o interface{}) bool {
	return false
}

//boolean	removeAll(Collection<?> c)
func (q *SynchronousQueue) RemoveAll(c Collection) bool {
	return false
}

//boolean	retainAll(Collection<?> c)
func (q *SynchronousQueue) RetainAll(c Collection) bool {
	return false
}

//int	size()
func (q *SynchronousQueue) Size() int {
	return 0
}

//Spliterator<E>	spliterator()

//E	take()
//Object[]	toArray()
func (q *SynchronousQueue) ToArray() []interface{} {
	ret := make([]interface{}, 0)
	return ret
}

//<T> T[]	toArray(T[] a)
func (q *SynchronousQueue) ToArrayWithArray(a []interface{}) []interface{} {
	if len(a) > 0 {
		a[0] = nil
	}
	return a
}
