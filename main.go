package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/gocarina/gocsv"

	"github.com/kohpai/tcas-3rd-round-resolver/mapper"
)

func main() {
	readStudents()
	readCourses()
	readRankings()
}

func readStudents() {
	jsonFile, err := os.Open("data/TC01/con1_student_enroll.json")
	if err != nil {
		log.Println("cannot read JSON file", err)
		return
	}

	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println("cannot read bytes", err)
		return
	}

	var students []mapper.Student

	if err = json.Unmarshal(bytes, &students); err != nil {
		log.Println("cannot unmarshal", err)
		return
	}

	log.Println("length", len(students))
	log.Println(students)
}

func readCourses() {
	jsonFile, err := os.Open("data/TC01/all_course.json")
	if err != nil {
		log.Println("cannot read JSON file", err)
		return
	}

	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println("cannot read bytes", err)
		return
	}

	var courses []mapper.Course

	if err = json.Unmarshal(bytes, &courses); err != nil {
		log.Println("cannot unmarshal", err)
		return
	}

	log.Println("length", len(courses))
	log.Println(courses)
}

func readRankings() {
	csvFile, err := os.Open("data/TC01/con1_course_accept.csv")
	if err != nil {
		log.Println("cannot read CSV file", err)
		return
	}

	defer csvFile.Close()

	var rankings []mapper.Ranking
	if err := gocsv.UnmarshalFile(csvFile, &rankings); err != nil {
		log.Println("cannot unmarshal", err)
		return
	}

	log.Println("length", len(rankings))
	log.Println(rankings)
}
