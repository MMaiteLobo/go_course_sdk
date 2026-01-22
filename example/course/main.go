package main

import (
	"errors"
	"fmt"
	"os"

	courseSDK "github.com/MMaiteLobo/go_course_sdk/course"
)

func main() {
	courseTrans := courseSDK.NewHTTPClient("http://localhost:8082", "")

	course, err := courseTrans.Get("5ef81830-6806-4dc8-84bd-652b4570c023")
	if err != nil {
		if errors.As(err, &courseSDK.ErrNotFound{}) {
			fmt.Println("Not found", err.Error())
			os.Exit(1)
		}
		fmt.Println("Internal server error", err.Error())
		os.Exit(1)
	}
	fmt.Println(course)
}