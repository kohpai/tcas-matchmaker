package mapper

import (
	"log"
	"strconv"

	as "github.com/kohpai/tcas-3rd-round-resolver/model/applystrategy"
	"github.com/kohpai/tcas-3rd-round-resolver/model/common"
	"github.com/kohpai/tcas-3rd-round-resolver/model/course"
	jc "github.com/kohpai/tcas-3rd-round-resolver/model/jointcourse"
	st "github.com/kohpai/tcas-3rd-round-resolver/model/student"
)

type RankingMap map[string]common.Ranking      // course ID -> citizen ID -> rank
type JointCourseMap map[string]*jc.JointCourse // joint ID -> joint course
type CourseMap map[string]*course.Course       // course ID -> course
type StudentMap map[string]*st.Student         // citizen ID -> student

type RankInfoMap map[string]RankInfo       // citizen ID -> rank info
type RankingInfoMap map[string]RankInfoMap // course ID -> citizen ID -> rank info

func createRankingMap(rankingInfoMap RankingInfoMap) RankingMap {
	rankingMap := make(RankingMap)

	for courseId, rankInfoMap := range rankingInfoMap {
		ranking := make(common.Ranking)
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
		condition, err := strconv.Atoi(c.Condition)
		if err != nil {
			log.Fatal("condition cannot be parsed", err)
		}
		strategy := as.NewApplyStrategy(course.Condition(condition), c.AddLimit)
		if c.JointId == "" {
			jointCourseMap[c.Id] = jc.NewJointCourse(c.Id, c.Limit, strategy)
		} else if _, ok := jointCourseMap[c.JointId]; !ok {
			jointCourseMap[c.JointId] = jc.NewJointCourse(c.JointId, c.Limit, strategy)
		}
	}

	return jointCourseMap
}

func ExtractRankings(rankings []Ranking) RankingInfoMap {
	rankInfoMap := make(RankingInfoMap)

	for _, r := range rankings {
		courseId := r.CourseId
		citizenId := r.CitizenId

		if _, ok := rankInfoMap[courseId]; !ok {
			rankInfoMap[courseId] = make(RankInfoMap)
		}

		rankInfoMap[courseId][citizenId] = RankInfo{
			ApplicationDate:   r.ApplicationDate,
			InterviewLocation: r.InterviewLocation,
			InterviewDate:     r.InterviewDate,
			InterviewTime:     r.InterviewTime,
			Rank:              r.Rank,
			Round:             r.Round,
			// course
			UniversityId:   r.UniversityId,
			UniversityName: r.UniversityName,
			CourseId:       r.CourseId,
			FacultyName:    r.FacultyName,
			CourseName:     r.CourseName,
			ProjectName:    r.ProjectName,
			// student
			CitizenId:   r.CitizenId,
			Title:       r.Title,
			FirstName:   r.FirstName,
			LastName:    r.LastName,
			PhoneNumber: r.PhoneNumber,
			Email:       r.Email,
		}
	}

	return rankInfoMap
}

func CreateCourseMap(courses []Course, rankingInfoMap RankingInfoMap) CourseMap {
	jointCourseMap := createJointCourseMap(courses)
	rankingMap := createRankingMap(rankingInfoMap)
	courseMap := make(CourseMap)

	for _, c := range courses {
		var jointCourse *jc.JointCourse
		if c.JointId == "" {
			jointCourse = jointCourseMap[c.Id]
		} else {
			jointCourse = jointCourseMap[c.JointId]
		}

		courseId := c.Id
		courseMap[courseId] = course.NewCourse(
			courseId,
			jointCourse,
			rankingMap[courseId],
		)
	}

	return courseMap
}

func CreateStudentMap(students []Student, courseMap CourseMap) StudentMap {
	studentMap := make(StudentMap)

	for _, s := range students {
		citizenId := s.CitizenId
		if _, ok := studentMap[citizenId]; !ok {
			studentMap[citizenId] = st.NewStudent(citizenId)
		}

		if err := studentMap[citizenId].SetPreferredCourse(s.Priority, courseMap[s.CourseId]); err != nil {
			log.Fatal("could not set preferred course", err)
		}
	}

	return studentMap
}

func ToOutput(students []*st.Student, rankingInfoMap RankingInfoMap) []Ranking {
	outputs := make([]Ranking, 0, len(students)*6)

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

			courseId := course.Id()
			rankInfo := rankingInfoMap[courseId][citizenId]
			output := Ranking{
				UniversityId:      rankInfo.UniversityId,
				UniversityName:    rankInfo.UniversityName,
				CourseId:          rankInfo.CourseId,
				FacultyName:       rankInfo.FacultyName,
				CourseName:        rankInfo.CourseName,
				ProjectName:       rankInfo.ProjectName,
				CitizenId:         rankInfo.CitizenId,
				Title:             rankInfo.Title,
				FirstName:         rankInfo.FirstName,
				LastName:          rankInfo.LastName,
				PhoneNumber:       rankInfo.PhoneNumber,
				Email:             rankInfo.Email,
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
