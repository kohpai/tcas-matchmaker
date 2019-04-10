package mapper

import (
	"log"

	"github.com/kohpai/tcas-3rd-round-resolver/model"
)

type JointCourseMap map[string]*model.JointCourse // joint ID -> joint course
type CourseMap map[string]*model.Course           // course ID -> course
type StudentMap map[string]*model.Student         // citizen ID -> student

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

func CreateCourseMap(courses []Course) CourseMap {
	jointCourseMap := createJointCourseMap(courses)
	courseMap := make(CourseMap)

	for _, c := range courses {
		var jointCourse *model.JointCourse
		if c.JointId == "" {
			jointCourse = jointCourseMap[c.Id]
		} else {
			jointCourse = jointCourseMap[c.JointId]
		}

		courseMap[c.Id] = model.NewCourse(c.Id, jointCourse)
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
