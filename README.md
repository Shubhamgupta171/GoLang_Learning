
# GoLang (Golang) – Full Step-by-Step Learning Guide & Project Development Roadmap

Go ek fast, simple aur powerful backend programming language hai.  
Ye README specially beginners ke liye banaya gaya hai jo **Go + Backend + Production API Development** ekdum step-by-step seekhna chahte hain.

Isme do main parts hain:

1. **GoLang Learning Guide (Basics → Advanced)**  
2. **Project Development Roadmap (Daily Step-by-Step Feature Building)**  

---

# =============================
# PART 1: GoLang Learning Guide  
# =============================

## 1. Install & Setup GoLang

### Mac (Homebrew)
```
brew install go
```

### ✔ Version Check
```
go version
```

---

## 2. Hello World Program

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

## 3. GoLang Basics

### Variables
```go
var a int = 10
b := 20
```

### Data Types  
- int  
- float64  
- string  
- bool  

### Conditions
```go
if a > 10 { ... } else { ... }
```

### Loops (Only `for`)
```go
for i := 0; i < 5; i++ { ... }
```

---

## 4. Functions in Go

```go
func add(a int, b int) int {
    return a + b
}
```

### Multiple Returns
```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

---

## 5. Structs and Methods

```go
type User struct {
    ID   int
    Name string
}
```

---

## 6. Slices & Maps

### Slice
```go
nums := []int{1,2,3}
```

### Map
```go
scores := map[string]int{"A": 10}
```

---

## 7. Pointers

```go
x := 10
p := &x
*p = 20
```

---

## 8. Interfaces (Polymorphism)

```go
type Speaker interface {
    Speak() string
}
```

---

## 9. Concurrency (Goroutines + Channels)

### Goroutine
```go
go func(){ fmt.Println("Hi") }()
```

### Channel
```go
ch := make(chan int)
```

### Select
```go
select { case <-ch: ... }
```

---

## 10. Simple API (net/http)

```go
http.HandleFunc("/hello", ...)
```

---

## 11. Using Gin Framework

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

## 12. Database (PostgreSQL)

Install driver:
```
go get github.com/lib/pq
```

Connect:
```go
db, _ := sql.Open("postgres", "postgres://user:pass@localhost/db?sslmode=disable")
```

---

## 13. Testing in Go

```go
func TestAdd(t *testing.T){
    if Add(2,3) != 5 {
        t.Fail()
    }
}
```

---

## 14. Recommended Folder Structure

```
project/
  cmd/
  internal/
  pkg/
  go.mod
  README.md
```

---

## 15. Practice Projects

### Beginner:
- CLI Calculator  
- TODO CLI app  

### Intermediate:
- CRUD API  
- JWT Authentication  
- URL Shortener  

### Advanced:
- API Gateway  
- Microservices  
- WebSocket Chat  

---

## 16. 4-Week GoLang Learning Roadmap  

### Week 1:
Basics, loops, slices, maps  

### Week 2:
Functions, structs, interfaces  

### Week 3:
Concurrency (goroutines + channels)  

### Week 4:
REST API + PostgreSQL + JWT Auth  

---

# =============================
# PART 2: Step-by-Step Learning + Project Feature Build Plan  
# =============================

Ye roadmap aapko **daily learning + daily project feature building** me help karega.  
Isse aap 30 days me **Go backend developer + production-ready API** build kar loge.

---

# WEEK 1 — Go Basics + Initial Project Structure

## **Day 1 — Install Go + Build First API**
**Learn:** go run, go build  
**Project:**  
- `/health` route return `{status:"ok"}`  

---

## **Day 2 — Variables + Structs**
**Learn:** Variables, functions  
**Project:**  
- `/info` static API  
- Make `User` struct  

---

## **Day 3 — Slices, Maps**
**Learn:** Slices, maps  
**Project:**  
- `/users` => return user list  

---

## **Day 4 — Interfaces + Folder Structure**
**Learn:** Interfaces  
**Project:**  
- Move logic → `services/`  
- Move routes → `handlers/`  

---

## **Day 5 — Middlewares**
**Learn:** custom middleware  
**Project:**  
- Logging middleware  
- Execution time middleware  

---

## **Day 6 — .env Config**
**Learn:** godotenv  
**Project:**  
- PORT & DB URL `.env` me rakho  

---

## **Day 7 — Refactoring**
Project cleanup

---

# WEEK 2 — Database + Authentication System

## **Day 8 — Database Setup**
**Learn:** SQL basics  
**Project:**  
- Connect SQLite/PostgreSQL  
- `users` table create  

---

## **Day 9 — CRUD API**
**Learn:** SQL queries  
**Project:**  
- Create, Get, Update, Delete  

---

## **Day 10 — Password Hashing**
**Learn:** bcrypt  
**Project:**  
- Hash password before saving  

---

## **Day 11 — JWT Authentication**
**Learn:** JWT sign/verify  
**Project:**  
- `/login` => return JWT  

---

## **Day 12 — Protected Routes**
**Learn:** Auth middleware  
**Project:**  
- `/profile` protected route  

---

## **Day 13 — Refresh Token**
**Project:**  
- `/auth/refresh`  

---

## **Day 14 — Cleanup**
Authentication service improvements.

---

# WEEK 3 — Advanced Backend Features

## **Day 15 — Redis Cache**
**Learn:** Redis client  
**Project:**  
- Cache `/users` list  

---

## **Day 16 — Rate Limiting**
**Learn:** Token bucket  
**Project:**  
- Limit login attempts  

---

## **Day 17 — Token Blacklist**
**Project:**  
- `/logout` => invalidate JWT  

---

## **Day 18 — Pagination**
**Project:**  
- `/users?page=&limit=`  

---

## **Day 19 — Goroutines**
**Project:**  
- Background job (send email)  

---

## **Day 20 — File Upload**
**Project:**  
- `/upload/profile`  

---

## **Day 21 — Search Feature**
**Project:**  
- `/users/search?q=`  

---

# WEEK 4 — Production Ready API

## **Day 22 — Unit Tests**
**Project:**  
- Test user + auth services  

---

## **Day 23 — Logging**
**Project:**  
- Add Zerolog / Zap  

---

## **Day 24 — Config System**
Yaml/JSON config loader

---

## **Day 25 — Docker**
**Project:**  
- Dockerfile + docker-compose  

---

## **Day 26 — Swagger Docs**
**Project:**  
- `/docs` auto API UI  

---

## **Day 27 — CI/CD**
GitHub Actions workflow

---

## **Day 28 — Deployment**
Deploy on Render/Railway/AWS

---

# Final Result

Agar aap ye README follow karte ho toh:

✔ Aap **Go backend developer** ban jaoge  
✔ Aapka project **production-ready** ho jayega  
✔ Aap Go + DB + Redis + Auth + Docker + CI/CD sab master kar loge  

---

# 💛 Happy Coding & Learning GoLang!
