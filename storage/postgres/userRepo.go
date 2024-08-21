package postgres

import (
	"context"
	"log"

	"github.com/book-shop-grpc/user_service/genproto/user_service"
	"github.com/jackc/pgx/v5"
)

type userRepo struct {
	db *pgx.Conn
}

func NewUserRepo(db *pgx.Conn) UserRepoI {
	return &userRepo{db: db}
}

// CreateUser creates a new user in the database
func (r *userRepo) CreateUser(ctx context.Context, user *user_service.User) (*user_service.User, error) {
	query := `
		INSERT INTO
			 users(
			 	user_id, 
				username, 
				phone_number, 
				password
			) VALUES ($1, $2, $3, $4) 
		RETURNING 
			user_id`
	err := r.db.QueryRow(
		ctx, query, 
		user.UserId, 
		user.Username, 
		user.PhoneNumber, 
		user.Password).Scan(&user.UserId)
	if err != nil {
		log.Println("Error creating user:", err)
		return nil, err
	}
	return user, nil
}

// ListUser retrieves a list of users from the database
func (r *userRepo) ListUser(ctx context.Context, req *user_service.GetListReq) (*user_service.GetListResp, error) {
	query := `
		SELECT 
			user_id, 
			username, 
			phone_number 
		FROM 
			users 
		LIMIT $1 
		OFFSET $2`
	rows, err := r.db.Query(
			ctx, query, 
			req.Limit, 
			req.Offset)
	if err != nil {
		log.Println("Error listing users:", err)
		return nil, err
	}
	defer rows.Close()

	var users []*user_service.User
	var count int32
	for rows.Next() {
		var user user_service.User
		if err := rows.Scan(
				&user.UserId, 
				&user.Username, 
				&user.PhoneNumber); err != nil {
			log.Println("Error scanning user row:", err)
			return nil, err
		}
		users = append(users, &user)
	}

	return &user_service.GetListResp{
		Users: users,
		Count: count}, nil
}

// GetUser retrieves a user by their ID
func (r *userRepo) GetUser(ctx context.Context, userId string) (*user_service.User, error) {
	query := `
		SELECT 
			user_id, 
			username, 
			phone_number 
		FROM 
			users 
		WHERE 
			user_id = $1`
	var user user_service.User
	err := r.db.QueryRow(ctx, query,userId).Scan(
				&user.UserId, 
				&user.Username, 
				&user.PhoneNumber)
	if err != nil {
		if err == pgx.ErrNoRows {
			log.Println("User not found:", err)
			return nil, err
		}
		log.Println("Error retrieving user:", err)
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates an existing user in the database
func (r *userRepo) UpdateUser(ctx context.Context, user *user_service.User) (*user_service.User, error) {
	query := `
		UPDATE 
			users 
		SET 
			username = $1, 
			phone_number = $2, 
			password = $3 
		WHERE 
			user_id = $4`
	_, err := r.db.Exec(
			ctx, query, 
			user.Username, 
			user.PhoneNumber, 
			user.Password, 
			user.UserId)
	if err != nil {
		log.Println("Error updating user:", err)
		return nil, err
	}
	return user, nil
}

// DeleteUser removes a user from the database by their ID
func (r *userRepo) DeleteUser(ctx context.Context, userId string) error {
	query := `
		DELETE FROM 
			users 
		WHERE user_id = $1`
	_, err := r.db.Exec(ctx, query, userId)
	if err != nil {
		log.Println("Error deleting user:", err)
		return err
	}
	return nil
}
