package common

import (
	"fmt"
	"log"
	"os"
)

func LogOnChainCronError(lastBlock uint64, err error) bool {
	if err != nil {
		err = fmt.Errorf("[Block: %d] %s", lastBlock, err)
	}
	return LogError("logs/cron_errors.log", err)
}

func LogError(fileName string, err error) bool {
	if err == nil {
		return false
	}

	f, e := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if e != nil {
		log.Fatalf("error opening file %v", e)
	}

	defer f.Close()

	log.SetOutput(f)
	log.Println(err)
	return true
}
