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
			Name:  "applications",
			Value: "apps.csv/universities.csv",
			Usage: "a file to input enrollment applications",
		},
		cli.StringFlag{
			Name:  "courses",
			Value: "courses.csv/round.csv",
			Usage: "a file to input all courses in the system",
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
	apps, err := util.ReadApps(c.String("applications"))
	if err != nil {
		return errors.Wrap(err, "applications")
	}
	courses, err := util.ReadCourses(c.String("courses"))
	if err != nil {
		return errors.Wrap(err, "courses")
	}

	rankingMap := mapper.ExtractRankings(apps)
	courseMap := mapper.CreateCourseMap(
		courses,
		rankingMap,
	)

	clearingHouse := model.NewClearingHouse(
		util.GetPendingStudents(
			mapper.CreateStudentMap(
				apps,
				courseMap,
			),
		),
	)

	clearingHouse.Execute()

	allStudents := clearingHouse.Students()

	// @ASSERTION, this shouldn't happen!
	if len(allStudents) != len(clearingHouse.AcceptedStudents())+len(clearingHouse.RejectedStudents()) {
		return errors.New("some students are missing")
	}

	outputs := mapper.ToOutput(allStudents, courseMap)
	if err := util.WriteCsvFile(c.String("output"), &outputs); err != nil {
		return err
	}

	return nil
}
