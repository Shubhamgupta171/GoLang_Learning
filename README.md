
# GoLang (Golang) – Step-by-Step Learning Guide 

Go ek fast, simple aur powerful programming language hai. Ye guide aapko bilkul zero se leke advance level tak step-by-step le jayegi.

---

## 🚀 1. Install & Setup GoLang

### 🖥 Mac (Homebrew)
```
brew install go
```

### ✔ Version Check
```
go version
```

---

## 📌 2. Hello World Program

`main.go`
```go
package main
import "fmt"

func main() {
    fmt.Println("Hello Go!")
}
```

### Run Program
```
go run main.go
```

### Build Executable
```
go build main.go
./main
```

---

## 📘 3. GoLang Basics

### ⭐ Variables
```go
var a int = 10
b := 20 // auto type infer
```

### ⭐ Data Types  
- int  
- float64  
- string  
- bool  

### ⭐ Conditions
```go
if a > 10 {
    fmt.Println("Greater")
} else {
    fmt.Println("Smaller")
}
```

### ⭐ Loops (Only `for`)
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

---

## 📦 4. Functions in Go

```go
func add(a int, b int) int {
    return a + b
}
```

### 🔁 Multiple Returns
```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

---

## 🧱 5. Structs and Methods

```go
type User struct {
    ID   int
    Name string
}

func (u User) Greet() string {
    return "Hello " + u.Name
}
```

---

## 📚 6. Slices & Maps

### Slice (Dynamic Array)
```go
nums := []int{1,2,3}
nums = append(nums, 4)
```

### Map (Key-Value Store)
```go
scores := map[string]int{"A": 10}
scores["B"] = 20
```

---

## 🧠 7. Pointers in Go

```go
x := 10
p := &x
*p = 20
```

---

## 🎭 8. Interfaces (Polymorphism in Go)

```go
type Speaker interface {
    Speak() string
}
```

---

## ⚡ 9. Concurrency – Go ka Superpower

### Goroutine
```go
go func() {
    fmt.Println("Running in background")
}()
```

### Channel
```go
ch := make(chan int)
go func(){ ch <- 5 }()
val := <-ch
```

### Select Statement
```go
select {
case msg := <-ch:
    fmt.Println(msg)
case <-time.After(2 * time.Second):
    fmt.Println("Timeout")
}
```

---

## 🌐 10. Simple API (net/http)

```go
http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintln(w, "Hello from API")
})
http.ListenAndServe(":8080", nil)
```

---

## 🧬 11. Using Gin Framework (Fastest)

Install:
```
go get github.com/gin-gonic/gin
```

Example:
```go
r := gin.Default()
r.GET("/ping", func(c *gin.Context){
    c.JSON(200, gin.H{"message": "pong"})
})
r.Run(":8080")
```

---

## 🗄 12. Database (PostgreSQL)

Install driver:
```
go get github.com/lib/pq
```

Connect:
```go
db, _ := sql.Open("postgres", "postgres://user:pass@localhost/db?sslmode=disable")
```

---

## 🧪 13. Testing in Go

`add_test.go`
```go
func TestAdd(t *testing.T){
    if Add(2,3) != 5 {
        t.Fail()
    }
}
```

Run:
```
go test ./...
```

---

## 📁 14. Recommended Project Structure

```
project/
  cmd/
  internal/
  pkg/
  go.mod
  README.md
```

---

## 🔥 15. Practice Projects (Highly Recommended)

### Beginner:
- CLI Calculator  
- TODO App (JSON file)

### Intermediate:
- CRUD API  
- JWT Auth API  
- URL Shortener  

### Advanced:
- API Gateway  
- Microservices + RabbitMQ / Kafka  
- Chat App using WebSockets  

---

## 🗺 16. 4-Week GoLang Roadmap 

### Week 1:
Basics, variables, loops, slices, maps

### Week 2:
Functions, structs, interfaces  
Build CLI project

### Week 3:
Concurrency (goroutines + channels)  
Worker pool project

### Week 4:
REST API + PostgreSQL + JWT Auth  
Deploy on AWS or Render

---

## ✨ Conclusion

Agar aap daily 1–2 hours GoLang practice karoge,  
toh 1 month me **solid backend developer** ban jaoge.

Happy Coding! 🚀💛
