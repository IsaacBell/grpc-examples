package main

import (
	"fmt"
	"time"
)

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().Format("15:04:05") + " " + string(bytes))
}
