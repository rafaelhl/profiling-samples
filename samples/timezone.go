package samples

import (
	"time"
)

func Timezone(times int) (tz *time.Location) {
	for i := 0; i <= times; i++ {
		var err error
		tz, err = time.LoadLocation("America/Sao_Paulo")
		if err != nil {
			panic(err)
		}
	}

	return
}
