package mapper

import (
	"log"

	"github.com/kohpai/tcas-3rd-round-resolver/model"
)

type RankingMap map[string]model.Ranking          // course ID -> citizen ID -> rank
type JointCourseMap map[string]*model.JointCourse // joint ID -> joint course
type CourseMap map[string]*model.Course           // course ID -> course
type StudentMap map[string]*model.Student         // citizen ID -> student

func createRankingMap(rankings []Ranking) RankingMap {
	rankingMap := make(RankingMap)

	for _, r := range rankings {
		courseId := r.CourseId
		if rankingMap[courseId] == nil {
			rankingMap[courseId] = make(model.Ranking)
		}
		rankingMap[courseId][r.CitizenId] = r.Rank
	}

	return rankingMap
}

func createJointCourseMap(courses []Course) JointCourseMap {
	jointCourseMap := make(JointCourseMap)

	for _, c := range courses {
		if c.JointId == "" {
			jointCourseMap[c.Id] = model.NewJointCourse(c.Id, c.Limit)
		} else if jointCourseMap[c.JointId] == nil {
			jointCourseMap[c.JointId] = model.NewJointCourse(c.JointId, c.Limit)
		}
	}

	return jointCourseMap
}

func CreateCourseMap(courses []Course, rankings []Ranking) CourseMap {
	jointCourseMap := createJointCourseMap(courses)
	rankingMap := createRankingMap(rankings)
	courseMap := make(CourseMap)

	for _, c := range courses {
		var jointCourse *model.JointCourse
		if c.JointId == "" {
			jointCourse = jointCourseMap[c.Id]
		} else {
			jointCourse = jointCourseMap[c.JointId]
		}

		courseMap[c.Id] = model.NewCourse(
			c.Id,
			c.Condition,
			jointCourse,
			rankingMap[c.Id],
		)
	}

	return courseMap
}

func CreateStudentMap(students []Student, courseMap CourseMap) StudentMap {
	studentMap := make(StudentMap)

	for _, s := range students {
		citizenId := s.CitizenId
		if studentMap[citizenId] == nil {
			studentMap[citizenId] = model.NewStudent(citizenId)
		}

		if err := studentMap[citizenId].SetPreferredCourse(s.Priority, courseMap[s.CourseId]); err != nil {
			log.Fatal("could not set preferred course", err)
		}
	}

	return studentMap
}

func ToOutput(students []*model.Student) []Output {
	outputs := make([]Output, 0)

	for _, student := range students {
		courseIndex := student.CourseIndex()

		for i := 0; i < 6; i++ {
			course, _ := student.PreferredCourse(uint8(i) + 1)

			if course == nil {
				continue
			}

			citizenId := student.CitizenId()
			ranking := course.Ranking()[citizenId]
			if ranking == 0 {
				continue
			}

			output := Output{
				CourseId:  course.Id(),
				CitizenId: citizenId,
				Ranking:   ranking,
			}

			statuses := AdmitStatuses()
			switch {
			case i < courseIndex || courseIndex == -1:
				output.AdmitStatus = statuses.Full
			case i == courseIndex:
				output.AdmitStatus = statuses.Admitted
			case i > courseIndex:
				output.AdmitStatus = statuses.Late
			}

			outputs = append(outputs, output)
		}
	}

	return outputs
}
