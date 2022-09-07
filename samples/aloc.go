package samples

import "fmt"

func Aloc[O any, T []O](obj O, times int) (result T) {

	for i := 0; i <= times; i++ {
		result = append(result, obj)
	}

	values := make([]any, 0)
	for _, o := range result {
		values = append(values, fmt.Sprintf("%+v\n", o))
	}
	fmt.Println(values...)

	return
}
