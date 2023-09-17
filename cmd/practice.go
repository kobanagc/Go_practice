package main
import (
	"fmt"
)

type Student struct {
	name string
}

func main() {
	student := Student{
		name : "luke",
	}
	average_score := student.avg([]float64{50, 20, 70, 80})
	fmt.Println(average_score)
	fmt.Println(student.judge(average_score))
}

func (s Student) avg(data[] float64) (average_score float64) {
	sum := 0.0
	for _, v := range data {
		sum += v
	}
	average_score = sum / float64(len(data))
	return
}

func (s Student) judge(average_score float64) (judge string) {
	if average_score >= 60 {
		judge = "passed"
	} else {
		judge = "failed"
	}
	return
}