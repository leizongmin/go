package iterutil

type Array = []Item
type Item = interface{}

func Map(arr Array, mapper func(value Item) Item) Array {
	size := len(arr)
	newArr := make([]Item, size)
	i := 0
	for i < size {
		newArr[i] = mapper(arr[i])
		i++
	}
	return newArr
}

func Filter(arr Array, predicate func(value Item) bool) Array {
	size := len(arr)
	newArr := make(Array, 0)
	i := 0
	for i < size {
		if predicate(arr[i]) {
			newArr = append(newArr, arr[i])
		}
		i++
	}
	return newArr
}

func Reduce(arr Array, initialValue Item, f func(total, currentValue Item, currentIndex int) Item) Item {
	size := len(arr)
	i := 0
	total := initialValue
	for i < size {
		total = f(total, arr[i], i)
		i++
	}
	return total
}
