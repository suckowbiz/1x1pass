package school

import (
	"github.com/suckowbiz/1x1pass/internal/calculus"
	"github.com/suckowbiz/1x1pass/internal/interaction"
)

// Lesson represents a school lesson.
type Lesson struct {
}

// NewLesson constructs a new Lesson.
func NewLesson() Lesson {
	return Lesson{}
}

// Run initiates and runs a lesson until duration elapsed.
func (l Lesson) Run(io interaction.CmdLiner, observer Observable) {
	var problem calculus.Problemer
	io.PrintStart()
	for {
		observer.RoundTick()
		if observer.RepeatRound() {
			problem, _ = observer.PopMissed()
		} else {
			problem = calculus.GenerateProblem([]calculus.Arithmetic{calculus.Multiplication, calculus.Division})
		}

		answer := io.AskProblem(observer.Rounds(), problem.Operator(), problem.LOperand(), problem.ROperand())
		if answer == problem.Solve() {
			if observer.Over() {
				io.PrintDone(observer.Rounds(), observer.Miscalculations())
				break
			}
		} else {
			io.PrintFailure()
			observer.PushMissed(problem)
		}
	}
}
