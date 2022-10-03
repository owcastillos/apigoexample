package persistence

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/owcastillos/apigoexample/models"
	"github.com/owcastillos/apigoexample/utils"
)

func GetUsers() []*models.User {
	log.Println("GetUsers init")
	db := createDBConnection()
	if db == nil {
		return nil
	}
	defer db.Close()
	result, err := db.Query("SELECT * FROM `USER`")
	if err != nil {
		log.Panic(err)
		return nil
	}
	users := make([]*models.User, 0)
	for result.Next() {
		var user models.User
		if err := result.Scan(&user.ID, &user.Email, &user.Nombre, &user.Apellido, &user.Fecha); err != nil {
			log.Println(err)
			return nil
		}
		users = append(users, &user)
	}
	log.Println("GetUsers correct!")
	return users
}

func GetUserById(idUser string) *models.User {
	log.Println("GetUserById init")
	db := createDBConnection()
	if db == nil {
		return nil
	}
	defer db.Close()
	result := db.QueryRow("SELECT * FROM `USER` WHERE ID = ?", idUser)
	var user models.User
	if err := result.Scan(&user.ID, &user.Email, &user.Nombre, &user.Apellido, &user.Fecha); err != nil {
		log.Println(err)
		return nil
	}
	log.Println("GetUserById correct!")
	return &user
}

func InsertUser(user *models.User) *models.Result {
	log.Println("InsertUser init")
	if user == nil {
		return utils.HandleString("OBJ NULL")
	}
	db := createDBConnection()
	defer db.Close()
	insertString := fmt.Sprintf("INSERT INTO `USER` (EMAIL, NOMBRE, APELLIDO, FECHA) VALUES ('%s','%s','%s','%s')",
		user.Email, user.Nombre, user.Apellido, user.Fecha)
	result, err := db.Exec(insertString)
	if err != nil {
		return utils.HandleError(err)
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return utils.HandleError(err)
	}
	log.Println("InsertUser correct!")
	return &models.Result{
		Status: fmt.Sprintf("ID: %d", lastId),
	}
}

func UpdateUser(user *models.User) *models.Result {
	log.Println("UpdateUser init")
	if user == nil {
		return utils.HandleString("OBJ NULL")
	}
	db := createDBConnection()
	defer db.Close()
	updateString := fmt.Sprintf("UPDATE `USER` SET EMAIL = '%s', NOMBRE = '%s', APELLIDO = '%s', FECHA = '%s' WHERE ID = %d",
		user.Email, user.Nombre, user.Apellido, user.Fecha, user.ID)
	result, err := db.Exec(updateString)
	if err != nil {
		return utils.HandleError(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return utils.HandleError(err)
	}
	var status string
	if rowsAffected > 0 {
		status = "OK"
	} else {
		status = "No row has been affected"
	}
	log.Println("UpdateUser correct!")
	return &models.Result{
		Status: status,
	}
}

func DeleteUsers() *models.Result {
	log.Println("DeleteUsers init")
	db := createDBConnection()
	if db == nil {
		return nil
	}
	defer db.Close()
	result, err := db.Exec("DELETE FROM `USER`")
	if err != nil {
		return utils.HandleError(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return utils.HandleError(err)
	}
	var status string
	if rowsAffected > 0 {
		status = "OK"
	} else {
		status = "No row has been affected"
	}
	log.Println("DeleteUsers correct!")
	return &models.Result{
		Status: status,
	}
}

func DeleteUserById(idUser string) *models.Result {
	log.Println("DeleteUserById init")
	db := createDBConnection()
	if db == nil {
		return nil
	}
	defer db.Close()
	result, err := db.Exec("DELETE FROM `USER` WHERE ID = ?", idUser)
	if err != nil {
		return utils.HandleError(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return utils.HandleError(err)
	}
	var status string
	if rowsAffected > 0 {
		status = "OK"
	} else {
		status = "No row has been affected"
	}
	log.Println("DeleteUserById correct!")
	return &models.Result{
		Status: status,
	}
}

func createDBConnection() *sql.DB {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", "local", "local", "localhost", "3316", "api-test")
	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Panic(err)
		return nil
	}
	err = db.Ping()
	if err != nil {
		log.Panic(err)
		defer db.Close()
		return nil
	}
	return db
}
