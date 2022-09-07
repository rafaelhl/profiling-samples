package samples

import (
	"bytes"
	"github.com/goccy/go-json"
	"fmt"
	"os"
)

var Path = "samples/"

func Unmarshall[T any](times int) (result T) {
	file, err := os.ReadFile(fmt.Sprintf("%stest.json", Path))
	if err != nil {
		panic(err)
	}

	for i := 0; i <= times; i++ {
		err = json.NewDecoder(bytes.NewReader(file)).Decode(&result)
		if err != nil {
			panic(err)
		}
	}

	return
}
