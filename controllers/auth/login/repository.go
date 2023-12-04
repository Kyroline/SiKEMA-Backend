package loginAuth

import (
	model "attendance-is/models"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type Repository interface {
	LoginEmailAuthRepository(email string, password string) (*model.User, string)
	LoginNimAuthRepository(nim string, password string) (*model.User, string)
	LoginNipAuthRepository(nip string, password string) (*model.User, string)
}

type repository struct {
	db *gorm.DB
}

func NewLoginAuthRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) LoginEmailAuthRepository(email string, password string) (*model.User, string) {
	var user model.User
	if err := r.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return nil, "ERR_UNEXPECTED_500"
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, "ERR_UNAUTHENTICATED_401"
	}
	return &user, ""
}

func (r *repository) LoginNimAuthRepository(nim string, password string) (*model.User, string) {
	var student model.Student
	if r.db.Model(model.Student{}).Where("nim = ?", nim).Preload("User").Find(&student).Error != nil {
		return nil, "ERR_UNAUTHENTICATED_401"
	}

	if bcrypt.CompareHashAndPassword([]byte(student.User.Password), []byte(password)) != nil {
		return nil, "ERR_UNAUTHENTICATED_401"
	}
	return &student.User, ""
}

func (r *repository) LoginNipAuthRepository(nip string, password string) (*model.User, string) {
	var lecturer model.Lecturer
	if r.db.Model(model.Lecturer{}).Where("nip = ?", nip).Preload("User").Find(&lecturer).Error != nil {
		return nil, "ERR_UNAUTHENTICATED_401"
	}
	if bcrypt.CompareHashAndPassword([]byte(lecturer.User.Password), []byte(password)) != nil {
		return nil, "ERR_UNAUTHENTICATED_401"
	}
	return &lecturer.User, ""
}
