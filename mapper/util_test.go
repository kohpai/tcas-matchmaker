package mapper

import (
	"testing"
)

func TestExtractRankings_Always_ReturnRankingMap(t *testing.T) {
	apps := []Application{
		{
			CourseId:  "1234",
			CitizenId: "13499",
			Ranking:   1,
		},
		{
			CourseId:  "1234",
			CitizenId: "13501",
			Ranking:   2,
		},
		{
			CourseId:  "1234",
			CitizenId: "13502",
			Ranking:   3,
		},
		{
			CourseId:  "1235",
			CitizenId: "13500",
			Ranking:   2,
		},
		{
			CourseId:  "1236",
			CitizenId: "13499",
			Ranking:   1,
		},
		{
			CourseId:  "1237",
			CitizenId: "13500",
			Ranking:   1,
		},
		{
			CourseId:  "1237",
			CitizenId: "13499",
			Ranking:   2,
		},
	}

	rankingMap := ExtractRankings(apps)

	if rank := rankingMap["1234"]["13499"]; rank != 1 {
		t.Error("Rank is incorrect, got", rank, apps[0])
	}
	if rank := rankingMap["1234"]["13501"]; rank != 2 {
		t.Error("Rank is incorrect, got", rank, apps[1])
	}
	if rank := rankingMap["1234"]["13502"]; rank != 3 {
		t.Error("Rank is incorrect, got", rank, apps[2])
	}
	if rank := rankingMap["1235"]["13500"]; rank != 2 {
		t.Error("Rank is incorrect, got", apps[3])
	}
	if rank := rankingMap["1236"]["13499"]; rank != 1 {
		t.Error("Rank is incorrect, got", rank, apps[4])
	}
	if rank := rankingMap["1237"]["13500"]; rank != 1 {
		t.Error("Rank is incorrect, got", rank, apps[5])
	}
	if rank := rankingMap["1237"]["13499"]; rank != 2 {
		t.Error("Rank is incorrect, got", rank, apps[6])
	}
}

func TestCreateJointCourseMap_Always_ReturnsJointCourseMap(t *testing.T) {
	courses := []Course{
		{
			CourseId:            "1234",
			UniversityId:        "4321",
			JointId:             "",
			ReceiveAmount:       10,
			ExceedAllowedAmount: "A",
		},
		{
			CourseId:            "1235",
			UniversityId:        "4321",
			JointId:             "",
			ReceiveAmount:       11,
			ExceedAllowedAmount: "A",
		},
		{
			CourseId:            "1236",
			UniversityId:        "4322",
			JointId:             "123",
			ReceiveAmount:       12,
			ExceedAllowedAmount: "A",
		},
		{
			CourseId:            "1237",
			UniversityId:        "4322",
			JointId:             "123",
			ReceiveAmount:       12,
			ExceedAllowedAmount: "A",
		},
	}

	jointCourseMap := createJointCourseMap(courses)

	if jc := jointCourseMap["1234"]; jc.Students().AvailableSpots() != 10 {
		t.Error("Available spots is incorrect", jc)
	}

	if jc := jointCourseMap["1235"]; jc.Students().AvailableSpots() != 11 {
		t.Error("Available spots is incorrect", jc)
	}

	if jc := jointCourseMap["4322123"]; jc.Students().AvailableSpots() != 12 {
		t.Error("Available spots is incorrect", jc)
	}
}

