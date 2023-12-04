// usermanager.go
package usermanager

import (
	"fmt"
	"sync"
)

// User struct represents a user with a unique ID, penalty duration, and penalty amount.
type User struct {
	ID          int
	PenaltyDays int
	PenaltyAmt  float64
}

// UserManager manages a collection of users.
type UserManager struct {
	Users map[int]User
	mu    sync.Mutex
}

// NewUserManager creates a new UserManager.
func NewUserManager() *UserManager {
	return &UserManager{
		Users: make(map[int]User),
	}
}

// AddUser adds a new user to the manager.
func (um *UserManager) AddUser(userID int) {
	um.mu.Lock()
	defer um.mu.Unlock()
	um.Users[userID] = User{ID: userID}
}

// ApplyPenalty applies a penalty to a user for a specified number of days and amount.
func (um *UserManager) ApplyPenalty(userID, penaltyDays int, penaltyAmt float64) {
	um.mu.Lock()
	defer um.mu.Unlock()
	user, exists := um.Users[userID]
	if !exists {
		fmt.Println("User not found.")
		return
	}

	user.PenaltyDays = penaltyDays
	user.PenaltyAmt = penaltyAmt
	um.Users[userID] = user
}

// CheckPenalty checks if a user has a penalty and returns the details.
func (um *UserManager) CheckPenalty(userID int) User {
	um.mu.Lock()
	defer um.mu.Unlock()
	user, exists := um.Users[userID]
	if !exists {
		fmt.Println("User not found.")
		return User{}
	}
	return user
}

// PayPenalty removes the penalty for a user if they pay the required amount.
func (um *UserManager) PayPenalty(userID int, paymentAmt float64) {
	um.mu.Lock()
	defer um.mu.Unlock()
	user, exists := um.Users[userID]
	if !exists {
		fmt.Println("User not found.")
		return
	}

	if user.PenaltyDays > 0 && paymentAmt >= user.PenaltyAmt {
		user.PenaltyDays = 0
		user.PenaltyAmt = 0
		um.Users[userID] = user
	}
}
