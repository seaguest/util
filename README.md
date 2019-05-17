# util
some auxiliry utils on IMEI generator/validator, time parse/format, collection operation.

## IMEI
### [Generate](https://godoc.org/github.com/seaguest/util#Generate)
~~~ go
func Generate() string
~~~
 Generate returns random IMEI

### [Validate](https://godoc.org/github.com/seaguest/util#Validate)
~~~ go
func Validate(imei string) bool
~~~
 Validate check if imei is valid.


## Slice
### [Contains](https://godoc.org/github.com/seaguest/util#Contains)
~~~ go
func Contains(list interface{}, e interface{}) bool
~~~
 Contains check if list contains an element.

~~~ go
// EXAMPLE
var list = []int{1, 1, 2, 3}
var e int = 1

ok := Contains(list, e)
fmt.Println(ok) // true
~~~

### [FieldSlice](https://godoc.org/github.com/seaguest/util#FieldSlice)
~~~ go
func FieldSlice(list interface{}, field string) (reflect.Value, bool)
~~~
 FieldSlice returns a slice extracted by specifed field from slice.

[1, 2, 2, 4, 6] & [2, 4, 5] >> [1, 5, 6]

[1, 1, 3, 4, 5, 6] >> [1, 3, 4, 5, 6]
~~~ go
// EXAMPLE
type user struct{
    Name string
    Age uint32
}

var users = []*user{&user{Name:"hello", Age:31}}

z, ok := FieldSlice(users, "Name")
if !ok {
    fmt.Println("Cannot find field")
}

slice, ok := z.Interface().([]string)
if !ok {
    fmt.Println("Cannot convert to slice")
}
fmt.Println(slice, reflect.TypeOf(slice)) // ["hello"] []string
~~~

### [Difference](https://godoc.org/github.com/seaguest/util#Difference)
~~~ go
func Difference(arrs ...interface{}) (reflect.Value, bool)
~~~
 Difference returns a slice of values that are only present in one of the input slices

[1, 2, 2, 4, 6] & [2, 4, 5] >> [1, 5, 6]

[1, 1, 3, 4, 5, 6] >> [1, 3, 4, 5, 6]
~~~ go
// EXAMPLE
var a = []int{1, 1, 2, 3}
var b = []int{2, 4}

z, ok := Difference(a, b)
if !ok {
    fmt.Println("Cannot find difference")
}

slice, ok := z.Interface().([]int)
if !ok {
    fmt.Println("Cannot convert to slice")
}
fmt.Println(slice, reflect.TypeOf(slice)) // [1, 3, 4] []int
~~~

### [Distinct](https://godoc.org/github.com/seaguest/util#Distinct)
~~~ go
func Distinct(arr interface{}) (reflect.Value, bool)
~~~

Distinct returns the unique vals of a slice

[1, 1, 2, 3] >> [1, 2, 3]

~~~ go
// EXAMPLE
var a = []int{1, 1, 2, 3}

z, ok := Distinct(a)
if !ok {
    fmt.Println("Cannot find distinct")
}

slice, ok := z.Interface().([]int)
if !ok {
    fmt.Println("Cannot convert to slice")
}
fmt.Println(slice, reflect.TypeOf(slice)) // [1, 2, 3] []int
~~~

### [Intersect](https://godoc.org/github.com/seaguest/util#Intersect)
~~~ go
func Intersect(arrs ...interface{}) (reflect.Value, bool)
~~~

Intersect returns a slice of values that are present in all of the input slices

[1, 1, 3, 4, 5, 6] & [2, 3, 6] >> [3, 6]

[1, 1, 3, 4, 5, 6] >> [1, 3, 4, 5, 6]

~~~ go
// EXAMPLE
var a = []int{1, 1, 2, 3}
var b = []int{2, 4}

z, ok := Intersect(a, b)
if !ok {
    fmt.Println("Cannot find intersect")
}

slice, ok := z.Interface().([]int)
if !ok {
    fmt.Println("Cannot convert to slice")
}
fmt.Println(slice, reflect.TypeOf(slice)) // [2] []int
~~~

### [Union](https://godoc.org/github.com/seaguest/util#Union)
~~~ go
func Union(arrs ...interface{}) (reflect.Value, bool)
~~~

Union returns a slice that contains the unique values of all the input slices

[1, 2, 2, 4, 6] & [2, 4, 5] >> [1, 2, 4, 5, 6]

[1, 1, 3, 4, 5, 6] >> [1, 3, 4, 5, 6] 

~~~ go
// EXAMPLE
var a = []int{1, 1, 2, 3}
var b = []int{2, 4}

z, ok := Union(a, b)
if !ok {
    fmt.Println("Cannot find union")
}

slice, ok := z.Interface().([]int)
if !ok {
    fmt.Println("Cannot convert to slice")
}
fmt.Println(slice, reflect.TypeOf(slice)) // [1, 2, 3, 4] []int
~~~


## Time
### [ParseTime](https://godoc.org/github.com/seaguest/util#ParseTime)
~~~ go
func ParseTime(ts interface{}) time.Time
~~~
 ParseTime parse different time type.

### [FormatTime](https://godoc.org/github.com/seaguest/util#FormatTime)
~~~ go
func FormatTime(ts interface{}) string
~~~
 FormatTime format string to 2006-01-02 15:04:05 format


## crypto
### [AESBase64Encrypt](https://godoc.org/github.com/seaguest/util#AESBase64Encrypt)
~~~ go
func AESBase64Encrypt(clear, key, iv string) (encryoted string, err error)
~~~
 AESBase64Encrypt encrypt clear data with key and iv.

### [AESBase64Decrypt](https://godoc.org/github.com/seaguest/util#AESBase64Decrypt)
~~~ go
func AESBase64Decrypt(encrypted, key, iv string) (clear string, err error)
~~~
 AESBase64Decrypt decrypt encrypted data with key and iv.

## License
MIT