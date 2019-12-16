package users

import (
	"fmt"
	"github.com/johnwoz123/pharmacy-user-api/persistence/users_gateway"
	"github.com/johnwoz123/pharmacy-user-api/utils/errors"
	"strings"
)

const (
	NoRowsInResultSetError = "no rows in result set"
	insertUserQuery        = "INSERT INTO users(first_name, last_name, email,status , password, date_created) VALUES(?, ?, ?, ?, ?, ?);"
	getUserByIdQuery       = "SELECT id, first_name, last_name, email, date_created, status FROM users where id=?"
	updateUserQuery        = "update users set first_name =?, last_name=?, email=? where id = ?;"
	deleteUserQuery        = "DELETE FROM users WHERE id=?;"
	getUserByStatusQuery   = "SELECT id, first_name, last_name, email, status, date_created FROM users WHERE status=?;"
)

func (user *User) Get() *errors.RestErrors {
	stmt, err := users_gateway.Client.Prepare(getUserByIdQuery)
	if err != nil {
		return errors.NotFoundError(fmt.Sprintf("user with the id %d was not found in the database.", user.Id))
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		if strings.Contains(err.Error(), NoRowsInResultSetError) {
			return errors.NotFoundError(fmt.Sprintf("user %d not found", user.Id))
		}
		return errors.InternalServerError(fmt.Sprintf("error while trying to get user %d:%s", user.Id, err.Error()))
	}
	return nil
}

func (user *User) Persist() *errors.RestErrors {

	stmtInsert, err := users_gateway.Client.Prepare(insertUserQuery)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	// since there are no errors defer close until complete
	defer stmtInsert.Close()
	result, err := stmtInsert.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.Password, user.DateCreated)
	if err != nil {
		return errors.InternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	insertedId, err := result.LastInsertId()
	if err != nil {
		return errors.InternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	user.Id = insertedId
	return nil
}

func (user *User) Update() *errors.RestErrors {
	stmtUpdate, err := users_gateway.Client.Prepare(updateUserQuery)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	// since there are no errors defer close until complete
	defer stmtUpdate.Close()
	_, err = stmtUpdate.Exec(user.FirstName, user.LastName, user.Email, user.Id)

	if err != nil {
		return errors.InternalServerError(err.Error())
	}

	return nil
}

func (user *User) Delete() *errors.RestErrors {
	stmtDelete, err := users_gateway.Client.Prepare(deleteUserQuery)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmtDelete.Close()
	if _, err = stmtDelete.Exec(user.Id); err != nil {
		return errors.InternalServerError(err.Error())
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErrors) {
	stmtFindByStatus, err := users_gateway.Client.Prepare(getUserByStatusQuery)
	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}
	defer stmtFindByStatus.Close()
	rows, err := stmtFindByStatus.Query(status)
	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}
	defer rows.Close()

	result := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, errors.NotFoundError(err.Error())
		}
		result = append(result, user)
	}

	if len(result) == 0 {
		return nil, errors.NotFoundError("no users found with that criteria")
	}
	return result, nil
}
