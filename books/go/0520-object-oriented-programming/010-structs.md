---
Title: Structs
Id: 282
Score: 1
SOId: 27425
---
Go supports user defined types in the form of structs and type aliases. structs are composite types, the component pieces of data that constitute the struct type are called *fields*. a field has a type and a name which must be unqiue.

```go
package main

type User struct {
    ID uint64
    FullName string
    Email    string
}

func main() {
    user := User{
        1,
        "Zelalem Mekonen",
        "zola.mk.27@gmail.com",
    }

    fmt.Println(user) // {1 Zelalem Mekonen zola.mk.27@gmail.com}
}
```

this is also a legal syntax for definining structs

```go
type User struct {
    ID uint64
    FullName, Email string
}

user := new(User)

user.ID = 1
user.FullName = "Zelalem Mekonen"
user.Email = "zola.mk.27@gmail.com"
```
