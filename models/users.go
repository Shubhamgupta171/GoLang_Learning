package models

// DB Model (sqlx will map from SQLite columns)
type User struct {
    ID       int    `db:"id"`
    Name     string `db:"name"`
    Email    string `db:"email"`
    Password string `db:"password"`
}

// DTO for Register
type RegisterRequest struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

// DTO for Login
type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}
