package catchy

type Never = interface{}

var never = Never(nil)

type IDoable interface {
	Do() error
}

type Catchy[T interface{}] struct {
	GetValue func() (T, error)
	OnSucess func(T)
	OnError  func(error)
	next     *DoableNode
	last     *DoableNode
}

func (c *Catchy[T]) WithOnSuccess(useValue func(T)) *Catchy[T] {
	c.OnSucess = useValue
	return c
}

func (c *Catchy[T]) WithOnError(onError func(error)) *Catchy[T] {
	c.OnError = onError
	return c
}

func (c *Catchy[T]) WithGetValueFunc(getValue func() (T, error)) *Catchy[T] {
	c.GetValue = getValue
	return c
}

func (c *Catchy[T]) doSelf() error {
	value, err := c.GetValue()
	if err != nil {
		return err
	}
	if c.OnSucess != nil {
		c.OnSucess(value)
	}
	return nil
}

func (c *Catchy[T]) Do() error {
	err := c.doSelf()
	if err != nil && c.OnError != nil {
		c.OnError(err)
	}
	if err != nil {
		return err
	}
	if c.next != nil {
		err = c.next.Do()
	}
	if err != nil && c.OnError != nil {
		c.OnError(err)
	}
	return err
}

func (c *Catchy[T]) DoNext(next IDoable) *Catchy[T] {
	node := newDoableNode(next)
	if c.next == nil {
		c.next = node
		c.last = node
	} else {
		c.last.Next = node
		c.last = node
	}
	return c
}
