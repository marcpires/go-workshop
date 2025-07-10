---
marp: true
paginate: true
footer: '(c) 2025 Marcelo da Silva Pires'
---

# **LHC Go workshop**

![bg left:40% 80%](https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Aqua.png)

---

# Agenda

![bg left:40% 80%](https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Aqua.png)

- Revisao do feed-reader
   - Identificadores exportados e não exportados
   - Mais erros comuns
- Tipos da linguagem
  - struct
    - tags
    - struct embedding
    - problemas com struct embedding

---

# Struct
![bg left:40% 80%](https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Aqua.png)

Permite a criacao de tipos personalizados pelo desenvolvedor.

Vejamos o pacote search.
[feed.go](../internal/search/feed.go)

---

![bg left:40% 80%](https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Aqua.png)

# Struct

Podemos declarar uma struct de algumas formas:

1. struct literal

```go
type user struct {
	name string
	email string
	ext int
    admin bool
}

marcp := user{
  name: "Marc Pires",
  email: "marcpiresrj@gmail.com",
  ext: 123,
  admin: true,

}
```
---
![bg left:40% 80%](https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Aqua.png)

# Struct

2. Declarando uma struct apenas com os valores

```go
bruno := User{"Bruno", "bruno@lhc.net.br", 223, false}
```
Aqui a ordem ao especificar os valores importa. 

---
# Struct

![bg left:40% 80%](https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Aqua.png)

Altere a linha 31 do arquivo [main.go](../cmd/langtypes/main.go), como o exemplo abaixo e veja o que acontece ao executar o programa.

```go
bruno := User{223, "Bruno", false, "bruno@lhc.net.br"}
```
**Dica**: Execute o programa via debug na sua IDE preferida

---
![bg left:40% 80%](https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Aqua.png)

# Type embedding

Quando uma `struct` possuí um campo não nomeado, definimos esse campo como `embedded`. Veja [cmd/embedded/main.go](../cmd/embedded/main.go)

```go
type Foo struct {
    Bar
}

type Bar struct {
    Baz int
}

foo := Bar{}
foo.Baz = 23
```
---
![bg left:40% 80%](https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Aqua.png)

# Type embedding

O campo Baz é promovido parac `Foo` e pode ser acessado de duas formas:

```go
func main() {
	foo := Foo{}
	foo.Baz = 23

	fmt.Printf("foo.Baz value is %d is promoted", foo.Baz)
	fmt.Printf("foo.Bar.Baz value is %d", foo.Bar.Baz)
}

```
---

![bg left:40% 80%](https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Aqua.png)

# Type embedding

Vejamos um pouco mais sobre alguns problemas que podem ocorrer com o mau uso de [*type enbedding*](../cmd/embedded/main.go)

---

![bg left:40% 80%](https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Aqua.png)

# Problemas com Type embedding

```go
package inmem

import "sync"

type MemData struct {
	sync.Mutex // Mutex é promovido para MemData
	memo map[string]int
}
```
---

![bg left:40% 80%](https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Aqua.png)

# Problemas com Type embedding

No exemplo, temos sync.Mutex sendo provido para MemData, com isso conseguimos acessar os métodos `Lock()` e `UnLock()` diretamente.

---

![bg left:40% 80%](https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Aqua.png)

```go
package main

import (
	"fmt"

	"github.com/marcpires/rss/pkg/inmem"
)

func embedMisuse() {
	m := inmem.New()
	m.Lock() // Oops ! deadlock
	m.Get("teste")
}

func main() {A
	embedMisuse()
}
```