package main

import (
	"errors"
	"log"
	"os"

	"github.com/kohpai/tcas-3rd-round-resolver/mapper"
	"github.com/kohpai/tcas-3rd-round-resolver/model"
	"github.com/kohpai/tcas-3rd-round-resolver/util"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "resolver"
	app.Usage = "resolve for admitted students"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "students",
			Value: "students.json",
			Usage: "a file to input as student-course enrollment",
		},
		cli.StringFlag{
			Name:  "courses",
			Value: "courses.json",
			Usage: "a file to input all courses in the system",
		},
		cli.StringFlag{
			Name:  "rankings",
			Value: "ranking.csv",
			Usage: "a file to input as ranking for students in each course",
		},
		cli.StringFlag{
			Name:  "output",
			Value: "output.csv",
			Usage: "a file to be saved as a result output",
		},
	}
	app.Action = action

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func action(c *cli.Context) error {
	students, err := util.ReadStudents(c.String("students"))
	if err != nil {
		return err
	}
	courses, err := util.ReadCourses(c.String("courses"))
	if err != nil {
		return err
	}
	rankings, err := util.ReadRankings(c.String("rankings"))
	if err != nil {
		return err
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
		return errors.New("some students are missing")
	}

	outputs := mapper.ToOutput(allStudents)
	if err := util.WriteCsvFile("output.csv", &outputs); err != nil {
		return err
	}

	return nil
}
