## Basic syntax

```go
func main(){}
```

## Parameters

```go
func main(){
  for i := 0; i < 5; i++ {
    sayMessage("Hello Go!", i)
  }
}

func sayMessage(msg string, idx int) {
  fmt.Println(msg)
  fmt.Println("The value of the index is", idx)
}
```

```go
func main() {
  greeting := "Hello"
  name := "Stacey"
  sayGreeting(greeting, name)
  fmt.Println(name) // Stacey
}

func sayGreeting(greeting, name string) {
  fmt.Println(greeting, name) // Hello Stacey
  name = "Ted"
  fmt.Println(name) // Ted
}
```

- Manipulate parameters

```go
func main() {
  greeting := "Hello"
  name := "Stacey"
  sayGreeting(&greeting, &name)
  fmt.Println(name) // Ted
}

func sayGreeting(greeting, name *string) {
  fmt.Println(&greeting, &name) // Hello Stacey
  *name = "Ted"
  fmt.Println(name) // Ted
}
```

- Passing pointer --> faster
- ...

```go
sum("The sum is", 1, 2, 3, 4, 5)
func sum(msg string, values ...int) {
  fmt.Println(values) // [1 2 3 4 5]
  result := 0
  for _, v := range values {
    result += v
  }
  fmt.Println(msg, result)
}
```

## Return values

```go
s := sum(1, 2, 3, 4, 5)
fmt.Println("The sum is ", s)
func sum(values ...int) {
  fmt.Println(values) // [1 2 3 4 5]
  result := 0
  for _, v := range values {
    result += v
  }
  return result
}
```

- Can return addresses of local variables
  - Automatically promoted from local memory(stack) to shared memory(heap)

```go
s := sum(1, 2, 3, 4, 5)
fmt.Println("The sum is ", *s)
func sum(values ...int) *int {
  fmt.Println(values) // [1 2 3 4 5]
  result := 0
  for _, v := range values {
    result += v
  }
  return &result // not freed --> heap
}
```

- Multi return

```go
func main() {
  d, err := divide(5.0, 0.0)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(d)
}

func divide(a, b float64) (float64, error) {
  if b == 0.0 {
    return 0.0 fmt.Errorf("Cannot divide by zero")
  }
  return a/b, nil
}
```

## Anonymous functions

```go
func () {
  msg := "Hello Go!"
  fmt.Println(msg)
}() // invoke this func
```

## Function as types

```go
var f func() = func() {
  fmt.Println("Hello Go!")
}
f()
```

## Methods

```go
func main() {
  g := greeter {
    greeting: "Hello",
    name: "Go"
  }
  g.greet()
  fmt.Println("The new name is:", g.name) // Go
}

type greeter struct {
  greeting string
  name string
}

func (g greeter) greet() {
  fmt.Println(g.greeting, g.name) // Hello Go
  g.name = "" // change on the copy not the object itself
}

func (g *greeter) greet() {
  fmt.Println(g.greeting, g.name) // Hello Go
  g.name = "" // change on the object itself
}
```
