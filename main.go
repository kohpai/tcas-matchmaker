package main

import (
	"log"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/kohpai/tcas-3rd-round-resolver/mapper"
	"github.com/kohpai/tcas-3rd-round-resolver/model"
	"github.com/kohpai/tcas-3rd-round-resolver/util"
)

func main() {
	students, err := util.ReadStudents()
	if err != nil {
		log.Fatalln(err)
	}
	courses, err := util.ReadCourses()
	if err != nil {
		log.Fatalln(err)
	}
	rankings, err := util.ReadRankings()
	if err != nil {
		log.Fatalln(err)
	}

	clearingHouse := model.NewClearingHouse(
		util.GetPendingStudents(
			mapper.CreateStudentMap(
				students,
				mapper.CreateCourseMap(
					courses,
					rankings,
				),
			),
		),
	)

	clearingHouse.Execute()

	allStudents, acceptedStudents, rejectedStudents := clearingHouse.Students(), clearingHouse.AcceptedStudents(), clearingHouse.RejectedStudents()
	if len(allStudents) != len(acceptedStudents)+len(rejectedStudents) {
		log.Fatal("some students are missing")
	}

	outputs := mapper.ToOutput(allStudents)
	outputFile, err := os.Create("output.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer outputFile.Close()

	err = gocsv.MarshalFile(&outputs, outputFile)
	if err != nil {
		log.Fatal(err)
	}
}
