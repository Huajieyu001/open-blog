package repository

import (
	"blog-server/model"
	"database/sql"
	"time"
)

// UserRepository 用户仓储
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository 创建用户仓储
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// FindByUsername 根据用户名查找用户
func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	user := &model.User{}
	var avatar sql.NullString
	err := r.db.QueryRow(
		"SELECT id, username, password, nickname, avatar, created_at, updated_at FROM users WHERE username = ?",
		username,
	).Scan(&user.ID, &user.Username, &user.Password, &user.Nickname, &avatar, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	if avatar.Valid {
		user.Avatar = avatar.String
	}
	return user, nil
}

// FindByID 根据ID查找用户
func (r *UserRepository) FindByID(id int64) (*model.User, error) {
	user := &model.User{}
	var avatar sql.NullString
	err := r.db.QueryRow(
		"SELECT id, username, password, nickname, avatar, created_at, updated_at FROM users WHERE id = ?",
		id,
	).Scan(&user.ID, &user.Username, &user.Password, &user.Nickname, &avatar, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	if avatar.Valid {
		user.Avatar = avatar.String
	}
	return user, nil
}

// Create 创建用户
func (r *UserRepository) Create(user *model.User) error {
	now := time.Now()
	result, err := r.db.Exec(
		"INSERT INTO users (username, password, nickname, avatar, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		user.Username, user.Password, user.Nickname, user.Avatar, now, now,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = id
	user.CreatedAt = now
	user.UpdatedAt = now
	return nil
}

// Update 更新用户
func (r *UserRepository) Update(user *model.User) error {
	now := time.Now()
	_, err := r.db.Exec(
		"UPDATE users SET nickname = ?, avatar = ?, updated_at = ? WHERE id = ?",
		user.Nickname, user.Avatar, now, user.ID,
	)
	if err != nil {
		return err
	}
	user.UpdatedAt = now
	return nil
}
