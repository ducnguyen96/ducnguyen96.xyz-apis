## Boolean

- Values are true or false
- Not an alias for other types (e.g. int)
- Every variable are initialize with zero value (bool --> false)

```go
func main() {
  n := 1 == 1
  m := 1 ==  2
  fmt.Printf("%v , %T\n", n, n) // true, bool
  fmt.Printf("%v , %T\n", m, m) // false, bool
}
```

## Numeric

```go
// signed intergers
// int8   -128 - 127
// int16  -32,768 - 32,767
// int32  -2,147,483,648 - 2,147,483,647
// int64  -9,223,372,036,854,775,808 - 9,223,372,036,854,775,807

// uint8  0 - 255
// uint16 0 - 65,535
// uint32 0 - 4,294,967,295
```

- Bit operations

```go
a := 10 // 1010
b := 3  // 0011
fmt.Println(a & b) // 0010 --> 2
fmt.Println(a | b) // 1011 --> 11
fmt.Println(a ^ b) // 1001 --> 9
fmt.Println(a &^ b) // 0100 --> 8
```

- Bit shifting

```go
a := 8 // 2^3
fmt.Println(a << 3) // 2^3 * 2^3 = 2^6 = 64
fmt.Println(a >> 3) // 2^3 / 2^3 = 2^0 = 1
```

- Floating point number

```go
float32
float64
```

## Complex types

## String

```go
s := "this is a string"
fmt.Printf("%v, %T\n", s[2], s[2]) // 105, uint8

fmt.Printf("%v, %T\n", string[2], s[2]) // i, uint8
```

- string is generally immutable

```go
s := "this is a string"
s[2] = "u"
fmt.Printf("%v, %T\n", s, s) // error
```

- Can be concatenated with plus(+) operator
- Can be converted to []byte

```go
s := "this is a string"
b := []byte(s)
fmt.Printf("%v, %T\n", b, b) // [116 104 105 ... 103], []uint8
```

## Rune : int32

```go
var r rune = 'a'
fmt.Printf("%v, %T\n", r, r) // 97, int32

```
