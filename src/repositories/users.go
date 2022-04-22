package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// Represents a user respository
type Users struct {
	db *sql.DB
}

// NewUsersRepository creates a user repository
func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

// Create insert a user in database
func (repository Users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"INSERT INTO users (name, nick, email, passwd) VALUES (?, ?, ? ,?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Passwd)
	if err != nil {
		return 0, err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ID), nil
}

// Return users according to the filters (name or nick)
func (repository Users) Index(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick%

	lines, err := repository.db.Query(
		"SELECT id, name, nick, email, createdAt from users where name LIKE ? OR nick LIKE ?",
		nameOrNick,
		nameOrNick,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository Users) Show(ID uint64) (models.User, error) {
	lines, err := repository.db.Query(
		"SELECT id, name, nick, email, createdAt from users where id = ?",
		ID,
	)
	if err != nil {
		return models.User{}, err
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// Update a user
func (repository Users) Update(ID uint64, user models.User) error {
	statement, err := repository.db.Prepare(
		"UPDATE users SET name = ?, nick = ?, email = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, ID); err != nil {
		return err
	}

	return nil
}

// Delete a user
func (repository Users) Destroy(ID uint64) error {
	statement, err := repository.db.Prepare(
		"DELETE FROM users WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

// Find a user by email
func (repository Users) FindByEmail(email string) (models.User, error) {
	line, err := repository.db.Query(
		"SELECT id, passwd FROM users WHERE email = ?",
		email,
	)
	if err != nil {
		return models.User{}, err
	}
	defer line.Close()

	var user models.User
	if line.Next() {
		if err = line.Scan(&user.ID, &user.Passwd); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// Follow allows one user to follow another
func (repository Users) Follow(userID, followerID uint64) error {
	statement, err := repository.db.Prepare("INSERT IGNORE INTO followers (user_id, follower_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

// Unfollow allows one user to unfollow another
func (repository Users) Unfollow(userID, followerID uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM followers WHERE user_id = ? AND follower_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

// Unfollow allows one user to unfollow another
func (repository Users) SearchFollowers(userID uint64) ([]models.User, error) {
	lines, err := repository.db.Query(`
		SELECT 
			u.id,
			u.name,
			u.nick,
			u.nick,
			u.createdAt
		FROM
			users u
		INNER JOIN followers f
		ON u.id = f.follower_id
		WHERE f.user_id = ?
	`, userID)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User
	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// SearchFollowers find all users that a user follow
func (repository Users) SearchFollowing(userID uint64) ([]models.User, error) {
	lines, err := repository.db.Query(`
		SELECT 
			u.id,
			u.name,
			u.nick,
			u.nick,
			u.createdAt
		FROM
			users u
		INNER JOIN followers f
		ON u.id = f.user_id
		WHERE f.follower_id = ?
	`, userID)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User
	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// GetPasswd returns a user password by userID
func (repository Users) GetPasswd(userID uint64) (string, error) {
	line, err := repository.db.Query("SELECT passwd FROM users WHERE id = ?", userID)
	if err != nil {
		return "", err
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if err = line.Scan(&user.Passwd); err != nil {
			return "", err
		}
	}

	return user.Passwd, nil
}

// UpdatePassword change user password
func (repository Users) UpdatePassword(userID uint64, passwd string) error {
	statement, err := repository.db.Prepare("UPDATE users SET passwd = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(passwd, userID); err != nil {
		return err
	}

	return nil
}
