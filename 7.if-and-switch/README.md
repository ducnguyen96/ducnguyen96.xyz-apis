# If statements

- Initializer
- Comparison operators
- Logical operators
- Short circuiting
- Equality and floats

```go
if true {
  fmt.Println("The test if true")
}
```

- Initializer

```go
statePopulations := make(map[string]int)
statePopulations = map[string]int {
  "California": 39250017,
  "Texas": 278629596,
  "Florida": 20612346
}

if pop, ok := statePopulations["Florida"]; ok {
  fmt.Println(pop) // 20612346
}
```

## Operations

```go
number := 50
guess := 30
if guess < 1 || guess >> 100 {
  fmt.Println("The guess must be between 1 and 100 !")
}

if guess >= 1 && guess <= 100 {
  if guess < number {
    fmt.Println("Too low")
  }

  if guess > number {
    fmt.Println("Too high")
  }

  if guess == number {
    fmt.Println("You got it !")
  }
  fmt.Println(number <= guess, number >= guess, number != guess)
}
```

- Short circuiting

```go
number := 50
guess := -5
if guess < 1 || returnTrue() || guess > 100 { // guess < 1 true --> don't executed any more code
  fmt.Println("The guess must be between 1 and 100 !")
}
```

## If-else and if-else if statements

```go
number := 50
guess := 30
if guess < 1 {
  fmt.Println("The guess must be greater than 1 !")
} else if guess > 100 {
  fmt.Println("The guess must be less than 100 !")
}
else {
  if guess < number {
    fmt.Println("Too low")
  }

  if guess > number {
    fmt.Println("Too high")
  }

  if guess == number {
    fmt.Println("You got it !")
  }
  fmt.Println(number <= guess, number >= guess, number != guess)
}
```

# Switch statements

- Switching on a tag
- Cases with multiple tests
- Initializers
- Switches with no tag
- Fallthrough
- Type switches
- Breaking out early

## Simple cases

```go
switch i := 2 + 3; i {
  case 1, 5, 10:
    fmt.Println("one, five or then")
  case 2, 4, 6:
    fmt.Println("two, four or six")
  default:
    fmt.Println("another number")
}
```

## Cases with multiple tests

```go
i := 10
switch {
  case i <= 10:
    fmt.Println("less than or equal to ten")
  case i <= 20:
    fmt.Println("less than or equal to twenty")
  default:
    fmt.Println("greater than twenty")
}
```

## Falling through

```go
i := 10
switch {
  case i <= 10:
    fmt.Println("less than or equal to ten")
    fallthrough
  case i >= 20:
    fmt.Println("less than or equal to twenty") // still be executed
  default:
    fmt.Println("greater than twenty")
}
```

## Type Switches

```go
var i interface{} = 1
switch i.(type) {
  case int:
    fmt.Println("i is an int")
    break
    fmt.Println("This will not be print")
  case float64:
    fmt.Println("i is a float64")
  case string:
    fmt.Println("i is a string")
  default:
    fmt.Println("another type")
}
```
