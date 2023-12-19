package day19

import (
	"strconv"
	"strings"
)

func PartOne() int {
	// INPUT = "px{a<2006:qkq,m>2090:A,rfg}\npv{a>1716:R,A}\nlnx{m>1548:A,A}\nrfg{s<537:gd,x>2440:R,A}\nqs{s>3448:A,lnx}\nqkq{x<1416:A,crn}\ncrn{x>2662:A,R}\nin{s<1351:px,qqz}\nqqz{s>2770:qs,m<1801:hdj,R}\ngd{a>3333:R,R}\nhdj{m>838:A,pv}\n\n{x=787,m=2655,a=1222,s=2876}\n{x=1679,m=44,a=2067,s=496}\n{x=2036,m=264,a=79,s=2244}\n{x=2461,m=1339,a=466,s=291}\n{x=2127,m=1623,a=2188,s=1013}" // 19114
	sorter := newSorter(INPUT)
	sorter.Sort("in")

	return sorter.AcceptedSum()
}

func PartTwo() int {
	return -1
}

func newSorter(input string) *Sorter {
	s := &Sorter{
		Workflows: map[string]*Workflow{},
		Parts:     []*Part{},
		Accepted:  []*Part{},
		Rejected:  []*Part{},
	}

	sections := strings.Split(input, "\n\n")

	// Workflows and rules
	for _, line := range strings.Split(sections[0], "\n") {
		wf := &Workflow{Rules: []*Rule{}, Raw: line}
		line = strings.TrimSuffix(line, "}")
		parts := strings.Split(line, "{")
		wf.Name = parts[0]

		for _, item := range strings.Split(parts[1], ",") {
			rule := &Rule{}
			if strings.Contains(item, ":") {
				pieces := strings.Split(item, ":")
				item = pieces[1]
				piecesSlice := []byte(pieces[0])
				rule.Category = piecesSlice[0]
				rule.Operator = piecesSlice[1]
				rule.Number = convertToNum(string(piecesSlice[2:]))
			}

			rule.Next = item
			wf.Rules = append(wf.Rules, rule)
		}

		s.Workflows[wf.Name] = wf
	}

	// Parts
	for _, line := range strings.Split(sections[1], "\n") {
		part := &Part{}

		line = strings.TrimPrefix(line, "{")
		line = strings.TrimSuffix(line, "}")

		for _, item := range strings.Split(line, ",") {
			pieces := strings.Split(item, "=")
			switch {
			case pieces[0] == "x":
				part.X = convertToNum(pieces[1])
			case pieces[0] == "m":
				part.M = convertToNum(pieces[1])
			case pieces[0] == "a":
				part.A = convertToNum(pieces[1])
			case pieces[0] == "s":
				part.S = convertToNum(pieces[1])
			}
		}
		s.Parts = append(s.Parts, part)
	}

	return s
}

type Sorter struct {
	Workflows map[string]*Workflow
	Parts     []*Part
	Accepted  []*Part
	Rejected  []*Part
}

func (s *Sorter) AcceptedSum() int {
	total := 0

	for _, accepted := range s.Accepted {
		total += accepted.X + accepted.M + accepted.A + accepted.S
	}

	return total
}

func (s *Sorter) Sort(initWorkFlow string) {
	for _, p := range s.Parts {
		currentWF := initWorkFlow

		for currentWF != "A" && currentWF != "R" {
			wf := s.Workflows[currentWF]
			for _, rule := range wf.Rules {
				result := rule.Apply(p)
				if result != "" {
					currentWF = result
					break
				}
			}
		}

		if currentWF == "A" {
			s.Accepted = append(s.Accepted, p)
		} else if currentWF == "R" {
			s.Rejected = append(s.Rejected, p)
		}
	}
}

type Workflow struct {
	Name  string
	Raw   string
	Rules []*Rule
}

func (w *Workflow) Process(part *Part) string {
	return ""
}

type Rule struct {
	Category byte
	Operator byte
	Number   int
	Next     string
}

func (r *Rule) Apply(part *Part) string {
	var value int
	switch r.Category {
	case 'x':
		value = part.X
	case 'm':
		value = part.M
	case 'a':
		value = part.A
	case 's':
		value = part.S
	case 0:
		return r.Next
	}

	if r.Operator == '>' && value > r.Number {
		return r.Next
	} else if r.Operator == '<' && value < r.Number {
		return r.Next
	}

	return ""
}

type Part struct {
	X, M, A, S int
}

func convertToNum(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}

	return result
}
