package samples

import (
	"time"
)

var cache = make(map[string]*time.Location)

func Timezone(times int) (tz *time.Location) {
	for i := 0; i <= times; i++ {
		var err error
		
		var ok bool
		location := "America/Sao_Paulo"
		tz, ok = cache[location]
		if ok {
			return tz
		}

		tz, err = time.LoadLocation(location)
		if err != nil {
			panic(err)
		}
		cache[location] = tz
	}

	return
}