func TestCreateCourseMap_Always_ReturnsCourseMap(t *testing.T) {
	apps := []Application{
		{
			CourseId:  "1234",
			CitizenId: "13499",
			Ranking:   1,
		},
		{
			CourseId:  "1234",
			CitizenId: "13501",
			Ranking:   2,
		},
		{
			CourseId:  "1234",
			CitizenId: "13502",
			Ranking:   3,
		},
		{
			CourseId:  "1235",
			CitizenId: "13500",
			Ranking:   2,
		},
		{
			CourseId:  "1236",
			CitizenId: "13499",
			Ranking:   1,
		},
		{
			CourseId:  "1237",
			CitizenId: "13500",
			Ranking:   1,
		},
		{
			CourseId:  "1237",
			CitizenId: "13499",
			Ranking:   2,
		},
	}
	rankingMap := ExtractRankings(apps)

	courses := []Course{
		{
			CourseId:            "1234",
			UniversityId:        "4321",
			JointId:             "",
			ReceiveAmount:       10,
			ExceedAllowedAmount: "A",
		},
		{
			CourseId:            "1235",
			UniversityId:        "4321",
			JointId:             "",
			ReceiveAmount:       11,
			ExceedAllowedAmount: "A",
		},
		{
			CourseId:            "1236",
			UniversityId:        "4322",
			JointId:             "123",
			ReceiveAmount:       12,
			ExceedAllowedAmount: "A",
		},
		{
			CourseId:            "1237",
			UniversityId:        "4322",
			JointId:             "123",
			ReceiveAmount:       12,
			ExceedAllowedAmount: "A",
		},
	}

	courseMap := CreateCourseMap(courses, rankingMap)

	if course := courseMap["1234"]; course.JointCourse().Students().AvailableSpots() != 10 {
		t.Error("Available spots is incorrect", course)
	}

	if course := courseMap["1235"]; course.JointCourse().Students().AvailableSpots() != 11 {
		t.Error("Available spots is incorrect", course)
	}

	if course := courseMap["1236"]; course.JointCourse().Students().AvailableSpots() != 12 {
		t.Error("Available spots is incorrect", course)
	}

	if course := courseMap["1237"]; course.JointCourse().Students().AvailableSpots() != 12 {
		t.Error("Available spots is incorrect", course)
	}

	jc1, jc2 := courseMap["1236"].JointCourse(), courseMap["1237"].JointCourse()

	if jc1 != jc2 {
		t.Error("Incorrect joint course", jc1, jc2)
	}

	if len(jc1.Courses()) != 2 {
		t.Error("Length of courses is incorrect", jc1)
	}
}

func TestCreateStudentMap_Always_ReturnsStudentMap(t *testing.T) {
	apps := []Application{
		{
			CourseId:      "1234",
			CitizenId:     "13499",
			Priority:      1,
			Ranking:       1,
			Gender:        1,
			SchoolProgram: 1,
		},
		{
			CourseId:      "1234",
			CitizenId:     "13501",
			Priority:      1,
			Ranking:       2,
			Gender:        1,
			SchoolProgram: 1,
		},
		{
			CourseId:      "1234",
			CitizenId:     "13502",
			Priority:      1,
			Ranking:       3,
			Gender:        1,
			SchoolProgram: 1,
		},
		{
			CourseId:      "1235",
			CitizenId:     "13500",
			Priority:      1,
			Ranking:       2,
			Gender:        1,
			SchoolProgram: 1,
		},
		{
			CourseId:      "1236",
			CitizenId:     "13499",
			Priority:      2,
			Ranking:       1,
			Gender:        1,
			SchoolProgram: 1,
		},
		{
			CourseId:      "1237",
			CitizenId:     "13500",
			Priority:      3,
			Ranking:       1,
			Gender:        1,
			SchoolProgram: 1,
		},
		{
			CourseId:      "1237",
			CitizenId:     "13499",
			Priority:      4,
			Ranking:       2,
			Gender:        1,
			SchoolProgram: 1,
		},
	}
	courses := []Course{
		{
			CourseId:            "1234",
			UniversityId:        "4321",
			JointId:             "",
			ReceiveAmount:       10,
			ExceedAllowedAmount: "A",
		},
		{
			CourseId:            "1235",
			UniversityId:        "4321",
			JointId:             "",
			ReceiveAmount:       11,
			ExceedAllowedAmount: "A",
		},
		{
			CourseId:            "1236",
			UniversityId:        "4322",
			JointId:             "123",
			ReceiveAmount:       12,
			ExceedAllowedAmount: "A",
		},
		{
			CourseId:            "1237",
			UniversityId:        "4322",
			JointId:             "123",
			ReceiveAmount:       12,
			ExceedAllowedAmount: "A",
		},
	}

	rankingMap := ExtractRankings(apps)
	courseMap := CreateCourseMap(courses, rankingMap)
	studentMap := CreateStudentMap(apps, courseMap)

	student1, student2 := studentMap["13499"], studentMap["13500"]
	if app, err := student1.Application(1); err != nil || app.Course().Id() != "1234" {
		t.Error("Course is incorrect", app.Course())
	}

	if app, err := student1.Application(2); err != nil || app.Course().Id() != "1236" {
		t.Error("Course is incorrect", app.Course())
	}

	if app, err := student2.Application(3); err != nil || app.Course().Id() != "1237" {
		t.Error("Course is incorrect", app.Course())
	}
}
