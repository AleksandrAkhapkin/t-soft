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
		return nil, errors.Wrap(err, "in pg GetAllUsers with Query")
	}
	defer rows.Close()
	user := types.User{}
	for rows.Next() {
		if err = rows.Scan(&user.ID, &user.Name, &user.Birthday, &user.IsMale); err != nil {
			return nil, errors.Wrap(err, "in pg GetAllUsers with Scan")
		}
		users = append(users, user)
	}

	return users, nil
}

//Получить пользователей с датой рождения не больше maxBDay (включительно)
func (p *Postgres) GetUserWithMaxBDay(maxBDay string) ([]types.User, error) {

	users := make([]types.User, 0)
	rows, err := p.db.Query("SELECT id, name, birthday, is_male FROM tb_user5 WHERE birthday <= $1", maxBDay)
	if err != nil {
		return nil, errors.Wrap(err, "in pg GetAllUsers with Query")
	}
	defer rows.Close()
	user := types.User{}
	for rows.Next() {
		if err = rows.Scan(&user.ID, &user.Name, &user.Birthday, &user.IsMale); err != nil {
			return nil, errors.Wrap(err, "in pg GetAllUsers with Scan")
		}
		users = append(users, user)
	}

	return users, nil
}

//Получить пользователя по ID
func (p *Postgres) GetUserByID(userID int) (*types.User, error) {

	user := types.User{ID: userID}
	err := p.db.QueryRow("SELECT name, birthday, is_male FROM tb_user5 WHERE id = $1", userID).Scan(
		&user.Name,
		&user.Birthday,
		&user.IsMale,
	)
	if err != nil {
		return nil, errors.Wrap(err, "in pg GetAllUsers with QueryRow")
	}

	return &user, nil
}
