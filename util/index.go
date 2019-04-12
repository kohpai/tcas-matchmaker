package util

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/kohpai/tcas-3rd-round-resolver/mapper"
	"github.com/kohpai/tcas-3rd-round-resolver/model"
)

func ReadStudents(filename string) ([]mapper.Student, error) {
	var students []mapper.Student
	err := readJsonFile(filename, &students)
	if err != nil {
		return nil, err
	}
	return students, nil
}

func ReadCourses(filename string) ([]mapper.Course, error) {
	var courses []mapper.Course
	err := readJsonFile(filename, &courses)
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func ReadRankings(filename string) ([]mapper.Ranking, error) {
	var rankings []mapper.Ranking
	err := readCsvFile(filename, &rankings)
	if err != nil {
		return nil, err
	}
	return rankings, nil
}

func readJsonFile(filename string, data interface{}) error {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return errors.New("cannot read JSON file: " + err.Error())
	}

	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return errors.New("cannot read bytes: " + err.Error())
	}

	if err = json.Unmarshal(bytes, data); err != nil {
		return errors.New("cannot unmarshal: " + err.Error())
	}

	return nil
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
