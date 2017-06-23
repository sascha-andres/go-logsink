package cmd

import (
	"fmt"
	"log"

	"github.com/nightlyone/lockfile"
	"github.com/spf13/viper"
)

func handleLock(method func()) {
	if "" != viper.GetString("lockfile") {
		lock, err := lockfile.New(viper.GetString("lockfile"))
		if err != nil {
			log.Fatal(err) // handle properly please!
		}
		err = lock.TryLock()

		// Error handling is essential, as we only try to get the lock.
		if err != nil {
			log.Fatal(fmt.Errorf("Cannot lock %q, reason: %v", lock, err))
		}

		defer lock.Unlock()
	}
	method()
}
