package samples

func Aloc[O any, T []O](obj O, times int) (result T) {

	for i := 0; i <= times; i++ {
		result = append(result, obj)
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
