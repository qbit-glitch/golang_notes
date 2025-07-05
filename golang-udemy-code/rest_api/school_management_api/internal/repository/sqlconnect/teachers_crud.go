package sqlconnect

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"school_management_api/internal/models"
	"school_management_api/pkg/utils"
	"strconv"
)

func GetTeachersDbHandler(teachers []models.Teacher, r *http.Request) ([]models.Teacher, error) {
	db, err := ConnectDb()
	if err != nil {
		return nil, utils.ErrorHandler(err, "error retrieving data")
	}
	defer db.Close()

	query := "SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE 1=1"
	var args []interface{}

	query, args = utils.AddFilters(r, query, args)

	query = utils.AddSorting(r, query)

	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println("err")

		return nil, utils.ErrorHandler(err, "error retrieving data")
	}
	defer rows.Close()

	for rows.Next() {
		teacher := models.Teacher{}
		err = rows.Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)
		if err != nil {

			return nil, utils.ErrorHandler(err, "error retrieving data")
		}
		teachers = append(teachers, teacher)
	}
	return teachers, nil
}

func GetTeacherByID(id int) (models.Teacher, error) {
	db, err := ConnectDb()
	if err != nil {

		return models.Teacher{}, utils.ErrorHandler(err, "error retrieving data ")
	}
	defer db.Close()

	var teacher models.Teacher
	err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", id).Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)
	if err == sql.ErrNoRows {

		return models.Teacher{}, utils.ErrorHandler(err, "error retrieving data ")
	} else if err != nil {
		fmt.Println(err)

		return models.Teacher{}, utils.ErrorHandler(err, "error retrieving data ")
	}
	return teacher, nil
}

func AddTeachersDBHandler(newTeachers []models.Teacher) ([]models.Teacher, error) {

	fmt.Println("------ AddTeachersDBHandler Called -------")

	db, err := ConnectDb()
	if err != nil {

		return nil, utils.ErrorHandler(err, "error adding data")
	}
	defer db.Close()
	stmt, err := db.Prepare(utils.GenerateInsertQuery("teachers", models.Teacher{}))
	if err != nil {

		return nil, utils.ErrorHandler(err, "error adding data")
	}
	defer stmt.Close()
	fmt.Printf("Teachers Add Handler")

	addedTeachers := make([]models.Teacher, len(newTeachers))

	for i, newTeacher := range newTeachers {

		values := utils.GetStructValues(newTeacher)
		fmt.Println(newTeacher)

		fmt.Println("VALUES:", values)
		res, err := stmt.Exec(values...)

		if err != nil {

			return nil, utils.ErrorHandler(err, "error adding data")
		}
		lastID, err := res.LastInsertId()
		if err != nil {

			return nil, utils.ErrorHandler(err, "error adding data")
		}
		newTeacher.ID = int(lastID)
		addedTeachers[i] = newTeacher
	}
	return addedTeachers, nil
}

func UpdateTeacher(id int, updatedTeacher models.Teacher) (models.Teacher, error) {
	db, err := ConnectDb()
	if err != nil {

		return models.Teacher{}, utils.ErrorHandler(err, "error updating data")
	}
	defer db.Close()

	var existingTeacher models.Teacher
	err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", id).Scan(&existingTeacher.ID, &existingTeacher.FirstName, &existingTeacher.LastName, &existingTeacher.Email, &existingTeacher.Class, &existingTeacher.Subject)
	if err != nil {
		if err != sql.ErrNoRows {

			return models.Teacher{}, utils.ErrorHandler(err, "error updating data")
		}

		return models.Teacher{}, utils.ErrorHandler(err, "error updating data")
	}

	updatedTeacher.ID = existingTeacher.ID
	_, err = db.Exec("UPDATE teachers SET first_name = ?, last_name = ?, email = ?, class = ?, subject = ? WHERE id = ?", updatedTeacher.FirstName, updatedTeacher.LastName, updatedTeacher.Email, updatedTeacher.Class, updatedTeacher.Subject, updatedTeacher.ID)
	if err != nil {

		return models.Teacher{}, utils.ErrorHandler(err, "error updating data")
	}
	return updatedTeacher, nil
}

