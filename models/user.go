package models

import (
	"database/sql"
	"errors"
	"log"
	"regexp"

	"github.com/AdluAghnia/go-artic/db"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string
	Username string
	Email    string
	Password string
}

func NewUser(username, email, password string) *User {
	id := uuid.New().String()

	return &User{
		ID:       id,
		Username: username,
		Email:    email,
		Password: password,
	}
}

func GetUserByEmail(email string) (*User, error) {
	db, err := db.NewDB()
	if err != nil {
		return nil, err
	}

	var user User
	query := db.QueryRow("SELECT * FROM user WHERE email = ?", email)
	err = query.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Article Not Found")
		}
		return nil, err
	}

	return &user, nil
}

func usernameValidation(username string) bool {
	return len(username) >= 3
}

func emailValidation(email string) bool {
	// Define the regex pattern for validating an email address
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regex
	re := regexp.MustCompile(emailRegex)

	// Return true if the email matches the regex, false otherwise
	return re.MatchString(email)
}

func emailExist(email string) (bool, error) {
	db, err := db.NewDB()
	if err != nil {
		return false, err
	}
	defer db.Close()

	// Prepare the SQL query
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM user WHERE email = ?)`
	err = db.QueryRow(query, email).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	return exists, nil
}

func (u *User) ValidateRegisterUser() (bool, map[string]error) {
	errs := make(map[string]error)
	isValid := true

	if !usernameValidation(u.Username) {
		errs["username"] = errors.New("Username should contain atleast 3 characters")
		isValid = false
	}

	if !emailValidation(u.Email) {
		errs["email"] = errors.New("Email is not valid")
		isValid = false
	}

	email_exist, err := emailExist(u.Email)
	if err != nil {
		log.Println(err)
		return false, nil
	}

	if email_exist {
		errs["email"] = errors.New("Email already been used")
		isValid = false
	}

	if len(u.Password) < 6 {
		errs["password"] = errors.New("Password should atleast contain 6 characters")
		isValid = false
	}

	return isValid, errs
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (u *User) ComparePassword(password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u User) SaveUser() error {
	db, err := db.NewDB()
	if err != nil {
		return err
	}
	defer db.Close()

	hash_password, err := HashPassword(u.Password)

	_, err = db.Exec("INSERT INTO user (id, username, email, password) VALUE (?, ?, ?, ?)", u.ID, u.Username, u.Email, hash_password)
	if err != nil {
		return err
	}

	return nil
}
