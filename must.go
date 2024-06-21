package catchy

func Must[T interface{}](o T, err error) T {
	if err != nil {
		panic(err)
	}
	return o
}

func MustNoReturn(err error) {
	Must(never, err)
}
