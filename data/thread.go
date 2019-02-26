package data

import "time"

type (
	// Thread struct
	Thread struct {
		ID        int
		UUID      string
		Topic     string
		UserID    int
		CreatedAt time.Time
	}
	// Post struct
	Post struct {
		ID        int
		UUID      string
		Body      string
		UserID    int
		ThreadID  int
		CreatedAt time.Time
	}
)

// CreatedAtDate returns a formatted datetime string
func (thread *Thread) CreatedAtDate() string {
	return thread.CreatedAt.Format("Jan 2, 2006 at 3:00pm")
}

// CreatedAtDate returns a formatted datetime string
func (post *Post) CreatedAtDate() string {
	return post.CreatedAt.Format("Jan 2, 2006 at 3:00pm")
}

// NumReplies returns the number of posts in a thread
func (thread *Thread) NumReplies() (count int) {
	db := db()
	defer db.Close()
	rows, err := db.Query("SELECT count(*) FROM posts WHERE thread_id = $1", thread.ID)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	rows.Close()
	return
}

// CreateThread creates a new thread
func (user *User) CreateThread(topic string) (conv Thread, err error) {
	db := db()
	defer db.Close()
	statement := "INSERT INTO threads (uuid, topic, user_id, created_at) VALUES ($1, $2, $3, $4) RETURNING id, uuid, topic, user_id, created_at"
	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), topic, user.ID, time.Now()).Scan(&conv.ID, &conv.UUID, &conv.Topic, &conv.UserID, &conv.CreatedAt)
	if err != nil {
		return
	}
	return
}

// CreatePost creates a new post to a thread
func (user *User) CreatePost(conv Thread, body string) (post Post, err error) {
	db := db()
	defer db.Close()
	statement := "INSERT INTO posts (uuid, body, user_id, thread_id, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, uuid, body, user_id, thread_id, created_at"
	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), body, user.ID, conv.ID, time.Now()).Scan(&post.ID, &post.UUID, &post.Body, &post.UserID, &post.ThreadID, &post.CreatedAt)
	if err != nil {
		return
	}
	return
}

// Threads returns all threads in the database
func Threads() (threads []Thread, err error) {
	db := db()
	defer db.Close()
	rows, err := db.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		conv := Thread{}
		if err = rows.Scan(&conv.ID, &conv.UUID, &conv.Topic, &conv.UserID, &conv.CreatedAt); err != nil {
			return
		}
		threads = append(threads, conv)
	}
	rows.Close()
	return
}

// ThreadByUUID returns a thread by the UUID
func ThreadByUUID(uuid string) (conv Thread, err error) {
	db := db()
	defer db.Close()
	conv = Thread{}
	err = db.QueryRow("SELECT id, uuid, topic, user_id, created_at FROM threads WHERE uuid = $1", uuid).Scan(&conv.ID, &conv.UUID, &conv.Topic, &conv.UserID, &conv.CreatedAt)
	return
}

// User returns the user who started this thread
func (thread *Thread) User() (user User) {
	db := db()
	defer db.Close()
	user = User{}
	db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = $1", thread.UserID).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.CreatedAt)
	return
}

// User returns the user who wrote the post
func (post *Post) User() (user User) {
	db := db()
	defer db.Close()
	user = User{}
	db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = $1", post.UserID).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.CreatedAt)
	return
}
