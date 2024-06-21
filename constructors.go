package catchy

func WithGetValueFunc[T interface{}](getValue func() (T, error)) *Catchy[T] {
	return &Catchy[T]{
		GetValue: getValue,
	}
}

func WithNoReturnFunc(noReturnFunc func() error) *Catchy[Never] {
	return WithGetValueFunc(func() (Never, error) {
		return never, noReturnFunc()
	})
}
