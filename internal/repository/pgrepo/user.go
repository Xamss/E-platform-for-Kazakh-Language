package pgrepo

import (
	"context"
	"fmt"
	"strings"

	"github.com/georgysavva/scany/pgxscan"
	"xamss.diploma/internal/model"
)

func (p *Postgres) CreateUser(ctx context.Context, u *model.User) error {
	query := fmt.Sprintf(`
			INSERT INTO %s (
			                username,
			                first_name, 
			                last_name,
							email,
			                hashed_password
			                )
			VALUES ($1, $2, $3, $4, $5)
			`, "users")

	fmt.Println(u)
	_, err := p.Pool.Exec(ctx, query, u.Username, u.FirstName, u.LastName, u.Email, u.Password)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetUser(ctx context.Context, username string) (*model.User, error) {
	user := new(model.User)

	query := fmt.Sprintf("SELECT id, username, first_name, last_name, email, hashed_password FROM %s WHERE username = $1", "users")

	//rows, err := p.SQLDB.Query(query, username)
	//if err != nil {
	//	return nil, err
	//}
	//defer rows.Close()
	//
	//for rows.Next() {
	//	err := rows.Scan(&user.ID, &user.Username, &user.LastName, &user.LastName, &user.Password)
	//	if err != nil {
	//		return nil, err
	//	}
	//}
	//err = rows.Err()
	//if err != nil {
	//	return nil, err
	//}

	err := pgxscan.Get(ctx, p.Pool, user, query, strings.TrimSpace(username))
	if err != nil {
		return nil, err
	}

	return user, nil
}
