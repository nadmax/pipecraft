package store

import (
	"sync"
	"time"

	"github.com/nadmax/pipecraft/src/models"
)

type UserStore struct {
	mu     sync.RWMutex
	users  map[int]*models.User
	nextID int
}

func NewUserStore() *UserStore {
	return &UserStore{
		users:  make(map[int]*models.User),
		nextID: 1,
	}
}

func (s *UserStore) Create(name, email string) *models.User {
	s.mu.Lock()
	defer s.mu.Unlock()

	user := &models.User{
		ID:        s.nextID,
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	s.users[s.nextID] = user
	s.nextID++

	return user
}

func (s *UserStore) GetAll() []*models.User {
	s.mu.RLock()
	defer s.mu.RUnlock()

	users := make([]*models.User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return users
}

func (s *UserStore) GetByID(id int) (*models.User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, exists := s.users[id]
	return user, exists
}

func (s *UserStore) Update(id int, name, email string) (*models.User, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[id]
	if !exists {
		return nil, false
	}
	if name != "" {
		user.Name = name
	}
	if email != "" {
		user.Email = email
	}
	user.UpdatedAt = time.Now()
	return user, true
}

func (s *UserStore) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[id]; exists {
		delete(s.users, id)
		return true
	}
	return false
}
