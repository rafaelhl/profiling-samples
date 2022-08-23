package samples

import "fmt"

func Aloc[O any, T []O](obj O, times int) (result T) {

	for i := 0; i <= times; i++ {
		result = append(result, obj)
	}

	for _, o := range result {
		fmt.Printf("%+v\n", o)
	}

	return
}
