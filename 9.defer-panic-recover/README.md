## Defer

- Used to deplay execution of a statement until func exists
- Useful to group `open` and `close` funcs together
  - Be cafeful in loops
- Run in last in first out order
- Arguments evaluated at time defer is executed, not at time of called function execution

```go
a := "start"
defer fmt.Println(a) // start
a = "end"
```

- It's similar to callback queue in javascript, will be call when stack is empty
- Last in, first out

```go
defer fmt.Println("1")
defer fmt.Println("2")
defer fmt.Println("3")
// 3 -> 2 -> 1
```

- usages

```go
res, err := http.get("http://www.google.com/robots.txt")
if err != nil {
  log.Fatal(err)
}

defer res.Body.Close()
robots, err := ioutil.ReadAll(res.Body)

if err != nil {
  log.Fatal(err)
}
fmt.Println("%s", robots)
```

## Panic

- Something went wrong --> return an error || panic
- typically we don't consider a lot of things to make application shut down. For example make http request an don't get a response --> return error not shutdown
- Occur when program cannot continue at all
  - Don't use when file can't be opended, unless it is critical
  - Use for unrecoreable events - cannot obtain TCP port for web server
- Function will stop executing
  - Deferred functions will still fire
- If nothing handles panic, program will exit

## Recover

- Used to recover from panics
- Only useful in deferred functions
- Current function will not attempt to continue, but higher functions in call stack will

```go
func main() {
  fmt.Println("start")
  panicker()
  fmt.Println("end")
}

func pacnicker() {
  fmt.Println("about to pani")
  defer func() {
    if err := recover(); err != nil {
      log.Println("Error: ", err)
      // if we put panic function here --> application shutdown
    }
  }()
  panic("Somthing bad happened !")
  fmt.Println("done panicking")
}
```
