package samples

import (
	"bytes"
	"github.com/goccy/go-json"
	"fmt"
	"os"
)

var Path = "samples/"

func Unmarshall[T any](times int) (result T) {
	filename := fmt.Sprintf("%stest.json", Path)
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var cache map[string]T = make(map[string]T)

	for i := 0; i <= times; i++ {
		var ok bool
		result, ok = cache[filename]
		if ok {
			return
		}

		err = json.NewDecoder(bytes.NewReader(file)).Decode(&result)
		if err != nil {
			panic(err)
		}
		cache[filename] = result
	}

	return
}
