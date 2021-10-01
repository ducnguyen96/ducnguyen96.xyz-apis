# Creating goroutines

```go
func main(){
  sayHello() // Hello
  go sayHello() // spin off green thread --> hello function will be run in there
  // green thread : abstraction of a thread : go routines
  // in go runtime, scheduler map these go routines onto these os threads for periods of time
}
func sayHello()  {
  fmt.Println("Hello")
}
```

# Synchronization

- use `sync.WaitGroup` to wait for groups of goroutines to complete
- use `sync.Mutex` and `sync.RWMutex` to protect data access

## WaitGroups

```go
var wg = sync.WaitGroup{}

var msg = "Hello"
wg.Add(1)
go func(msg string) {
  fmt.Println(msg)
  wg.Done()
}(msg)
msg = "Goodbye"
wg.Wait()
```

## Mutexes

```go
var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main(){
  runtime.GOMAXPROCS(100)
  for i := 0; i< 10; i++ {
    wg.Add(2)
    m.RLock()
    go sayHello()
    m.Lock()
    go increment()
  }
  wg.Wait()
}

func sayHello() {
  fmt.Printf("Hello #%v\n", counter)
  wg.Done()
}

func increment(){
  counter++
  wg.Done()
}
```

## Best practices

- Don't create goroutines in libraries
  - Let consumer control concurrency
- When creating a goroutine, know how it will end
  - Avoids subtle memory leaks
- Check for race conditions `go run -race src/main.go`

# Parallelism

- By default, Go will use CPU threads equal to available cores
- Change with runtime.GOMAXPROCS
- More threads can increase performace, but too many can slow it down
