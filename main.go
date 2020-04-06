package main

import (
	"log"
	"os"

	"github.com/kohpai/tcas-matchmaker/mapper"
	"github.com/kohpai/tcas-matchmaker/model"
	"github.com/kohpai/tcas-matchmaker/util"
	"github.com/pkg/errors"
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
		return errors.Wrap(err, "students")
	}
	courses, err := util.ReadCourses(c.String("courses"))
	if err != nil {
		return errors.Wrap(err, "courses")
	}
	rankings, err := util.ReadRankings(c.String("rankings"))
	if err != nil {
		return errors.Wrap(err, "rankings")
	}

	rankingInfoMap := mapper.ExtractRankings(rankings)
	clearingHouse := model.NewClearingHouse(
		util.GetPendingStudents(
			mapper.CreateStudentMap(
				students,
				mapper.CreateCourseMap(
					courses,
					rankingInfoMap,
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

	outputs := mapper.ToOutput(allStudents, rankingInfoMap)
	if err := util.WriteCsvFile(c.String("output"), &outputs); err != nil {
		return err
	}

	return nil
}
