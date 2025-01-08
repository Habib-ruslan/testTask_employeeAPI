package Logger

import (
	"fmt"
	"log"
)

func Log(message string, args ...any) {
	log.Printf(message, args)
	fmt.Printf(message, args)
}

func Error(message string, args ...any) error {
	log.Printf(message, args)
	return fmt.Errorf(message, args)
}

func Fatal(message string, args ...any) {
	log.Fatalf(message, args)
}
