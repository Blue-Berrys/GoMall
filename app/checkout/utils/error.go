package utils

func MustHandleError(err error) {
	if err != nil {
		panic(err)
	}
}
