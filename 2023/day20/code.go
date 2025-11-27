package day20

import (
	"strings"
)

type Pulse int

const (
	lowPulse Pulse = iota
	highPulse
)

func PartOne() int {
	// INPUT = "broadcaster -> a, b, c\n%a -> b\n%b -> c\n%c -> inv\n&inv -> a"         // 32000000
	// INPUT = "broadcaster -> a\n%a -> inv, con\n&inv -> b\n%b -> con\n&con -> output" // 11687500
	controller := parse(INPUT)

	for i := 0; i < 1000; i++ {
		controller.PressButton()
	}

	return controller.LowPulses * controller.HighPulses
}

func PartTwo() int {
	return -1
}

func parse(input string) *Controller {
	controller := &Controller{Modules: map[string]*Module{}, Queue: []*Flow{}}

	for _, line := range strings.Split(input, "\n") {
		mod := &Module{
			Controller: controller,
			LastPulse:  lowPulse,
			Memory:     map[string]Pulse{},
		}
		parts := strings.Split(line, " -> ")
		name := parts[0]
		mod.Connections = strings.Fields(strings.ReplaceAll(parts[1], ",", ""))

		if parts[0] == "broadcaster" {
			controller.Broadcaster = mod
			continue
		}

		if trimmed, found := strings.CutPrefix(name, "%"); found {
			mod.IsFlipFlop = true
			name = trimmed
		}
		if trimmed, found := strings.CutPrefix(name, "&"); found {
			mod.IsConjunction = true
			name = trimmed
		}

		mod.Name = name
		controller.Modules[name] = mod
	}

	conjunctions := map[string]bool{}
	for _, mod := range controller.Modules {
		if mod.IsConjunction {
			conjunctions[mod.Name] = true
		}
	}

	for _, mod := range controller.Modules {
		for _, connection := range mod.Connections {
			if _, ok := conjunctions[connection]; ok {
				controller.Modules[connection].Memory[mod.Name] = lowPulse
			}
		}
	}

	return controller
}

type Module struct {
	Name          string
	Controller    *Controller
	IsFlipFlop    bool
	IsConjunction bool
	On            bool
	LastPulse     Pulse
	Connections   []string
	Memory        map[string]Pulse
}

type Flow struct {
	Pulse  Pulse
	Module *Module
	Sender string
}

func (m *Module) Conjunction(pulse Pulse, sender string) {
	m.Memory[sender] = pulse

	nextPulse := lowPulse
	for _, p := range m.Memory {
		if p == lowPulse {
			nextPulse = highPulse
			break
		}
	}

	for _, name := range m.Connections {
		if nextPulse == lowPulse {
			m.Controller.LowPulses++
		} else {
			m.Controller.HighPulses++
		}

		tmpMod, ok := m.Controller.Modules[name]
		if ok {
			m.Controller.Queue = append(m.Controller.Queue, &Flow{Module: tmpMod, Pulse: nextPulse, Sender: m.Name})
		}
	}
}

func (m *Module) FlipFlop(pulse Pulse, _ string) {
	if pulse == highPulse {
		return
	}

	nextPulse := lowPulse
	if !m.On {
		nextPulse = highPulse
	}

	for _, name := range m.Connections {
		tmpMod, ok := m.Controller.Modules[name]
		if ok {
			if nextPulse == lowPulse {
				m.Controller.LowPulses++
			} else {
				m.Controller.HighPulses++
			}
			m.Controller.Queue = append(m.Controller.Queue, &Flow{Module: tmpMod, Pulse: nextPulse, Sender: m.Name})
		} else {
			panic(name + " mod not found")
		}
	}

	m.On = !m.On
}

type Controller struct {
	Modules     map[string]*Module
	Queue       []*Flow
	Broadcaster *Module
	LowPulses   int
	HighPulses  int
}

func (c *Controller) PressButton() {
	c.LowPulses++

	for _, name := range c.Broadcaster.Connections {
		mod := c.Modules[name]
		c.LowPulses++
		c.Queue = append(c.Queue, &Flow{Module: mod, Pulse: lowPulse, Sender: "broadcaster"})
	}

	for len(c.Queue) != 0 {
		flow := c.Queue[0]
		c.Queue = c.Queue[1:]

		if flow.Module.IsFlipFlop {
			flow.Module.FlipFlop(flow.Pulse, flow.Sender)
		} else if flow.Module.IsConjunction {
			flow.Module.Conjunction(flow.Pulse, flow.Sender)
		} else {
			panic("WHAT!?")
		}
	}
}
