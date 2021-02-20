package postgres

import (
	"github.com/AleksandrAkhapkin/testTNS/task4/internal/task4/types"
	"github.com/pkg/errors"
)

//возвращает список всех пользователей (возраст не возращается)
func (p *Postgres) GetAllUsers() ([]types.User, error) {

	users := make([]types.User, 0)
	rows, err := p.db.Query("SELECT id, name, birthday, is_male FROM tb_user5 ")
	if err != nil {
		return nil, errors.Wrap(err, "err with Query")
	}
	defer rows.Close()
	user := types.User{}
	for rows.Next() {
		if err = rows.Scan(&user.ID, &user.Name, &user.Birthday, &user.IsMale); err != nil {
			return nil, errors.Wrap(err, "err with Scan")
		}
		users = append(users, user)
	}

	return users, nil
}
