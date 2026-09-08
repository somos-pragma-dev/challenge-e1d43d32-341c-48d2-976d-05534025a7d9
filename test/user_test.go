package test

import (
	"testing"
	"internal/model"
	"internal/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err!= nil {
		t.Fatal(err)
	}
	userRepo := repository.NewUserRepository(db)
	user := model.User{Name: "Test User", Email: "test@example.com", Password: "password", Role: "user"}
	if err := userRepo.CreateUser(&user); err!= nil {
		t.Fatal(err)
	}
}