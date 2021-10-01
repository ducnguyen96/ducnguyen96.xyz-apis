# Channel basic

- Create a channel with `make` command
  - make(chan int)
- Send message into channel
  - ch <- val
- Receive message from channel
  - val := <- ch
- Can have multiple senders and receivers

```go
var wg = sync.WaitGroup{}
ch := make(chan init)
wg.Add(2)
go func(){
  i := <- ch
  fmt.Println(i)
  wg.Done()
}()

go func(){
  ch <- 42
  wg.Done()
}()

wg.Wait( )

```

# Restricting data flow

- Channel can be cast into send-only or receive only versions
- Send-only: chan <- int

```go
go func(ch <-chan int){
  i := <- ch
  fmt.Println(i)
  wg.Done()
}(ch)
```

- Receive-only: <- chan int

```go
go func(ch chan<- int){
  ch <- 42
  wg.Done()
}(ch)
```

# Buffered channels

- Channels block sender side till receiver is available
- Block receiver side till message is available
- Can decouple sender and receiver with buffered channels
  - make(chan int, 50)
- Use buffered chanels when send and receiver have asymmetric loading

```go
ch := make(chan int, 50)
wg.Add(2)
go func(ch <-chan int){
  i := <-chan
  fmt.Println(i)
  i = <-chan
  fmt.Println(i)
  wg.Done()
}(ch)
go func(ch chan<- int) {
  ch <- 42
  ch <- 27
  wg.Done()
}(ch)
wg.Wait()
```

# Closing channels

# For ... range loops with channels

- Use to monitor channel and process messages as they arrive
- Loop exits when channel is closed

```go
ch := make(chan int, 50)
wg.Add(2)
go func(ch <-chan int){
  for i := range ch {
    fmt.Println(i)
  }
  wg.Done()
}(ch)
go func(ch chan<- int) {
  ch <- 42
  ch <- 27
  close(ch)
  wg.Done()
}(ch)
wg.Wait()
```

# Select statements

- Allows goroutine to monitor several channels at once
  - Blocks if all channels block
  - If multiple channels receive value simultaneously, behavior is undefined
