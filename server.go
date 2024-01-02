package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Course represents data about a record Course.
type course struct {
	CourseMajor  string `json:"CourseMajor"`
	CourseID     string `json:"CourseID"`
	ID           string `json:"ID"`
	CourseName   string `json:"CourseName" `
	CreditHour   int32  `json:"CreditHour"`
	LectureHour  int32  `json:"LectureHour"`
	LabHour      int32  `json:"LabHour"`
	Attribute    string `json:"Attribute"`
	GradeMode    string `json:"GradeMode"`
	Prerequisite string `json:"Prerequisite"`
	Corequisite  string `json:"Corequisite"`
}

// Courses slice to seed record Course data.
var courses = []course{
	{CourseMajor: "ANTH", CourseID: "ANTH 3308", ID: "3308", CourseName: "Cultural Resource Management and Archaeology", CreditHour: 3, LectureHour: 3, LabHour: 0, Attribute: "", GradeMode: "1", Prerequisite: "", Corequisite: ""},
	{CourseMajor: "ANTH", CourseID: "ANTH 3309", ID: "3309", CourseName: "Cultures Through Film", CreditHour: 3, LectureHour: 3, LabHour: 0, Attribute: "022", GradeMode: "1", Prerequisite: "", Corequisite: ""},
}

func main() {
	router := gin.Default()
	router.GET("/courses", getCourses)
	router.GET("/courses/:id", getCourseByID)
	router.POST("/courses", postCourses)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}

}

// getCourses responds with the list of all Courses as JSON.
func getCourses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, courses)
}

// postCourses adds a Course from JSON received in the request body.
func postCourses(c *gin.Context) {
	var newCourse course

	// Call BindJSON to bind the received JSON to
	// newCourse.
	if err := c.BindJSON(&newCourse); err != nil {
		return
	}

	// Add the new Course to the slice.
	courses = append(courses, newCourse)
	c.IndentedJSON(http.StatusCreated, newCourse)
}

// getCourseByID locates the Course whose ID value matches the id
// parameter sent by the client, then returns that Course as a response.
func getCourseByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of Courses, looking for
	// a Course whose ID value matches the parameter.
	for _, a := range courses {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "course not found"})
}
