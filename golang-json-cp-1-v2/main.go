package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Report struct {
	// TODO: answer here
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Date     string  `json:"date"`
	Semester int     `json:"semester"`
	Studies  []Study `json:"studies"`
}

type Study struct {
	StudyName   string `json:"study_name"`
	StudyCredit int    `json:"study_credit"`
	Grade       string `json:"grade"`
}

// gunakan fungsi ini untuk mengambil data dari file json
// kembalian berupa struct 'Report' dan error
func ReadJSON(filename string) (Report, error) {
	// TODO: answer here
	var report Report
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return report, err
	}
	err = json.Unmarshal(data, &report)
	if err != nil {
		return report, err
	}
	return report, nil
}

func GradePoint(report Report) float64 {
	totalScore := 0.0
	totalCredits := 0

	for _, study := range report.Studies {
		var gradeValue float64
		switch study.Grade {
		case "A":
			gradeValue = 4.0
		case "AB":
			gradeValue = 3.5
		case "B":
			gradeValue = 3.0
		case "BC":
			gradeValue = 2.5
		case "C":
			gradeValue = 2.0
		case "CD":
			gradeValue = 1.5
		case "D":
			gradeValue = 1.0
		case "DE":
			gradeValue = 0.5
		case "E":
			gradeValue = 0.0
		default:
			gradeValue = 0.0
		}

		score := gradeValue * float64(study.StudyCredit)
		totalScore += score
		totalCredits += study.StudyCredit
	}

	if totalCredits != 0 {
		gpa := totalScore / float64(totalCredits)
		return gpa
	}
	return 0.0 // TODO: replace this
}

func main() {
	// bisa digunakan untuk menguji test case
	report, err := ReadJSON("report.json")
	if err != nil {
		panic(err)
	}

	gradePoint := GradePoint(report)
	fmt.Println(gradePoint)
}
