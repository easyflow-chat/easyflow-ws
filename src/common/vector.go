package common

import (
	"sync"
)

type Vector[T comparable] struct {
	mtx  sync.Mutex
	data *[]T
}

func NewVector[T comparable]() *Vector[T] {
	data := make([]T, 0)
	return &Vector[T]{
		data: &data,
	}
}

func (v *Vector[T]) Write(elements []T) {
	v.mtx.Lock()
	defer v.mtx.Unlock()
	*v.data = elements
}

func (v *Vector[T]) Push(element T) {
	v.mtx.Lock()
	defer v.mtx.Unlock()
	*v.data = append(*v.data, element)
}

func (v *Vector[T]) Read() []T {
	v.mtx.Lock()
	defer v.mtx.Unlock()
	dataCopy := make([]T, len(*v.data))
	copy(dataCopy, *v.data)
	return dataCopy
}

func (v *Vector[T]) Len() int {
	v.mtx.Lock()
	defer v.mtx.Unlock()
	return len(*v.data)
}

func (v *Vector[T]) Clear() {
	v.mtx.Lock()
	defer v.mtx.Unlock()
	*v.data = make([]T, 0)
}

func (v *Vector[T]) Get(index int) (T, bool) {
	v.mtx.Lock()
	defer v.mtx.Unlock()
	if index >= 0 && index < len(*v.data) {
		return (*v.data)[index], true
	}
	var zero T
	return zero, false
}

func (v *Vector[T]) Update(index int, element T) bool {
	v.mtx.Lock()
	defer v.mtx.Unlock()
	if index >= 0 && index < len(*v.data) {
		(*v.data)[index] = element
		return true
	}
	return false
}

func (v *Vector[T]) RemoveByIndex(index int) bool {
	v.mtx.Lock()
	defer v.mtx.Unlock()
	if index >= 0 && index < len(*v.data) {
		*v.data = append((*v.data)[:index], (*v.data)[index+1:]...)
		return true
	}
	return false
}

func (v *Vector[T]) Pop() (T, bool) {
	v.mtx.Lock()
	defer v.mtx.Unlock()
	if len(*v.data) == 0 {
		var zero T
		return zero, false
	}
	element := (*v.data)[len(*v.data)-1]
	*v.data = (*v.data)[:len(*v.data)-1]
	return element, true
}
