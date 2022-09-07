package samples

func Aloc[O any, T []O](obj O, times int) (result T) {

	for i := 0; i <= times; i++ {
		result = append(result, obj)
	}

	for i, o := range result {
		result[len(result)-(i+1)]=o
	}

	return
}
