# go-core

My core stuff with types, code, etc that I use almost everywhere

## Slice generics

These functions were added before the Go standard library added generics support for slices in the [slices](https://pkg.go.dev/golang.org/x/exp/slices) package. While some of these functions are now available in the standard library, others are still not available. Wherever possible, prefer using the standard library functions (my functions will revert to using the standard library functions when possible).

### Contains

You can check if a slice contains a value with [Contains](https://pkg.go.dev/github.com/gildas/go-core#Contains) method:

```go
slice := []int{1, 2, 3, 4, 5}

fmt.Println(core.Contains(slice, 4)) // true
fmt.Println(core.Contains(slice, 6)) // false
```

`Contains` works with all `comparable` types.

If the slice type is more complex, you can use [ContainsWithFunc](https://pkg.go.dev/github.com/gildas/go-core#ContainsWithFunc) method:

```go
type User struct {
  ID int
  Name string
}

slice := []User{
  {ID: 1, Name: "John"},
  {ID: 2, Name: "Jane"},
  {ID: 3, Name: "Bob"},
}

fmt.Println(core.ContainsWithFunc(slice, User{Name: "John"}, func(item, john User) bool {
  return item.Name == john.Name
})) // true
```

If the struct implements the [core.Named](https://pkg.go.dev/github.com/gildas/go-core#Named) interface, you can use the [core.MatchNamed](https://pkg.go.dev/github.com/gildas/go-core#MatchNamed) method:

```go
fmt.Println(core.ContainsWithFunc(slice, User{Name: "John"}, core.MatchNamed)) // true
```

Same goes for [core.Identifiable](https://pkg.go.dev/github.com/gildas/go-core#Identifiable) and [core.StringIdentifiable](https://pkg.go.dev/github.com/gildas/go-core#StringIdentifiable) interfaces.

### Find

You can find a value in a slice with [Find](https://pkg.go.dev/github.com/gildas/go-core#Find) method:

```go
slice := []int{1, 2, 3, 4, 5}
number, found := core.Find(slice, 4)
```

[Find](https://pkg.go.dev/github.com/gildas/go-core#Find) works with all `comparable` types.

If the slice type is more complex, you can use [FindWithFunc](https://pkg.go.dev/github.com/gildas/go-core#FindWithFunc) method:

```go
type User struct {
  ID int
  Name string
}

slice := []User{
  {ID: 1, Name: "John"},
  {ID: 2, Name: "Jane"},
  {ID: 3, Name: "Bob"},
}

user, found := core.FindWithFunc(slice, User{Name: "John"}, func(a, b User) bool {
  return a.Name == b.Name
})
```

If the struct implements the [core.Named](https://pkg.go.dev/github.com/gildas/go-core#Named) interface, you can use the [core.MatchNamed](https://pkg.go.dev/github.com/gildas/go-core#MatchNamed) method:

```go
fmt.Println(core.FindWithFunc(slice, User{Name: "John"}, core.MatchNamed)) // true
```

Same goes for [core.Identifiable](https://pkg.go.dev/github.com/gildas/go-core#Identifiable) and [core.StringIdentifiable](https://pkg.go.dev/github.com/gildas/go-core#StringIdentifiable) interfaces.

### EqualSlices

You can check if two slices are equal with [EqualSlices](https://pkg.go.dev/github.com/gildas/go-core#EqualSlices) method:

```go
slice1 := []int{1, 2, 3, 4, 5}
slice2 := []int{1, 2, 3, 4, 5}

fmt.Println(core.EqualSlices(slice1, slice2)) // true
```

[EqualSlices](https://pkg.go.dev/github.com/gildas/go-core#EqualSlices) works with all `comparable` types.

If the slice type is more complex, you can use [EqualSlicesWithFunc](https://pkg.go.dev/github.com/gildas/go-core#EqualSlicesWithFunc) method:

```go
type User struct {
  ID int
  Name string
}

slice1 := []User{
  {ID: 1, Name: "John"},
  {ID: 2, Name: "Jane"},
  {ID: 3, Name: "Bob"},
}

slice2 := []User{
  {ID: 1, Name: "John"},
  {ID: 2, Name: "Jane"},
  {ID: 3, Name: "Bob"},
}

fmt.Println(core.EqualSlicesWithFunc(slice1, slice2, func(a, b User) bool {
  return a.Name == b.Name
})) // true
```

### Filter

You can filter a slice with [Filter](https://pkg.go.dev/github.com/gildas/go-core#Filter) method:

```go
slice := []int{1, 2, 3, 4, 5}

fmt.Println(core.Filter(slice, func(element int) bool {
  return element > 3
})) // [4 5]
```

### Join

You can join a slice with [Join](https://pkg.go.dev/github.com/gildas/go-core#Join) method:

```go
slice := []int{1, 2, 3, 4, 5}

fmt.Println(core.Join(slice, ",")) // 1,2,3,4,5
```

If the slice type is more complex, you can use [JoinWithFunc](https://pkg.go.dev/github.com/gildas/go-core#JoinWithFunc) method:

```go
type User struct {
  ID int
  Name string
}

slice := []User{
  {ID: 1, Name: "John"},
  {ID: 2, Name: "Jane"},
  {ID: 3, Name: "Bob"},
}

fmt.Println(core.JoinWithFunc(slice, ",", func(element User) string {
  return element.Name
})) // John,Jane,Bob
```

### Map

You can map a slice with [Map](https://pkg.go.dev/github.com/gildas/go-core#Map) method:

```go
slice := []int{1, 2, 3, 4, 5}

fmt.Println(core.Map(slice, func(element int) int {
  return element * 2
})) // [2 4 6 8 10]
```

The returned slice is a new slice, the original slice is not modified.

### Reduce

You can reduce a slice with [Reduce](https://pkg.go.dev/github.com/gildas/go-core#Reduce) method:

```go
slice := []int{1, 2, 3, 4, 5}

fmt.Println(core.Reduce(slice, func(accumulator, element int) int {
  return accumulator + element
})) // 15
```

### Sort

You can sort a slice with [Sort](https://pkg.go.dev/github.com/gildas/go-core#Sort) method:

```go
slice := []int{5, 2, 3, 1, 4}

fmt.Println(core.Sort(slice, func(a, b int) {
  return a < b
})) // [1 2 3 4 5]
```

The [Sort](https://pkg.go.dev/github.com/gildas/go-core#Sort) method is an in-place sort, the original slice is modified.

The [Sort](https://pkg.go.dev/github.com/gildas/go-core#Sort) method is now calling the [slices.SortFunc](https://pkg.go.dev/golang.org/x/exp/slices#SortFunc) from the standard library.

### Convert to/from []any

You can convert a slice to/from `[]any` with [ToAnySlice](https://pkg.go.dev/github.com/gildas/go-core#ToAnySlice) and [FromAnySlice](https://pkg.go.dev/github.com/gildas/go-core#FromAnySlice) methods:

```go
slice := []int{1, 2, 3, 4, 5}
anySlice := core.ConvertToAnySlice(slice)
fmt.Println(anySlice) // [1 2 3 4 5]
newSlice := core.ConvertFromAnySlice[int](anySlice)
fmt.Println(newSlice) // [1 2 3 4 5]
```

If the []any slice came from a JSON unmarshalling, [FromAnySlice](https://pkg.go.dev/github.com/gildas/go-core#FromAnySlice) will marshal/unmarshal each element to convert it to the target type.

```go
type Person struct {
  Name string
  Age  int
}
var anySlice []any

_ = json.Unmarshal([]byte(`[{"name": "Alice", "age": 30}, {"name": "Bob", "age": 25}]`), &anySlice)
personSlice := core.ConvertFromAnySlice[Person](anySlice)
fmt.Println(personSlice) // [{Alice 30} {Bob 25}]
```

## Time and Duration helpers

The [core.Time](https://pkg.go.dev/github.com/gildas/go-core#Time) mimics the [time.Time](https://pkg.go.dev/time#Time) and adds JSON serialization support to and from RFC 3339 time strings.

The [core.Duration](https://pkg.go.dev/github.com/gildas/go-core#Duration) mimics the [time.Duration](https://pkg.go.dev/time#Duration) and adds JSON serialization support to and from duration strings. Its [core.ParseDuration](https://pkg.go.dev/github.com/gildas/go-core#ParseDuration) also understands most of the ISO 8601 duration formats. It marshals to milliseconds. It can unmarshal from milliseconds, GO duration strings, and ISO 8601 duration strings.

Example:

```go
type User struct {
  Name      string
  CreatedAt core.Time
  Duration  core.Duration
}

user := User{
  Name:      "John",
  CreatedAt: core.Now(),
  Duration:  core.Duration(5 * time.Second),
}

data, err := json.Marshal(user)
if err != nil {
  panic(err)
}

fmt.Println(string(data))
// {"Name":"John","CreatedAt":"2021-01-01T00:00:00Z","Duration":"5000"}

string.Replace(string(data), "5000", "PT5S", 1)

var user2 User
err = json.Unmarshal(data, &user2)
if err != nil {
  panic(err)
}
fmt.Println(user2.Duration) // 5s
```

The [core.Timestamp](https://pkg.go.dev/github.com/gildas/go-core#Timestamp) type is an alias for [core.Time](https://pkg.go.dev/github.com/gildas/go-core#Time) and it is used to represent timestamps in milliseconds. It marshals into milliseconds and unmarshals from milliseconds (string or integer).

## Environment Variable helpers

You can get an environment variable with `GetEnvAsX` methods, where `X` is one of `bool`, [time.Duration](https://pkg.go.dev/time#Duration), `int`, `string`, [time.Time](https://pkg.go.dev/time#Time), [url.URL](https://pkg.go.dev/net/url#URL), [uuid.UUID](https://pkg.go.dev/github.com/google/uuid#UUID), if the environment variable is not set or the conversion fails, the default value is returned.

Example:

```go
// GetEnvAsBool
v := core.GetEnvAsBool("ENV_VAR", false)
// GetEnvAsDuration
v := core.GetEnvAsDuration("ENV_VAR", 5 * time.Second)
// GetEnvAsString
v := core.GetEnvAsString("ENV_VAR", "default")
// GetEnvAsTime
v := core.GetEnvAsTime("ENV_VAR", time.Now())
// GetEnvAsURL
v := core.GetEnvAsURL("ENV_VAR", &url.URL{Scheme: "https", Host: "example.com"})
// GetEnvAsUUID
v := core.GetEnvAsUUID("ENV_VAR", uuid.New())
```

**Notes**:

- [GetEnvAsBool](https://pkg.go.dev/github.com/gildas/go-core#GetEnvAsBool) returns `true` if the environment variable is set to `true`, `1`, `on` or `yes`, otherwise it returns `false`.
  It is also case-insensitive.  
- [GetEnvAsDuration](https://pkg.go.dev/github.com/gildas/go-core#GetEnvAsDuration) accepts any duration string that can be parsed by [core.ParseDuration](https://pkg.go.dev/github.com/gildas/go-core#ParseDuration), see below in the [Time and Duration Helpers](#time-and-duration-helpers) section.  
- [GetEnvAsTime](https://pkg.go.dev/github.com/gildas/go-core#GetEnvAsTime) accepts an RFC 3339 time string.  
- [GetEnvAsURL](https://pkg.go.dev/github.com/gildas/go-core#GetEnvAsURL) fallback can be a `url.URL`, a `*url.URL`, or a `string`.

## Common Interfaces

The [core.Identifiable](https://pkg.go.dev/github.com/gildas/go-core#Identifiable) interface is used to represent an object that has an ID in the form of a `uuid.UUID`.

The [core.StringIdentifiable](https://pkg.go.dev/github.com/gildas/go-core#StringIdentifiable) interface is used to represent an object that has an ID in the form of a `string`.

The [core.Named](https://pkg.go.dev/github.com/gildas/go-core#Named) interface is used to represent an object that has a name in the form of a `string`.

The [core.TypeCarrier](https://pkg.go.dev/github.com/gildas/go-core#TypeCarrier) interface is used to represent an object that has a type in the form of a `string`.

The [core.IsZeroer](https://pkg.go.dev/github.com/gildas/go-core#IsZeroer) interface is used to represent an object that can be checked if it is zero. This interace is compatible with the `IsZero` method used by the [encoding/json](https://pkg.go.dev/encoding/json) package.

The [core.GoString](https://pkg.go.dev/github.com/gildas/go-core#GoString) interface is used to represent an object that can be converted to a Go string.

## HTTP Response helpers

[core.RespondWithJSON](https://pkg.go.dev/github.com/gildas/go-core#RespondWithJSON) is a helper function that marshals a payload into an `http.ResponseWriter` as JSON. It also sets the `Content-Type` header to `application/json`.

[core.RespondWithHTMLTemplate](https://pkg.go.dev/github.com/gildas/go-core#RespondWithHTMLTemplate) is a helper function that executes a template on a given data and writes the result into an `http.ResponseWriter`. It also sets the `Content-Type` header to `text/html`.

[core.RespondWithError](https://pkg.go.dev/github.com/gildas/go-core#RespondWithError) is a helper function that marshals an error into an `http.ResponseWriter` as JSON. It also sets the `Content-Type` header to `application/json`.

The [core.GetReference](https://pkg.go.dev/github.com/gildas/go-core#GetReference) function gets a reference of an object, if the object implements [core.Identifiable](https://pkg.go.dev/github.com/gildas/go-core#Identifiable), [core.StringIdentifiable](https://pkg.go.dev/github.com/gildas/go-core#StringIdentifiable), or [fmt.Stringer](https://pkg.go.dev/fmt#Stringer), the reference will use the ID, otherwise it will return the object itself.

Example:

```go
type User struct {
  ID uuid.UUID
  Name string
}
func (user User) GetID() uuid.UUID {
  return user.ID
}

core.RespondWithJSON(w, http.StatusAccepted, core.GetReference(user)) // this will send a response like {"id": "12345678-1234-5678-1234-567812345678"}
```

The [core.Decorate](https://pkg.go.dev/github.com/gildas/go-core#Decorate) function gets a decorated object of an object, if the object implements [core.Identifiable](https://pkg.go.dev/github.com/gildas/go-core#Identifiable), [core.StringIdentifiable](https://pkg.go.dev/github.com/gildas/go-core#StringIdentifiable), or [fmt.Stringer](https://pkg.go.dev/fmt#Stringer), the decorated object will contain the object and a `selfURI` field.

Example:

```go
type User struct {
  ID uuid.UUID
  Name string
}

func (user User) GetID() uuid.UUID {
  return user.ID
}

core.RespondWithJSON(w, http.StatusAccepted, core.Decorate(user, "/users"))
// this will send a response like:
//  {"id": "12345678-1234-5678-1234-567812345678", "selfURI": "/users/12345678-1234-5678-1234-567812345678"}
```

## Type Registries

[core.TypeRegistry](https://pkg.go.dev/github.com/gildas/go-core#TypeRegistry) is a type registry that can be used to unmarshal JSON [core.TypeCarrier](https://pkg.go.dev/github.com/gildas/go-core#TypeCarrier) objects into the correct type:

```go
type User struct {
  ID   uuid.UUID
  Name string
}

func (user User) GetType() string {
  return "user"
}

type Product struct {
  ID   uuid.UUID
  Name string
}

func (product Product) GetType() string {
  return "product"
}

registry := core.TypeRegistry{}

registry.Add(User{}, Product{})

var user User

err := registry.UnmarshalJSON([]byte(`{"type": "user", "ID": "00000000-0000-0000-0000-000000000000", "Name": "John"}`), &user)
if err != nil {
  panic(err)
}

fmt.Println(user)

registry := core.CaseInsensitiveTypeRegistry{}

registry.Add(User{}, Product{})

var user User

err := registry.UnmarshalJSON([]byte(`{"type": "UsEr", "ID": "00000000-0000-0000-0000-000000000000", "Name": "John"}`), &user)
if err != nil {
  panic(err)
}

fmt.Println(user)
```

**Notes**:

- The default JSON property name for the type is `type`, but it can be changed by adding strings to the UnmarshalJSON method.

## Miscellaneous

[core.ExecEvery](https://pkg.go.dev/github.com/gildas/go-core#ExecEvery) executes a function every `duration`:

```go
stop, ping, change := core.ExecEvery(5 * time.Second, func() {
  fmt.Println("ping")
})

time.Sleep(15 * time.Second)
change <- 10 * time.Second
time.Sleep(15 * time.Second)
stop <- true
```

**Notes**:

- `stop` is a channel that can be used to stop the execution.
- `ping` is a channel that can be used to force the execution of the func at any time.
- `change` is a channel that can be used to change the execution duration.

[core.FlexInt](https://pkg.go.dev/github.com/gildas/go-core#FlexInt), [core.FlexInt8](https://pkg.go.dev/github.com/gildas/go-core#FlexInt8), [core.FlexInt16](https://pkg.go.dev/github.com/gildas/go-core#FlexInt16), [core.FlexInt32](https://pkg.go.dev/github.com/gildas/go-core#FlexInt32), [core.FlexInt64](https://pkg.go.dev/github.com/gildas/go-core#FlexInt64) are types that can be unmarshalled from a string or an integer:

```go
type User struct {
  ID core.FlexInt
}

user := User{}
json.Unmarshal([]byte(`{"ID": 1}`), &user)
json.Unmarshal([]byte(`{"ID": "1"}`), &user)
```

[core.Must](https://pkg.go.dev/github.com/gildas/go-core#Must) is a helper function that panics if the error is not `nil` from a function that returns a value and an error:

```go
func DoSomething() (string, error) {
  return "", errors.New("error")
}

func main() {
  value := core.Must(DoSomething()).(string)
}
```

[core.URL](https://pkg.go.dev/github.com/gildas/go-core#URL) is an alias for [url.URL](https://pkg.go.dev/net/url#URL) that marshals as a string and unmarshals from a string. When unmarshaling, if the value is nil or empty, the unmarshaled value is nil (it is not considered as an error).

[core.UUID](https://pkg.go.dev/github.com/gildas/go-core#UUID) is an alias for [uuid.UUID](https://pkg.go.dev/github.com/google/uuid#UUID) that marshals as a string and unmarshals from a string. When unmarshaling, if the value is nil or empty, the unmarshaled value is `uuid.Nil` (it is not considered as an error). It also implements the [fmt.Stringer](https://pkg.go.dev/fmt#Stringer), [core.Identifiable](https://pkg.go.dev/github.com/gildas/go-core#Identifiable) and [core.IsZeroer](https://pkg.go.dev/github.com/gildas/go-core#IsZeroer) interfaces.

You can cap strings with [core.CappedString](https://pkg.go.dev/github.com/gildas/go-core#CappedString):

```go
str := "Hello, World!"
cappedStr := core.CappedString(str, 8)
fmt.Println(cappedStr) // Hello...
```

You can give your own suffix to the capped string:

```go
str := "Hello, World!"
cappedStr := core.CappedStringWithSuffix(str, 8, "***")
fmt.Println(cappedStr) // Hello***
```
