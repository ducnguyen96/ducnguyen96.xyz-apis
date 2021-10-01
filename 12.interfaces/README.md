# Basics

```go
type Write interface {
  Write([]byte) (int, error)
}
type ConsoleWriter struct {}

func (cw ConsoleWriter) Write(data []byte)(int, error){
  n, err := fmt.Println(string(data))
  return n, err
}
```

# Composing interfaces

```go
type Write interface {
  Write([]byte) (int, error)
}

type Closer interface {
  Close() error
}

type WriterCloser interface {
  Writer
  Closer
}
```

# Type conversion

```go
var wc WriterCloser = NewBufferedWriterCloser()
bwc := wc.(*BufferedWriterCloser)
```

## The empty interface

- Them empty interface and type switches

```go
var i interface{} = 0
```

## Type switches

# Implementing with values vs. pointers

```go
- Method set of value is all methods with value receivers
- Method set of pointer is all methods, regardless of receiver types.
```

# Best practices

- Use many, small interfaces
  - Single method interfaces are some of the most powerful and flexible
    - io.Writer, io.Reader, interface{}
- Don't export interfaces for types that will be consumed
- Do export interfaces for types that will be used be package
- Design functions and methods to receive interfaces whenever possible
