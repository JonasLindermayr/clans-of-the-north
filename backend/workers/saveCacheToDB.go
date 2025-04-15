package workers

import (
	"fmt"
	"time"

	"github.com/JonasLindermayr/clans-of-the-north/backend/registry"
)

func SaveCacheToDB() {

	for {
		time.Sleep(time.Second * 20)

		fmt.Println("Saving cache to DB")

		registry.SaveCacheToDB()

	}
}