func PatchTeachers(updates []map[string]interface{}) error {
	db, err := ConnectDb()
	if err != nil {

		return utils.ErrorHandler(err, "error updating data")
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {

		return utils.ErrorHandler(err, "error updating data")
	}

	for _, update := range updates {
		idStr, ok := update["id"].(string)
		if !ok {
			tx.Rollback()

			return utils.ErrorHandler(err, "Invalid id")
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			tx.Rollback()
			return utils.ErrorHandler(err, "invalid id")
		}

		var teacherFromDb models.Teacher
		err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", id).Scan(&teacherFromDb.ID, &teacherFromDb.FirstName, &teacherFromDb.LastName, &teacherFromDb.Email, &teacherFromDb.Class, &teacherFromDb.Subject)

		if err != nil {
			tx.Rollback()
			if err == sql.ErrNoRows {

				return utils.ErrorHandler(err, "Teacher Not Found")
			}

			return utils.ErrorHandler(err, "error updating data")
		}

		teacherVal := reflect.ValueOf(&teacherFromDb).Elem()
		teacherType := teacherVal.Type()

		for k, v := range update {
			if k == "id" {
				continue
			}
			for i := 0; i < teacherVal.NumField(); i++ {
				field := teacherType.Field(i)
				if field.Tag.Get("json") == k+",omitempty" {
					fieldVal := teacherVal.Field(i)
					if fieldVal.CanSet() {
						val := reflect.ValueOf(v)
						if val.Type().ConvertibleTo(fieldVal.Type()) {
							fieldVal.Set(val.Convert(fieldVal.Type()))
						} else {
							tx.Rollback()
							log.Printf("cannot convert %v to %v", val.Type(), fieldVal.Type())
							return utils.ErrorHandler(err, "error updating data")
						}
					}
					break
				}
			}
		}
		_, err = tx.Exec("UPDATE teachers SET first_name = ?, last_name = ?, email = ?, class = ?, subject = ? WHERE id = ?", teacherFromDb.FirstName, teacherFromDb.LastName, teacherFromDb.Email, teacherFromDb.Class, teacherFromDb.Subject, teacherFromDb.ID)
		if err != nil {
			tx.Rollback()

			return utils.ErrorHandler(err, "error updating data")
		}
	}

	err = tx.Commit()
	if err != nil {

		return utils.ErrorHandler(err, "error updating data")
	}
	return nil
}

func PatchOneTeacher(id int, updates map[string]interface{}) (models.Teacher, error) {

	db, err := ConnectDb()
	if err != nil {
		log.Println(err)

		return models.Teacher{}, utils.ErrorHandler(err, "error updating data")
	}
	defer db.Close()

	var existingTeacher models.Teacher
	err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", id).Scan(&existingTeacher.ID, &existingTeacher.FirstName, &existingTeacher.LastName, &existingTeacher.Email,&existingTeacher.Class, &existingTeacher.Subject)
	if err != nil {
		if err != sql.ErrNoRows {

			return models.Teacher{}, utils.ErrorHandler(err, "Teacher not Found")
		}

		return models.Teacher{}, utils.ErrorHandler(err, "error updating data")
	}

	teacherVal := reflect.ValueOf(&existingTeacher).Elem()
	teacherType := teacherVal.Type()
	for k, v := range updates {

		for i := 0; i < teacherVal.NumField(); i++ {

			field := teacherType.Field(i)

			if field.Tag.Get("json") == k+",omitempty" {
				if teacherVal.Field(i).CanSet() {

					teacherVal.Field(i).Set(reflect.ValueOf(v).Convert(teacherVal.Field(i).Type()))
				}
			}
		}
	}

	_, err = db.Exec("UPDATE teachers SET first_name = ?, last_name = ?, email = ?, class = ?, subject = ? WHERE id = ?", existingTeacher.FirstName, existingTeacher.LastName, existingTeacher.Email, existingTeacher.Class, existingTeacher.Subject, existingTeacher.ID)
	if err != nil {

		return models.Teacher{}, utils.ErrorHandler(err, "error updating data")
	}
	return existingTeacher, nil
}

func DeleteOneTeacher(id int) error {
	db, err := ConnectDb()
	if err != nil {
		log.Println(err)

		return utils.ErrorHandler(err, "error deleting data")
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM teachers WHERE id = ?", id)
	if err != nil {

		return utils.ErrorHandler(err, "error deleting data")
	}

	fmt.Println(result.RowsAffected())

	rowsAffected, err := result.RowsAffected()
	if err != nil {

		return utils.ErrorHandler(err, "error deleting data")
	}

	if rowsAffected == 0 {

		return utils.ErrorHandler(err, "teacher not found")
	}
	return nil
}

func DeleteTeachers(ids []int) ([]int, error) {
	db, err := ConnectDb()
	if err != nil {

		return nil, utils.ErrorHandler(err, "error deleting data")
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {

		return nil, utils.ErrorHandler(err, "error deleting data")
	}

	stmt, err := tx.Prepare("DELETE FROM teachers WHERE id = ?")
	if err != nil {

		tx.Rollback()

		return nil, utils.ErrorHandler(err, "error deleting data")
	}
	defer stmt.Close()

	deleteIds := []int{}

	for _, id := range ids {
		result, err := stmt.Exec(id)
		if err != nil {
			tx.Rollback()

			return nil, utils.ErrorHandler(err, "error deleting data")
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			tx.Rollback()

			return nil, utils.ErrorHandler(err, "error deleting data")
		}

		if rowsAffected > 0 {
			deleteIds = append(deleteIds, id)
		}

		if rowsAffected < 1 {
			tx.Rollback()

			return nil, utils.ErrorHandler(err, fmt.Sprintf("ID %d does not exist", id))
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)

		return nil, utils.ErrorHandler(err, "error deleting data")
	}

	if len(deleteIds) < 1 {

		return nil, utils.ErrorHandler(err, "IDs do not exist")
	}
	return deleteIds, nil
}



func GetStudentsByTeacherIdFromDb(teacherId string, students []models.Student) ([]models.Student, error) {
	db, err := ConnectDb()
	if err != nil {
		return nil, utils.ErrorHandler(err, "error retrieving data")
	}
	defer db.Close()

	query := `SELECT id, first_name, last_name, email, class FROM students WHERE class = (SELECT class FROM teachers WHERE id = ?)`
	rows, err := db.Query(query, teacherId)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var student models.Student
		err := rows.Scan(&student.ID, &student.FirstName, &student.LastName, &student.Email, &student.Class)
		if err != nil {
			return nil, utils.ErrorHandler(err, "error retrieving data")
		}
		students = append(students, student)
	}

	err = rows.Err()
	if err != nil {
		return nil, utils.ErrorHandler(err, "error retrieving data")
	}
	return students, nil
}


func GetStudentCountByTeacherIdFromDb(teacherId string) (int, error) {
	db, err := ConnectDb()
	if err != nil {
		return 0, utils.ErrorHandler(err, "error retrieving data")
	}
	defer db.Close()

	query := `SELECT COUNT(*) FROM students WHERE class = (SELECT class FROM teachers WHERE id = ?)`
	var studentCount int
	err = db.QueryRow(query, teacherId).Scan(&studentCount)
	if err != nil {
		return 0, utils.ErrorHandler(err, "error retrieving data")
	}
	return studentCount, nil
}

