package mapper

import (
	"log"

	"github.com/kohpai/tcas-3rd-round-resolver/model"
)

type RankingMap map[string]model.Ranking          // course ID -> citizen ID -> rank
type JointCourseMap map[string]*model.JointCourse // joint ID -> joint course
type CourseMap map[string]*model.Course           // course ID -> course
type StudentMap map[string]*model.Student         // citizen ID -> student

type RankInfoMap map[string]RankInfo       // citizen ID -> rank info
type RankingInfoMap map[string]RankInfoMap // course ID -> citizen ID -> rank info
type CourseInfoMap map[string]CourseInfo   // course ID -> course info
type StudentInfoMap map[string]StudentInfo // citizen ID -> student info

func createRankingMap(rankingInfoMap RankingInfoMap) RankingMap {
	rankingMap := make(RankingMap)

	for courseId, rankInfoMap := range rankingInfoMap {
		ranking := make(model.Ranking)
		for citizenId, rankInfo := range rankInfoMap {
			ranking[citizenId] = rankInfo.Rank
		}
		rankingMap[courseId] = ranking
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

func ExtractRankings(rankings []Ranking) (RankingInfoMap, CourseInfoMap, StudentInfoMap) {
	rankInfoMap := make(RankingInfoMap)
	courseInfoMap := make(CourseInfoMap)
	studentInfoMap := make(StudentInfoMap)

	for _, r := range rankings {
		courseId := r.CourseId
		citizenId := r.CitizenId

		if _, ok := courseInfoMap[courseId]; !ok {
			courseInfoMap[courseId] = CourseInfo{
				UniversityId:   r.UniversityId,
				UniversityName: r.UniversityName,
				CourseId:       courseId,
				FacultyName:    r.FacultyName,
				CourseName:     r.CourseName,
				ProjectName:    r.ProjectName,
			}
		}

		if _, ok := studentInfoMap[citizenId]; !ok {
			studentInfoMap[citizenId] = StudentInfo{
				CitizenId:   r.CitizenId,
				Title:       r.Title,
				FirstName:   r.FirstName,
				LastName:    r.LastName,
				PhoneNumber: r.PhoneNumber,
				Email:       r.Email,
			}
		}

		if _, ok := rankInfoMap[courseId]; !ok {
			rankInfoMap[courseId] = make(RankInfoMap)
		}

		rankInfoMap[courseId][citizenId] = RankInfo{
			ApplicationId:     r.ApplicationId,
			ApplicationDate:   r.ApplicationDate,
			InterviewLocation: r.InterviewLocation,
			InterviewDate:     r.InterviewDate,
			InterviewTime:     r.InterviewTime,
			Rank:              r.Rank,
			Round:             r.Round,
		}
	}

	return rankInfoMap, courseInfoMap, studentInfoMap
}

func CreateCourseMap(courses []Course, rankingInfoMap RankingInfoMap) CourseMap {
	jointCourseMap := createJointCourseMap(courses)
	rankingMap := createRankingMap(rankingInfoMap)
	courseMap := make(CourseMap)

	for _, c := range courses {
		var jointCourse *model.JointCourse
		if c.JointId == "" {
			jointCourse = jointCourseMap[c.Id]
		} else {
			jointCourse = jointCourseMap[c.JointId]
		}

		courseId := c.Id
		courseMap[courseId] = model.NewCourse(
			courseId,
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

func ToOutput(
	students []*model.Student,
	courseInfoMap CourseInfoMap,
	studentInfoMap StudentInfoMap,
	rankingInfoMap RankingInfoMap,
) []Ranking {
	outputs := make([]Ranking, 0, len(students)*6)

	for _, student := range students {
		courseIndex := student.CourseIndex()

		for i := 0; i < 6; i++ {
			course, _ := student.PreferredCourse(uint8(i) + 1)

			if course == nil {
				continue
			}

			citizenId := student.CitizenId()
			rank := course.Ranking()[citizenId]
			if rank == 0 {
				continue
			}

			courseId := course.Id()
			courseInfo := courseInfoMap[courseId]
			studentInfo := studentInfoMap[citizenId]
			rankInfo := rankingInfoMap[courseId][citizenId]
			output := Ranking{
				UniversityId:      courseInfo.UniversityId,
				UniversityName:    courseInfo.UniversityName,
				CourseId:          courseId,
				FacultyName:       courseInfo.FacultyName,
				CourseName:        courseInfo.CourseName,
				ProjectName:       courseInfo.ProjectName,
				CitizenId:         citizenId,
				Title:             studentInfo.Title,
				FirstName:         studentInfo.FirstName,
				LastName:          studentInfo.LastName,
				PhoneNumber:       studentInfo.PhoneNumber,
				Email:             studentInfo.Email,
				ApplicationId:     rankInfo.ApplicationId,
				ApplicationDate:   rankInfo.ApplicationDate,
				InterviewLocation: rankInfo.InterviewLocation,
				InterviewDate:     rankInfo.InterviewDate,
				InterviewTime:     rankInfo.InterviewTime,
				Rank:              rank,
				Round:             rankInfo.Round,
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
