package mapper

type Course struct {
	Id        string `json:"course_id"`
	JointId   string `json:"round_3_join_id"`
	Limit     uint16 `json:"round_3_receive"`
	Condition uint8  `json:"round_3_condition"`
	AddLimit  uint16 `json:"round_3_add_limit"`
}

type Student struct {
	CitizenId string `json:"citizen_id"`
	CourseId  string `json:"course_id"`
	Priority  uint8  `json:"priority"`
}

type Ranking struct {
	CourseId  string `csv:"course_id"`
	CitizenId string `csv:"citizen_id"`
	Rank      uint8  `csv:"ranking"`
}
