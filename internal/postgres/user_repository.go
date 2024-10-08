package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/kosalnik/keeper/internal/entity"
	"github.com/kosalnik/keeper/internal/log"
)

type UserRepository struct {
	db *Conn
}

func NewUserRepository(conn *Conn) *UserRepository {
	return &UserRepository{
		db: conn,
	}
}

func (r *UserRepository) Create(ctx context.Context, m *entity.User) error {
	if m.ID == uuid.Nil {
		return errors.New("User.ID is nil")
	}
	log.Info(fmt.Sprintf("%s, %s, %s", m.ID, m.Password, m.Login))
	_, err := r.db.ExecContext(ctx, `INSERT INTO "user" (id, login, pass) VALUES ($1, $2, $3)`, m.ID, m.Login, m.Password)
	return err
}

func (r *UserRepository) Get(ctx context.Context, id uuid.UUID, m *entity.User) error {
	if m == nil {
		return errors.New("model should not be nil")
	}
	row := r.db.QueryRowContext(ctx, `SELECT id, login, pass FROM "user" WHERE id = $1`, id)
	if err := row.Err(); err != nil {
		if errors.Is(row.Err(), sql.ErrNoRows) {
			return nil
		}
		return err
	}
	return row.Scan(&m.ID, &m.Login, &m.Password)
}

func (r *UserRepository) GetByLogin(ctx context.Context, login string, m *entity.User) error {
	if m == nil {
		return errors.New("model should not be nil")
	}
	row := r.db.QueryRowContext(ctx, `SELECT id, login, pass FROM "user" WHERE login = $1`, login)
	if err := row.Err(); err != nil {
		if errors.Is(row.Err(), sql.ErrNoRows) {
			return nil
		}
		return err
	}
	return row.Scan(&m.ID, &m.Login, &m.Password)
}
