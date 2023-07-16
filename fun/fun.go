// Putting the FUN in FUNctions!
package fun

// execute a func over all elements in list, returning a new list with the
// results
func Map[T any, Y any](list []T, apply func(i int, x T) Y) []Y {
	nilListCheck(list)
	result := make([]Y, 0, len(list))
	for i, v := range list {
		result = append(result, apply(i, v))
	}
	return result
}

// Apply a func to each item in a slice
func Each[T any](list []T, apply func(x T)) {
	nilListCheck(list)
	for _, v := range list {
		apply(v)
	}
}

func Values[K comparable, T any](m map[K]T) []T {
	result := make([]T, 0, len(m))

	for _, v := range m {
		result = append(result, v)
	}
	return result
}

// convert a slice to a map, the provided function will return the key and value to populate the
// map with
func ToMap[K comparable, V any](list []V, fn func(idx int, v V) (K, V)) map[K]V {
	destMap := make(map[K]V)

	for idx, val := range list {
		k, v := fn(idx, val)
		destMap[k] = v

	}
	return destMap
}

// returns a new list that is a subset of the previous list with elements that
// evaluate true for the given filter function.
func Filter[T any](list []T, filter func(v T) bool) []T {
	nilListCheck(list)
	foundList := make([]T, 0)
	for _, v := range list {
		if found := filter(v); found {
			foundList = append(foundList, v)
		}
	}
	return foundList
}

// Retrieve the first element from the list or the provided `other`
func FirstOrElse[T any](list []T, other T) T {
	nilListCheck(list)
	if len(list) == 0 {
		return other
	}
	return list[0]
}

func Zip[T any, V any](slice1 []T, slice2 []V) []Tuple[T, V] {
	var (
		minLength int
		result    = []Tuple[T, V]{}
		length1   = len(slice1)
		length2   = len(slice2)
	)

	if length1 <= length2 {
		minLength = length1
	} else {
		minLength = length2
	}

	for i := 0; i < minLength; i++ {
		newTuple := Tuple[T, V]{
			E1: slice1[i],
			E2: slice2[i],
		}
		result = append(result, newTuple)
	}
	return result
}

// Reduce takes a list(slice) and produces a single value. The [initial] parameter
// will be what is passed to to the first iteration of the [reducer] fun.
func Reduce[T any, V any](list []T, initial V, reducer func(v T, accumulator V) V) V {
	nilListCheck(list)
	x := initial
	for _, v := range list {
		x = reducer(v, x)
	}
	return x
}

// Same as Reduce, but also passes the loop iterator
func ReduceI[T any, V any](list []T, initial V, reducer func(i int, v T, accumulator V) V) V {
	nilListCheck(list)
	x := initial
	for i, v := range list {
		x = reducer(i, v, x)
	}
	return x
}

type Tuple[T any, V any] struct {
	E1 T
	E2 V
}

func nilListCheck[T any](maybeList []T) {
	if maybeList == nil {
		panic("must be a list; nil not allowed")
	}
}
