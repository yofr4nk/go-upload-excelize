package files

import (
	"errors"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"go-upload-excelize/database"
	"go-upload-excelize/models"
	"mime/multipart"
)

// ProcessPeople get a file to return []People
func ProcessPeople(file multipart.File) ([]models.Person, error) {
	var people []models.Person
	excFile, excErr := excelize.OpenReader(file)

	if excErr != nil {
		return nil, errors.New("Something went wrong: " + excErr.Error())
	}

	rows, rowErr := excFile.Rows(excFile.GetSheetName(0))

	if rowErr != nil {
		return nil, errors.New("Something went wrong: " + rowErr.Error())
	}

	for rows.Next() {
		row, colErr := rows.Columns()
		if colErr != nil {
			return nil, errors.New("Something went wrong: " + colErr.Error())
		}

		p := models.Person{
			Name:     row[1],
			LastName: row[0],
			Role:     row[2],
		}

		database.DBConnection.Create(&p)

		people = append(people, p)
	}

	return people, nil
}
