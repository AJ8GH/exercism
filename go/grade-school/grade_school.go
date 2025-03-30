package school

import (
	"slices"
)

type (
	School struct {
		gradesMap map[int]*Grade
	}

	Grade struct {
		level    int
		students []string
	}
)

func New() *School {
	return &School{gradesMap: map[int]*Grade{}}
}

func (s *School) Add(student string, grade int) {
	g := s.gradesMap[grade]
	if g == nil {
		g = &Grade{level: grade, students: []string{student}}
		s.gradesMap[grade] = g
		return
	}
	for _, v := range g.students {
		if v == student {
			return
		}
	}
	g.students = append(g.students, student)
	slices.Sort(g.students)
}

func (s *School) Grade(level int) []string {
	g := s.gradesMap[level]

	if g == nil {
		return []string{}
	}
	return g.students
}

func (s *School) Enrollment() (grades []Grade) {
	for _, v := range s.gradesMap {
		grades = append(grades, *v)
	}
	slices.SortFunc(grades, sortGrade)
	return grades
}

func sortGrade(a, b Grade) int {
	switch {
	case a.level < b.level:
		return -1
	case a.level > b.level:
		return 1
	default:
		return 0
	}
}
