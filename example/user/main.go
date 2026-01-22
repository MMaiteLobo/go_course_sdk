package main

import (

	"errors"
	"fmt"
	"os"
	
	userSDK "github.com/MMaiteLobo/go_course_sdk/user"

)

func main() {

	userTrans := userSDK.NewHTTPClient("http://localhost:8081", "")

	user, err := userTrans.Get("419bef06-9800-4b23-98b7-8438a6596a22")

	if err != nil {
		if errors.As(err, &userSDK.ErrNotFound{}) {
			fmt.Println("Not found", err.Error())
			os.Exit(1)
		} 
		fmt.Println("Internal server error", err.Error())
		os.Exit(1)
	}
	fmt.Println(user)

}
