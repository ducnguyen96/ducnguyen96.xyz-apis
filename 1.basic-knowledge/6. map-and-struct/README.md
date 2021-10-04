# Maps

```go
statePopulations := map[string]int {
  "California": 39250017,
  "Texas": 278629596,
  "Florida": 20612346
}
```

## What are they ?

- Collections of value types that are accessed ia keys
- Created via literals or via `make` function
- Members accessed via `[key]` syntax
  - myMap["key"] = value
- Check for presence with "value, ok" form of result
- Multiple assignments refer to same underlying data

## Creating

```go
statePopulations := make(map[string]int)
statePopulations = map[string]int {
  "California": 39250017,
  "Texas": 278629596,
  "Florida": 20612346
}
```

## Manipulation

```go
statePopulations["Georgia"] = 24657354654
fmt.Println(statePopulations["Georgia"]) // 24657354654

delete(statePopulations, "Georgia")
fmt.Println(statePopulations["Georgia"]) // 0

pop, ok := statePopulations["Florida"]
fmt.Println(pop, ok) // 20612346, true

_, ok := statePopulations["Florida"]
fmt.Println(ok) // false
```

Reference

```go
sp := statePopulations
delete(sp, "Florida")
fmt.Println(sp) // Florida removed
fmt.Println(statePopulations) // Florida removed
```

# Structs

```go
type Doctor struct {
  number int
  actorName string
  companions []string
}

func main() {
  aDoctor := Doctor{
    number: 3,
    actorName: "Jon Pertwee",
    companions: []string {
      "Liz Shaw",
      "Jo Grant"
    }
  }
}
```

## What are they ?

- Collections of disparate data types that describe a single concept
- Keyed by named fields
- Normally created as types, but anonymous structs are allowed
- Structs are value types
- No inheritance but can use composition via embedding
- Tags can be added to struct fields to describe fields

## Creating

```go
aDoctor := struct{name string}{name: "John Pertwee"}
anotherDoctor := aDocter
anotherDoctor.name = "Tom Baker"
fmt.Println(aDoctor) // { John Pertwee }
fmt.Println(anotherDoctor) // { Tom Baker }
```

```go
aDoctor := struct{name string}{name: "John Pertwee"}
anotherDoctor := &aDocter
anotherDoctor.name = "Tom Baker"
fmt.Println(aDoctor) // { Tom Baker }
fmt.Println(anotherDoctor) // &{ Tom Baker }
```

## Naming conventions

## Embedding

Golang doesnt have inheritance but does have composition instead.

```go
type Animal struct {
  Name string
  Origin string
}

type Bird struct {
  Animal
  SpeedKPH float32
  CanFly bool
}

func main() {
  b := Bird{
    Animal: Animal{Name :"Emu", Origin: "Australia"},
    SppedKPH: 48,
    CanFly: false
  }

  fmt.Println(b.Name) // Emu
}
```

## Tags

```go
type Animal struct {
  Name string `required max: "100"`
  Origin string
}

func main() {
  t := reflect.TypeOf(Animal{})
  field, _ := t.FieldByName("Name")
  fmt.Println(field.Tag)
}
```
