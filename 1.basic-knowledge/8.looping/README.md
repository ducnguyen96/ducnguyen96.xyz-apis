# For statements

## Simple loops

```go
for i := 0; i < 5; i++ {
  fmt.Println(i)
}
```

```go
i := 0
for ; i < 5; i++ {
  fmt.Println(i)
}
```

```go
for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
  fmt.Println(i, j)
}
```

## Exiting early

```go
i := 0
for ; i < 5; i++ {
  if i%2 == 0 {
    break
  }
}
```

```go
i := 0
for ; i < 5; i++ {
  if i%2 == 0 {
    continue
  }
}
```

```go
Loop:
  for ; i < 5; i++ {
    if i%2 == 0 {
      break Loop
    }
  }
```

## Looping through collection

```go
s := []int{1, 2, 3}
for k, v := range s {
  fmt.Println(k, v)
}
```
