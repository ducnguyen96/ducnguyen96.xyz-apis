## Creating pointer

- Pinter types use an asterisk(\*) as a prefix to type pointed to
  - \*int - a pointer to an inter
- Use the addressof operator (&) to get address of variable
- Can use the addressof operator (&) if value type already exists
  - ms := myStruct{foo: 42}
  - p := &ms
- Use addressof operator before initializer
  - &myStruct{foo: 42}
- Use the new keyword
  - Can't initialize fields at the same time

```go
a := 42
b := a // deep copy
fmt.Println(a, b) // 42, 42
```

```go
var a int = 42
var b *int = &a
fmt.Println(a, b) // 42 0x104a124
```

## Dereferencing pointers

- Deference a pointer by preceding with an asterisk(\*)
- Complex types (e.g. structs) are automatically dereferenced

```go
var a int = 42
var b *int = &a
fmt.Println(a, b) // 42 0x104a124
fmt.Println(a, *b) // 42 42
a = 27
fmt.Println(a, *b) // 27 27
*b = 14
fmt.Println(a, *b) // 14 14
```

## The new function

```go
func main() {
  var ms *myStruct
  ms = new(myStruct)
  (*ms).foo = 42 // ms.foo = 42
  fmt.Println((*ms).foo) // ms.foo
}

type myStruct struct {
  foo int
}
```

## Working with nil

```go
var ms *myStruct
fmt.Println(ms) // <nil>
```

## Types with internal pointers

- All assignment operations in Go are copy operations
- Slices and maps contain internal pointers, so copies point to same underlying data

```go
// array
a := [3]int{1, 2, 3}
b := a
fmt.Println(a, b) // [1 2 3] [1 2 3]
a[1] = 42
fmt.Println(a, b) // [1 42 3] [1 2 3]
```

```go
// slice
a := []int{1, 2, 3}
b := a
fmt.Println(a, b) // [1 2 3] [1 2 3]
a[1] = 42
fmt.Println(a, b) // [1 42 3] [1 42 3]
```

```go
// map
a := map[string]string {"foo": "bar", "bar": "buz"}
b := a
fmt.Println(a, b) // map[foo: bar baz: buz] map[foo: bar baz: buz]
a["foo"] = "qux"
fmt.Println(a, b) // map[foo: qux baz: buz] map[foo: qux baz: buz]
```
