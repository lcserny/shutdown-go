package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/shutdown", func(writer http.ResponseWriter, request *http.Request) {
		values := request.URL.Query()
		secondsInt := "0"
		if seconds, exists := values["seconds"]; exists {
			secondsInt = seconds[0]
		}
		executeShutdownFromWSL(secondsInt)
	})
	fmt.Println("Server started on port 8082")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}

func executeShutdownFromWSL(seconds string) {
	var cmdErr bytes.Buffer
	cmd := exec.Command("shutdown", "-s", "-t", seconds)
	cmd.Stderr = &cmdErr
	if err := cmd.Run(); err != nil {
		cmdErr.WriteString(err.Error())
	}
}
