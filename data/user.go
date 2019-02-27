package data

import "time"

type (
	// User struct
	User struct {
		ID        int
		UUID      string
		Name      string
		Email     string
		Password  string
		CreatedAt time.Time
	}
	// Session struct
	Session struct {
		ID        int
		UUID      string
		Email     string
		UserID    int
		CreatedAt time.Time
	}
)

// CreateSession creates a new session for an existing user
func (user *User) CreateSession() (session Session, err error) {
	db := db()
	defer db.Close()
	statement := "INSERT INTO sessions (uuid, email, user_id, created_at) VALUES ($1, $2, $3, $4) RETURNING id, uuid, email, user_id, created_at"
	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), user.Email, user.ID, time.Now()).Scan(&session.ID, &session.UUID, &session.Email, &session.UserID, &session.CreatedAt)
	if err != nil {
		return
	}
	return
}

// Session gets the session for an existing user
func (user *User) Session() (session Session, err error) {
	db := db()
	defer db.Close()
	session = Session{}
	err = db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = $1", user.ID).Scan(&session.ID, &session.UUID, &session.Email, &session.UserID, &session.CreatedAt)
	return
}

// Check returns whether session is valid in the DB
func (session *Session) Check() (valid bool, err error) {
	db := db()
	defer db.Close()
	err = db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = $1", session.UUID).Scan(&session.ID, &session.UUID, &session.Email, &session.UserID, &session.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if session.ID != 0 {
		valid = true
	}
	return
}

// DeleteByUUID deletes session from DB
func (session *Session) DeleteByUUID() (err error) {
	db := db()
	defer db.Close()
	statement := "DELETE FROM sessions WHERE uuid = $1"
	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(session.UUID)
	if err != nil {
		return
	}
	return
}

// User returns the user from the session
func (session *Session) User() (user User, err error) {
	db := db()
	defer db.Close()
	user = User{}
	err = db.QueryRow("SELECT id, uuid, name, email, created_at FROM users, WHERE id = $1", session.UserID).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.CreatedAt)
	return
}

// SessionDeleteAll deletes all sessions from DB
func SessionDeleteAll() (err error) {
	db := db()
	defer db.Close()
	statement := "DELETE FROM sessions"
	_, err = db.Exec(statement)
	if err != nil {
		return
	}
	return
}

// Create creates a new user, saves user info into DB
func (user *User) Create() (err error) {
	db := db()
	defer db.Close()
	statement := "INSERT INTO users (uuid, name, email, password, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, uuid, created_at"
	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), user.Name, user.Email, Encrypt(user.Password), time.Now()).Scan(&user.ID, &user.UUID, &user.CreatedAt)
	if err != nil {
		return
	}
	return
}

// Delete deletes user from DB
func (user *User) Delete() (err error) {
	db := db()
	defer db.Close()
	statement := "DELETE FROM users WHERE id = $1"
	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID)
	if err != nil {
		return
	}
	return
}

// Update updates user info in the DB
func (user *User) Update() (err error) {
	db := db()
	defer db.Close()
	statement := "UPDATE users set name = $2, email = $3 WHERE id = $1"
	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID, user.Name, user.Email)
	if err != nil {
		return
	}
	return
}

// UserDeleteAll deletes all users from DB
func UserDeleteAll() (err error) {
	db := db()
	defer db.Close()
	statement := "DELETE FROM users"
	_, err = db.Exec(statement)
	if err != nil {
		return
	}
	return
}

// Users returns all users in the DB
func Users() (users []User, err error) {
	db := db()
	defer db.Close()
	rows, err := db.Query("SELECT id, uuid, name, email, password, created_at FROM users")
	if err != nil {
		return
	}
	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return
		}
		users = append(users, user)
	}
	rows.Close()
	return
}

// UserByEmail returns a sing user given an email
func UserByEmail(email string) (user User, err error) {
	db := db()
	defer db.Close()
	user = User{}
	err = db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = $1", email).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return
	}
	return
}

// UserByUUID returns a sing user given an email
func UserByUUID(uuid string) (user User, err error) {
	db := db()
	defer db.Close()
	user = User{}
	err = db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = $1", uuid).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return
	}
	return
}
