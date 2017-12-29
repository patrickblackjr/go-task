package task

import (
	"fmt"
	"sort"
	"text/tabwriter"
)

// PrintTasksHelp prints help os tasks that have a description
func (e *Executor) PrintTasksHelp() {
	tasks := e.tasksWithDesc()
	if len(tasks) == 0 {
		e.outf("task: No tasks with description available")
		return
	}
	e.outf("task: Available tasks for this project:")

	// Format in tab-separated columns with a tab stop of 8.
	w := tabwriter.NewWriter(e.Stdout, 0, 8, 0, '\t', 0)
	for _, task := range tasks {
		fmt.Fprintf(w, "* %s: \t%s\n", task.Task, task.Desc)
	}
	w.Flush()
}

func (e *Executor) tasksWithDesc() (tasks []*Task) {
	tasks = make([]*Task, 0, len(e.Taskfile.Tasks))
	for _, task := range e.Taskfile.Tasks {
		if task.Desc != "" {
			tasks = append(tasks, task)
		}
	}
	sort.Slice(tasks, func(i, j int) bool { return tasks[i].Task < tasks[j].Task })
	return
}
