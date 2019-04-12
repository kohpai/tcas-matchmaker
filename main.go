package main

import (
	"log"

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

	allStudents := clearingHouse.Students()

	// @ASSERTION, this shouldn't happen!
	if len(allStudents) != len(clearingHouse.AcceptedStudents())+len(clearingHouse.RejectedStudents()) {
		log.Fatal("some students are missing")
	}

	outputs := mapper.ToOutput(allStudents)
	if err := util.WriteCsvFile("output.csv", &outputs); err != nil {
		log.Fatal(err)
	}
}
