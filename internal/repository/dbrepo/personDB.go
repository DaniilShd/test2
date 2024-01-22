package dbrepo

import (
	"context"
	"fmt"
	"time"

	"github.com/DaniilShd/test2/models"
)

func (m *postgresDBRepo) GetAllPersons() (*[]models.Person, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	querySelectAll := `
	select  id, name, surname, patronymic, gender, nationality, age
	from persons
	`

	rows, err := m.DB.QueryContext(ctx, querySelectAll)
	if err != nil {
		return nil, err
	}

	var persons []models.Person

	for rows.Next() {
		var person models.Person
		err = rows.Scan(&person.ID,
			&person.Name,
			&person.Surname,
			&person.Patronymic,
			&person.Gender,
			&person.Nationality,
			&person.Age)

		if err != nil {
			return nil, err
		}

		persons = append(persons, person)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &persons, nil
}

func (m *postgresDBRepo) GetPersonByID(id int) (*models.Person, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	querySelectByID := `
	select  id, name, surname, patronymic, gender, nationality, age
	from persons
	where id = $1
	`

	var person models.Person

	row := m.DB.QueryRowContext(ctx, querySelectByID, id)

	err := row.Scan(&person.ID,
		&person.Name,
		&person.Surname,
		&person.Patronymic,
		&person.Gender,
		&person.Nationality,
		&person.Age)

	if err != nil {
		return nil, err
	}

	if err = row.Err(); err != nil {
		return nil, err
	}

	return &person, nil
}

func (m *postgresDBRepo) DeletePersonByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	deleteItem := `
	delete
	from persons
	where id = $1
	`
	_, err := m.DB.ExecContext(ctx, deleteItem, id)
	if err != nil {
		return err
	}

	return nil
}

func (m *postgresDBRepo) UpdatePersonByID(person *models.Person) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	queryUpdate := `update persons 
	set name=$1, surname=$2, patronymic=$3, gender=$4, nationality=$5, age=$6
	where id = $7
	`

	_, err := m.DB.ExecContext(ctx, queryUpdate,
		&person.Name,
		&person.Surname,
		&person.Patronymic,
		&person.Gender,
		&person.Nationality,
		&person.Age,
		&person.ID)

	if err != nil {
		return err
	}
	return nil
}

func (m *postgresDBRepo) InsertPerson(person *models.Person) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	queryInsert := `
	insert into persons 
	(name, surname, patronymic, gender, nationality, age)
	VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := m.DB.ExecContext(ctx, queryInsert,
		&person.Name,
		&person.Surname,
		&person.Patronymic,
		&person.Gender,
		&person.Nationality,
		&person.Age,
	)
	if err != nil {
		return err
	}

	return nil
}

func (m *postgresDBRepo) GetPersonsByOffset(filter *models.Filter) (*[]models.Person, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	var persons []models.Person

	querySelectByFilter := `
	select  id, name, surname, patronymic, gender, nationality, age
	from persons
	`
	if filter.Set != 0 {
		querySelectByFilter = querySelectByFilter + `WHERE `
	}

	if filter.Name != "" {
		querySelectByFilter = querySelectByFilter + `name=` + `'` + filter.Name + `'`
		if filter.Set != 1 {
			querySelectByFilter = querySelectByFilter + ` AND `
		}
		filter.Set -= 1
	}

	if filter.Surname != "" {
		querySelectByFilter = querySelectByFilter + `surname=` + `'` + filter.Surname + `'`
		if filter.Set != 1 {
			querySelectByFilter = querySelectByFilter + ` AND `
		}
		filter.Set -= 1
	}

	if filter.Patronymic != "" {
		querySelectByFilter = querySelectByFilter + `patronymic=` + `'` + filter.Patronymic + `'`
		if filter.Set != 1 {
			querySelectByFilter = querySelectByFilter + ` AND `
		}
		filter.Set -= 1
	}

	if filter.Nationality != "" {
		querySelectByFilter = querySelectByFilter + `nationality=` + `'` + filter.Nationality + `'`
		if filter.Set != 1 {
			querySelectByFilter = querySelectByFilter + ` AND `
		}
		filter.Set -= 1
	}

	if filter.Gender != "" {
		querySelectByFilter = querySelectByFilter + `gender=` + `'` + filter.Gender + `'`
		if filter.Set != 1 {
			querySelectByFilter = querySelectByFilter + ` AND `
		}
		filter.Set -= 1
	}

	if filter.Age != "" {
		querySelectByFilter = querySelectByFilter + `age=` + filter.Age
		if filter.Set != 1 {
			querySelectByFilter = querySelectByFilter + ` AND `
		}
		filter.Set -= 1
	}

	//querySelectByFilter = querySelectByFilter + ` OFFSET=` + filter.Offset + ` LIMIT=` + filter.Limit

	fmt.Println(querySelectByFilter)

	rows, err := m.DB.QueryContext(ctx, querySelectByFilter)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var person models.Person
		err = rows.Scan(&person.ID,
			&person.Name,
			&person.Surname,
			&person.Patronymic,
			&person.Gender,
			&person.Nationality,
			&person.Age)

		if err != nil {
			return nil, err
		}
		persons = append(persons, person)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &persons, nil
}
