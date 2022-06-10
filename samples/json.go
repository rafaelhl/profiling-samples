package samples

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var Path = "samples/"

func Unmarshall[T any](times int) (result T) {
	file, err := ioutil.ReadFile(fmt.Sprintf("%stest.json", Path))
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
