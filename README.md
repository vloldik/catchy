# catchy

The catchy package provides a convenient way to handle errors and success cases in Go programming language.

## Usage

The package includes the following functions and types:

### Must[T interface{}](o T, err error) T

The Must function takes a value o and an error err. If err is not nil, it panics with the given error. Otherwise, it returns the value o.

### MustNoReturn(err error)

The MustNoReturn function is a helper function that calls Must with a Never value and the given error. It is useful when you have a function that doesn't return any value.

### Never

The Never type is an empty interface used as a placeholder for functions that don't return any value.

### Catchy[T interface{}]

The Catchy struct represents a catchy operation. It has three fields:

- GetValue: A function that returns a value of type T and an error.
- OnSuccess: A function that is called when the operation succeeds.
- OnError: A function that is called when the operation encounters an error.

### WithGetValueFunc[T interface{}](getValue func() (T, error)) Catchy[T]

The WithGetValueFunc function creates a new Catchy instance with the given getValue function.

### WithNoReturnFunc(noReturnFunc func() error) Catchy[Never]

The WithNoReturnFunc function creates a new Catchy instance for functions that don't return any value.

### WithOnSuccess(useValue func(T)) Catchy[T]

The WithOnSuccess method sets the OnSuccess function for a Catchy instance.

### WithOnError(onError func(error)) Catchy[T]

The WithOnError method sets the OnError function for a Catchy instance.

### Do()

The Do method executes the catchy operation. It calls the GetValue function and handles the returned value and error based on the OnSuccess and OnError functions.

## Example

Here's an example usage of the catchy package:

```go
func main() {
    catchy.WithGetValueFunc(func() (int, error) {
        return 42, nil
    }).WithOnSuccess(func(value int) {
        fmt.Println("Success:", value)
    }).WithOnError(func(err error) {
        fmt.Println("Error:", err)
    }).Do()
}
```

In this example, we create a Catchy instance with a getValue function that returns the value 42 and no error. We set the OnSuccess function to print the success message and the OnError function to print the error message. Finally, we call the Do method to execute the catchy operation.

Output:
`Success: 42`


If the getValue function returns an error, the OnError function will be called instead:

```go
func main() {
    catchy.WithGetValueFunc(func() (int, error) {
        return 0, errors.New("something went wrong")
    }).WithOnSuccess(func(value int) {
        fmt.Println("Success:", value)
    }).WithOnError(func(err error) {
        fmt.Println("Error:", err)
    }).Do()
}
```

Output:
`Error: something went wrong`


That's it! The catchy package helps you handle errors and success cases in a clean and concise way.
