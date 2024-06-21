Catchy is a generic Go language structure that provides convenient ways to process values ​​using processing chains. It includes handling of success and error functions, as well as the ability to build a chain of actions.
## TODO: 

Add the ability to pass an argument along a chain

#### Note - If an error occurs in a function chain, the chain is interrupted and if the `OnError` function is specified, `OnError` is called.

## Constructor functions
`func WithGetValueFunc[T interface{}](getValue func() (T, error)) Catchy[T] `

It is used to create `Catchy` at once with a function to get a value and an error

`func WithNoReturnFunc(noReturnFunc func() error) Catchy[Never]`

Used to create `Catchy` at once with the function of getting only the error

## Example usage

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/vloldik/catchy"
)

func main() {
	var result string

	catchy.WithGetValueFunc(func() (string, error) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		return reader.ReadString('\n')
	}).WithOnSuccess(func(input string) {
		result = input
	}).DoNext(catchy.WithGetValueFunc(func() (int, error) {
		return strconv.Atoi(result[:len(result)-2])
	}).WithOnSuccess(func(i int) {
		fmt.Printf("You entered valid int %d", i)
	})).WithOnError(func(err error) {
		fmt.Printf("Error %s", err)
	}).Do()
}
```

## Basic types and structures

### Never

` type Never = interface{} `

` var never = Never(nil) `


Never is an interface that represents “nothing” and is used where there is no return value.

### IDoable

```go
type IDoable interface {
    Do() error
}
```


IDoable is an interface that declares a Do() error method. This interface is used in processing chains.

### Catchy
```go
type Catchy[T interface{}] struct {
    GetValue func() (T, error)
    OnSucess func(T)
    OnError func(error)
    next *DoableNode
    last *DoableNode
}
```


Catchy is a generic structure with type T that includes functions for getting a value, handling success and errors, and nodes for constructing a chain.

## Catchy Methods

### WithOnSuccess

`func (c Catchy[T]) WithOnSuccess(useValue func(T)) Catchy[T]`


The method sets the useValue function to handle the successful result. Returns the modified Catchy structure.

### WithOnError

`func (c Catchy[T]) WithOnError(onError func(error)) Catchy[T]`


The method sets up the onError function to handle errors. Returns the modified Catchy structure.

### WithGetValueFunc

`func (c Catchy[T]) WithGetValueFunc(getValue func() (T, error)) Catchy[T]`


The method sets up the getValue function to get the value. Returns the modified Catchy structure.

### Do

`func (c Catchy[T]) Do() error`


The method executes the current Catchy structure and, if there is a next one in the chain, executes it. Returns an error if one occurs.

### DoNext

`func (c Catchy[T]) DoNext(next IDoable) Catchy[T]`


The method adds the next IDoable node to the processing chain. Returns the modified Catchy structure.
