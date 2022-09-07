package samples

func Aloc[O any, T []O](obj O, times int) (result T) {

	result = make(T, times+1)
	for i := 0; i <= times; i++ {
		result[i] = obj
	}

	for _, o := range result {
		if isError(o) {
			result = nil
			return
		}
	}

	return
}

func isError(o any) bool {
	_, ok := o.(error)
	return ok
}
