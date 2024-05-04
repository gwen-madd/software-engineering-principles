package list

//see - https://github.com/emirpasic/gods

type List[T comparable] interface {
	Get(index int) (T, bool)
	Remove(index int)
	Add(value T)
	Contains(value T) bool
	Set(index int, value T)
	Clear()
	Size()
}

type ArrayList[T comparable] struct {
	elements []T
	size     int
}

func (list *ArrayList[T]) Add(value T) {
	list.growSlice()
	list.elements[list.size] = value
	list.size++
}

func (list *ArrayList[T]) Contains(value T) bool {
	for i := 0; i < list.size; i++ {
		if list.elements[i] == value {
			return true
		}
	}
	return false
}

func (list *ArrayList[T]) Get(index int) (T, bool) {
	if list.size > index && index >= 0 {
		return list.elements[index], true
	}
	var t T
	return t, false
}

func (list *ArrayList[T]) Size() int {
	return list.size
}

func (list *ArrayList[T]) Set(index int, value T) {
	if list.size == index {
		list.Add(value)
		return
	}
	if list.size > index && index >= 0 {
		list.elements[index] = value
	}
}

func (list *ArrayList[T]) Remove(index int) {
	if list.size > index && index >= 0 {
		clear(list.elements[index : index+1])
		copy(list.elements[index:], list.elements[index+1:list.size])
		list.size--
	}
}

func (list *ArrayList[T]) Clear() {
	list.elements = make([]T, 0, 0)
	list.size = 0
}

func (list *ArrayList[T]) growSlice() {
	currSize := cap(list.elements)
	if list.size+1 >= currSize {
		newSize := currSize * 2
		newElements := make([]T, newSize, newSize)
		copy(newElements, list.elements)
		list.elements = newElements
	}
}
