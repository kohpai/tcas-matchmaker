package util

import (
	"errors"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/kohpai/tcas-matchmaker/mapper"
	"github.com/kohpai/tcas-matchmaker/model"
)

func ReadApps(filename string) ([]mapper.Application, error) {
	var apps []mapper.Application
	err := readCsvFile(filename, &apps)
	if err != nil {
		return nil, err
	}
	return apps, nil
}

func ReadCourses(filename string) ([]mapper.Course, error) {
	var courses []mapper.Course
	err := readCsvFile(filename, &courses)
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func readCsvFile(filename string, data interface{}) error {
	csvFile, err := os.Open(filename)
	if err != nil {
		return errors.New("cannot read CSV file: " + err.Error())
	}

	defer csvFile.Close()

	if err := gocsv.UnmarshalFile(csvFile, data); err != nil {
		return errors.New("cannot unmarshal: " + err.Error())
	}

	return nil
}

func WriteCsvFile(filename string, data interface{}) error {
	outputFile, err := os.Create(filename)
	if err != nil {
		return errors.New("cannot create file: " + err.Error())
	}

	defer outputFile.Close()

	if err := gocsv.MarshalFile(data, outputFile); err != nil {
		return errors.New("cannot write file: " + err.Error())
	}

	return nil
}

func GetPendingStudents(studentMap mapper.StudentMap) []*model.Student {
	pendingStudents := make([]*model.Student, len(studentMap))

	i := 0
	for _, student := range studentMap {
		pendingStudents[i] = student
		i++
	}

	return pendingStudents
}
