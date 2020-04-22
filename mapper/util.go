package mapper

import (
	"log"
	"strconv"

	"github.com/kohpai/tcas-matchmaker/model"
)

type RankingMap map[string]model.Ranking          // course ID -> citizen ID -> rank
type JointCourseMap map[string]*model.JointCourse // joint ID -> joint course
type CourseMap map[string]*model.Course           // course ID -> course
type StudentMap map[string]*model.Student         // citizen ID -> student

// type RankInfoMap map[string]RankInfo       // citizen ID -> rank info
// type RankingInfoMap map[string]RankInfoMap // course ID -> citizen ID -> rank info

func createJointCourseMap(courses []Course) JointCourseMap {
	jointCourseMap := make(JointCourseMap)

	for _, c := range courses {
		condition := c.ExceedAllowedAmount[0:1]
		exceedAllowed := 0
		if condition == "C" {
			var err error
			exceedAllowed, err = strconv.Atoi(c.ExceedAllowedAmount[1:])
			if err != nil {
				log.Fatal("exceed allowed amount cannot be parsed", err)
			}
		}
		strategy := model.NewApplyStrategy(model.Condition(condition), uint16(exceedAllowed))
		var jointId string
		if c.JointId == "" {
			jointId = c.CourseId
		} else if _, ok := jointCourseMap[c.JointId]; !ok {
			jointId = c.UniversityId + c.JointId
		}

		if jointId != "" {
			jointCourseMap[jointId] = model.NewJointCourse(
				jointId,
				model.NewAvailableSpots(
					c.ReceiveAmount,
					c.MaleReceiveAmount,
					c.FemaleReceiveAmount,
					c.FormalReceiveAmount,
					c.InterReceiveAmount,
					c.VocatReceiveAmount,
					c.NonformalReceiveAmount,
				),
				strategy,
			)
		}
	}

	return jointCourseMap
}

func ExtractRankings(rankings []Application) RankingMap {
	rankingMap := make(RankingMap)

	for _, r := range rankings {
		courseId := r.CourseId
		citizenId := r.CitizenId

		if _, ok := rankingMap[courseId]; !ok {
			rankingMap[courseId] = make(model.Ranking)
		}

		rankingMap[courseId][citizenId] = r.Ranking
	}

	return rankingMap
}

func CreateCourseMap(courses []Course, rankingMap RankingMap) CourseMap {
	jointCourseMap := createJointCourseMap(courses)
	courseMap := make(CourseMap)

	for _, c := range courses {
		var jointCourse *model.JointCourse
		if c.JointId == "" {
			jointCourse = jointCourseMap[c.CourseId]
		} else {
			jointCourse = jointCourseMap[c.UniversityId+c.JointId]
		}

		courseId := c.CourseId
		courseMap[courseId] = model.NewCourse(
			courseId,
			jointCourse,
			rankingMap[c.CourseId],
		)
	}

	return courseMap
}

func CreateStudentMap(applications []Application, courseMap CourseMap) StudentMap {
	studentMap := make(StudentMap)

	for _, a := range applications {
		citizenId := a.CitizenId
		if _, ok := studentMap[citizenId]; !ok {
			gender := a.Gender
			if gender < 1 || 2 < gender {
				log.Fatal("wrong gender value")
			}
			program := a.SchoolProgram
			if program < 1 || 4 < program {
				log.Fatal("wrong program value")
			}
			studentMap[citizenId] = model.NewStudent(
				citizenId,
				model.Gender(gender),
				model.Program(a.SchoolProgram),
			)
		}

		if err := studentMap[citizenId].SetPreferredApp(a.Priority, courseMap[a.CourseId], a.ApplicationId); err != nil {
			log.Fatal("could not set preferred course", err)
		}
	}

	return studentMap
}

func ToOutput(
	students []*model.Student,
	courseMap CourseMap,
) []Application {
	outputs := make([]Application, 0, len(students)*6)

	for _, student := range students {
		courseIndex := student.AppIndex()

		for i := 0; i < 6; i++ {
			appPriority := uint8(i) + 1
			app, _ := student.Application(appPriority)

			if app == nil {
				continue
			}

			citizenId := student.CitizenId()
			course := app.Course()
			rank := course.Ranking()[citizenId]
			if rank == 0 {
				continue
			}

			output := Application{
				ApplicationId: app.Id(),
				CitizenId:     citizenId,
				// Gender:           0,
				// SchoolProgram:    0,
				// FormalApplicable: 0,
				CourseId: course.Id(),
				Priority: uint8(i),
				Ranking:  rank,
			}

			statuses := AdmitStatuses()
			switch {
			case i < courseIndex || courseIndex == -1:
				output.Status = statuses.Full
			case i == courseIndex:
				output.Status = statuses.Admitted
			case i > courseIndex:
				output.Status = statuses.Late
			}

			outputs = append(outputs, output)
		}
	}

	return outputs
}
