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
		return nil, err
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
		return nil, err
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
func (p *Postgres) GetUserByID(userID int64) (*types.User, error) {

	user := types.User{ID: userID}
	err := p.db.QueryRow("SELECT name, birthday, is_male FROM tb_user5 WHERE id = $1", userID).Scan(
		&user.Name,
		&user.Birthday,
		&user.IsMale,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

//Создать пользователя и получить его ID
func (p *Postgres) MakeUser(newUser *types.User) (int64, error) {

	var userID int64
	err := p.db.QueryRow("INSERT INTO tb_user5 (name, birthday, age, is_male) VALUES ($1, $2, $3, $4) RETURNING id",
		newUser.Name, newUser.Birthday, newUser.Age, newUser.IsMale).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

//Изменить пользователя
func (p *Postgres) PutUserByID(newUser *types.User) error {

	_, err := p.db.Exec("UPDATE tb_user5 SET name = $1, birthday = $2, age = $3, is_male = $4 WHERE id = $5",
		newUser.Name, newUser.Birthday, newUser.Age, newUser.IsMale, newUser.ID)
	if err != nil {
		return err
	}

	return nil
}

//удалить пользователя
func (p *Postgres) DelUserByID(userID int64) error {

	_, err := p.db.Exec("DELETE FROM tb_user5 WHERE id = $1", userID)
	if err != nil {
		return err
	}

	return nil
}
