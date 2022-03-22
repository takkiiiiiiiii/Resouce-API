package repository

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/takkiiiiiiiii/rest-api/model/entity"

	"github.com/takkiiiiiiiii/rest-api/model/database"
)

type ApiRepository interface {
	Get_Api() (data []entity.ApiEntity, err error)
	Insert_Api(datum entity.ApiEntity) (id int64, err error)
	Update_Api(datum entity.ApiEntity) (err error)
	Delete_Api(id int) (err error)
}

type apiRepository struct {
}

func NewApiRepository() *apiRepository{
	return &apiRepository{}
}

var Db = database.Db

func (tr *apiRepository) Get_Api() (data []entity.ApiEntity, err error) {
	data = []entity.ApiEntity{}
	rows, err := Db.Query("SELECT * FROM User_Info")
	if err != nil {
         log.Print(err)
		 return
	}
	for rows.Next(){
		datum := entity.ApiEntity{}
		err = rows.Scan(&datum.Id, &datum.Name, &datum.Contents, &datum.Created)
		if err != nil {
			log.Print(err)
			return
		}
		data = append(data, datum)
	}
	return
}

func (tr *apiRepository) Insert_Api(datum entity.ApiEntity) (id int64, err error) {
    stmt, err := Db.Prepare("INSERT INTO User_Info(NAME, CONTENTS, CREATED) VALUES(?,?,now())")
	if err != nil {
		log.Print(err)
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(datum.Name, datum.Contents)
	if err != nil {
		log.Print(err)
		return
	}
	id, err = result.LastInsertId()
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (tr *apiRepository) Update_Api(datum entity.ApiEntity) (err error) {
	stmt1, err := Db.Prepare("UPDATE User_Info SET NAME = ?, CONTENTS = ?, CREATED = now() WHERE ID = ?") 
	if err != nil {
		log.Print(err)
		return
	}
	defer stmt1.Close()
	if err != nil {
		log.Print(err)
		return
	}
	_, err = stmt1.Exec(datum.Name, datum.Contents, datum.Id)
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (tr *apiRepository) Delete_Api(id int)(err error) {
	stmt, err := Db.Prepare("DELETE FROM User_Info WHERE ID = ?")
	if err != nil {
		log.Print(err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		log.Print(err)
		return
	}
	return
}

