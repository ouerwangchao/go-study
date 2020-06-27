package ch2

import (
	"encoding/json"
	"fmt"
	"testing"
)

type BasicInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type JobInfo struct {
	Skills []string `json:"skills"`
}

type Employee struct {
	BasicInfo BasicInfo `json:"basic_info"`
	JobInfo   JobInfo   `json:"job_info"`
}

var jsonStr = `{
	"basic_info":{
		"name":"Mike",
		"age" : 30
	},
	"job_info":{
		"skills":["Java","Go","C"]
	}
}`

func TestJson(t *testing.T) {
	//var employee Employee
	e := new(Employee)
	//t.Log(&employee)
	//t.Log(e)
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(*e)
	if v, err := json.Marshal(e); err == nil {
		t.Log(string(v))
	} else {
		t.Error(err)
	}
	//t.Log(employee)
	//employee.JobInfo.Skills = []string{"a", "b"}
	//t.Log(employee)
}

func TestA(t *testing.T) {
	a, b := 3, 4
	t.Log(swap(a, b))
}
func swap(a, b int) (int, int) {
	a, b = b, a
	return a, b
}
