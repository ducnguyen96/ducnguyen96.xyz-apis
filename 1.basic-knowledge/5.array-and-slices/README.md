# Arrays

- Collection of items with same type
- Fixed size
- Access via zero-based index
- `len` functions returns size of array
- Copies refer to different underlying data

## Creation

```go
grades := [3]int{97, 85, 93}
grades := [...]int{97, 85, 93}
var students [3]string
students[0] = "Lisa"
students[2] = "Ahmed"
students[1] = "Arnord"
```

## Built-in functions

## Working with arrays

```go
a := [...]int{1, 2, 3}
b := a
b[1] = 5
fmt.Println(a) // [1 2 3]
fmt.Println(b) // [1 5 3]
```

```go
a := [...]int{1, 2, 3}
b := &a
b[1] = 5
fmt.Println(a) // [1 5 3]
fmt.Println(b) // &[1 5 3]
```

# Slicse // reference type

- Backed by array

## Creation

- Slice existing array or slice
- Literal style
- Via make function

```go
a := make([]int, 10) // create slice with capacity and length == 10
a := make([]int, 10, 100) // slice with length == 10 and capacity == 100
```

## Built-in functions

- `len` function returns length of slice
- `cap` function returns length of underlying array
- `append` function to add elements to slice
  - May cause expensice copy operation if underlying array is too small

## Working with slices

- Copies refer to same underlying array
