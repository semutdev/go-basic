package repositories

import (
	"fmt"
	"go-crud/internal/models"
	"sync"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id int64) (*models.User, error)
	GetAll() ([]*models.User, error)
	Update(user *models.user) error
	Delete(id int64) error 
}

type inMemoryUserRepository struct {
	users map[int64]*models.user
	mutex sync.RWMutex
	nextID int64 
}

// Membuat constructor untuk repository
func NewInMemoryUserRrepository() UserRepository {
	return &inMemoryUserRepository{
		users: make(map[int64]*models.user),
		nextID: 1
	}
}

// membuat implementasi fungsi create
func (r *inMemoryUserRepository) Create(user *models.user) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Cek email duplikat
	for _, u := range r.users {
		if u.Email == user.Email {
			return fmt.Errorf("email %s already exists", user.Email)
		}
	}

	user.ID = r.nextID
	r.users[user.ID] = user
	r.nextID++
	return nil 
}


