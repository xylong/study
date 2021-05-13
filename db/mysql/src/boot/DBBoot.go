package boot

import "study/db/mysql/src/config"

func DBBoot() {
	go func() {
		if err := config.InitDB(); err != nil {
			config.ErrorChan <- err
		}
	}()
}
