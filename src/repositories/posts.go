package repositories

import (
	"api/src/models"
	"database/sql"
)

// Posts represents a posts repository
type Posts struct {
	db *sql.DB
}

// NewPostsRepository create a post repository
func NewPostsRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

// Create create a post in database
func (repository Posts) Create(post models.Post) (uint64, error) {
	statement, err := repository.db.Prepare(
		"INSERT INTO posts (title, content, autor_id) VALUES (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.AutorID)
	if err != nil {
		return 0, err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ID), nil
}

// Show one post by id
func (repository Posts) Show(postID uint64) (models.Post, error) {
	line, err := repository.db.Query(`
	SELECT p.*, u.nick FROM posts p
	inner join users u
	on u.id = p.autor_id
	WHERE p.id = ?`,
		postID,
	)
	if err != nil {
		return models.Post{}, err
	}
	defer line.Close()

	var post models.Post

	if line.Next() {
		if err = line.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AutorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AutorNick,
		); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}

// Show all posts by user and users that him follows
func (repository Posts) FindAllPosts(userID uint64) ([]models.Post, error) {
	lines, err := repository.db.Query(
		`SELECT DISTINCT p.*, u.nick FROM posts p
		INNER JOIN users u ON u.id = p.autor_id
		LEFT JOIN followers f ON p.autor_id = f.user_id
		WHERE u.id = ? 
		OR f.follower_id = ?
		ORDER BY 1 DESC`,
		userID,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var posts []models.Post

	for lines.Next() {
		var post models.Post

		if err = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AutorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AutorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

// Update a post
func (repository Posts) Update(postID uint64, post models.Post) error {
	statement, err := repository.db.Prepare("UPDATE posts SET title = ?, content = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(post.Title, post.Content, postID); err != nil {
		return err
	}

	return err
}

// Delete a post
func (repository Posts) Destroy(postID uint64) error {
	statement, err := repository.db.Prepare(
		"DELETE FROM posts WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postID); err != nil {
		return err
	}

	return nil
}

func (repository Posts) ShowPostsByUser(userID uint64) ([]models.Post, error) {
	lines, err := repository.db.Query(
		`SELECT p.*, u.nick FROM posts p
		JOIN users u ON u.id = p.autor_id
		WHERE p.autor_id = ?`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var posts []models.Post

	for lines.Next() {
		var post models.Post

		if err = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AutorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AutorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

// Like add like to a post
func (repository Posts) Like(postID uint64) error {
	statement, err := repository.db.Prepare("UPDATE posts SET likes = likes + 1 WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postID); err != nil {
		return err
	}

	return nil
}

// Unlike remove like to a post
func (repository Posts) Unlike(postID uint64) error {
	statement, err := repository.db.Prepare(
		`UPDATE posts SET likes =
			CASE
				WHEN likes > 0
			THEN 
				likes - 1
			ELSE 
				likes END
		WHERE id = ?`,
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postID); err != nil {
		return err
	}

	return nil
}
