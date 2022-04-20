package exceptions

import "fmt"

func Panic(context string) func(err error) {
	return func(err error) {
		if err != nil {
			errMsg := fmt.Sprintf("ERROR: %-10s => %s", context, err)
			panic(errMsg)
		}
	}
}
