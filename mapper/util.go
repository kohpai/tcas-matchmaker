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

		courseMap[c.Id] = model.NewCourse(c.Id, jointCourse, rankingMap[c.Id])
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

func ToOutput(acceptedStudents []*Student, rejectedStudents []*Student) []Output {
	outputs := make([]Output, len(acceptedStudents)+len(rejectedStudents))
}
