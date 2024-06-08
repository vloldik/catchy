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

type Never = interface{}

var never = Never(nil)

type Catchy[T interface{}] struct {
	GetValue func() (T, error)
	OnSucess func(T)
	OnError  func(error)
}

func WithGetValueFunc[T interface{}](getValue func() (T, error)) Catchy[T] {
	return Catchy[T]{
		GetValue: getValue,
	}
}

func WithNoReturnFunc(noReturnFunc func() error) Catchy[Never] {
	return Catchy[Never]{
		GetValue: func() (Never, error) {
			return never, noReturnFunc()
		},
	}
}

func (c Catchy[T]) WithOnSuccess(useValue func(T)) Catchy[T] {
	c.OnSucess = useValue
	return c
}

func (c Catchy[T]) WithOnError(onError func(error)) Catchy[T] {
	c.OnError = onError
	return c
}

func (c Catchy[T]) Do() {
	value, err := c.GetValue()
	if err != nil && c.OnError != nil {
		c.OnError(err)
		return
	} else if err != nil {
		panic(err)
	}
	if c.OnSucess != nil {
		c.OnSucess(value)
	}
}
