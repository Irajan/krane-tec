package errorUtils

import (
	"fmt"
)

func HandleError(err error, message ...string) {

	if err == nil {
		return
	}

	var errorMessage string

	if len(message) == 0 {
		errorMessage = "Something went wrong"
	} else {
		errorMessage = message[0]
	}

	// logger.Error(errorMessage)

	fmt.Println(err)

	fmt.Println(errorMessage)

	panic(err)

}
