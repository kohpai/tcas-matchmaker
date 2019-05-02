package mapper

import (
	"log"

	"github.com/kohpai/tcas-3rd-round-resolver/model"
	"github.com/kohpai/tcas-3rd-round-resolver/model/common"
	st "github.com/kohpai/tcas-3rd-round-resolver/model/student"
)

type RankingMap map[string]common.Ranking         // course ID -> citizen ID -> rank
type JointCourseMap map[string]*model.JointCourse // joint ID -> joint course
type CourseMap map[string]*model.Course           // course ID -> course
type StudentMap map[string]*st.Student            // citizen ID -> student

func createRankingMap(rankings []Ranking) RankingMap {
	rankingMap := make(RankingMap)

	for _, r := range rankings {
		courseId := r.CourseId
		if rankingMap[courseId] == nil {
			rankingMap[courseId] = make(common.Ranking)
		}
		rankingMap[courseId][r.CitizenId] = r.Rank
	}

	return rankingMap
}

func createJointCourseMap(courses []Course) JointCourseMap {
	jointCourseMap := make(JointCourseMap)

	for _, c := range courses {
		strategy := model.NewApplyStrategy(c.Condition, c.AddLimit)
		if c.JointId == "" {
			jointCourseMap[c.Id] = model.NewJointCourse(c.Id, c.Limit, strategy)
		} else if jointCourseMap[c.JointId] == nil {
			jointCourseMap[c.JointId] = model.NewJointCourse(c.JointId, c.Limit, strategy)
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
			studentMap[citizenId] = st.NewStudent(citizenId)
		}

		if err := studentMap[citizenId].SetPreferredCourse(s.Priority, courseMap[s.CourseId]); err != nil {
			log.Fatal("could not set preferred course", err)
		}
	}

	return studentMap
}

func ToOutput(students []*st.Student) []Output {
	outputs := make([]Output, 0)

	for _, student := range students {
		courseIndex := student.CourseIndex()

		for i := 0; i < 6; i++ {
			course, _ := student.PreferredCourse(i + 1)

			if course == nil {
				continue
			}

			citizenId := student.CitizenId()
			rank := course.Ranking()[citizenId]
			if rank == 0 {
				continue
			}

			output := Output{
				CourseId:  course.Id(),
				CitizenId: citizenId,
				Ranking:   rank,
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
