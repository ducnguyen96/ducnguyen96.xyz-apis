## 1. Variable declaration

```go
var i int
i = 42

var j int = 42

j := 42
```

```go
var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sara Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
)

var (
	counter int = 0
)
```

## 2. Redeclaration and shadowing

- Can't redeclare variables, but can shadow them

```go
var i int = 27
func main() {
  fmt.Println(i) // 27
  var i int = 42
  fmt.Println(i) // 42
}
```

- All variables must be used

## 3. Visibility

- lowercase variables are scoped to the package
- uppercase will export it
- no private scope

## 4. Naming conventions

- length of variable should reflect the life of the variable.
- acronym should be all uppercase (`var theHTTPRequest string = https://google.com`)

## 5. Type conventions

```go
var i int = 42
fmt.Printf("%v , %T\n", i, i)  // 42, int

var j string // alias for a stream of bytes
j = string(i) // unicode of *
fmt.Printf("%v, %T\n", j, j) // *, string

j = strconv.Itoa(i) //
```
