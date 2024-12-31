# Understanding Go: Concurrency, Multi-threading, and Features

## Multi-threading vs Concurrency

- **Multi-threading**: The ability to perform multiple tasks simultaneously.
- **Concurrency**: Handling multiple tasks at once by managing their progress independently.

### Key Points for Developers

- Developers must write code carefully to prevent **conflicts** when:
  - Tasks run **in parallel**.
  - **Shared data** is being accessed.

---

## Go and Concurrency

- **Designed for Modern Hardware**: Go was created to run efficiently on multiple cores.
- **Built-in Concurrency**: Go makes concurrency **cheap** and **easy** to implement with its native support.
- **Readable Syntax**: Combines:
  - The simplicity of dynamically typed languages like **Python**.
  - The efficiency and safety of statically typed languages like **C++**.

---

## Compiled Language

- Go is a **compiled language**, meaning:
  - Go code is transformed into **machine code** by a compiler before execution.
  - This enhances performance and safety.

---

## Go Program Requirements

- Every Go program must include:
  - A file with the **`package main`** declaration.
  - A **`main` function** as the entry point for execution.

Example of a simple Go program:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

## Web API in Go

- Go makes it easy to create **web APIs** using the `net/http` package.
- A web server listens for requests on a specified port and responds with content.

### Example: Simple Web Server

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
        fmt.Fprintln(writer, "Welcome to GoTasker API!")
    })

    http.ListenAndServe(":8080", nil)
}

