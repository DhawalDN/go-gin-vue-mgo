package models

import (
	"go-gin-mgod/db"
	"go-gin-mgod/forms"

	"gopkg.in/mgo.v2/bson"
)

type Student struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"name"`
	Email    string        `json:"email"`
	Password string        `json:"password"`
	Gender   string        `json:"gender"`
}

type StudentModel struct{}

var server = "127.0.0.1"

var dbConnect = db.NewConnection(server)

func (m *StudentModel) Create(data forms.CreateStudentCommand) error {
	collection := dbConnect.Use("db-mgo", "students")
	err := collection.Insert(bson.M{"name": data.Name, "email": data.Email, "password": data.Password, "gender": data.Gender})
	return err
}

func (m *StudentModel) Find() (list []Student, err error) {
	collection := dbConnect.Use("db-mgo", "students")
	err = collection.Find(bson.M{}).All(&list)
	return list, err
}

func (m *StudentModel) Get(id string) (student Student, err error) {
	collection := dbConnect.Use("db-mgo", "students")
	err = collection.FindId(bson.ObjectIdHex(id)).One(&student)
	return student, err
}

func (m *StudentModel) Update(id string, data forms.UpdateStudentCommand) (err error) {
	collection := dbConnect.Use("db-mgo", "students")
	err = collection.UpdateId(bson.ObjectIdHex(id), data)

	return err
}

func (m *StudentModel) Delete(id string) (err error) {
	collection := dbConnect.Use("db-mgo", "students")
	err = collection.RemoveId(bson.ObjectIdHex(id))

	return err
}
