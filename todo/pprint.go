package todo

import (
	"fmt"
	"strings"
	"time"
)

type TodoLine Item

func (tl TodoLine) Pprint() string {

	// Standard Todo Printf format; e.g.
	// X  (1)  This is a description  +project  @tag1 @tag2 @tag3
	//    (2)  Another description    +chores   @tag1 @tag3
	format := "%v\t(%v)\t%v\t%v\t|\t%v\t%v"

	doneStr := ""

	if tl.Done {
		doneStr = "X"
	}

	projectStr := fmt.Sprintf("+%v", tl.Project)

	var tagsList []string

	for _, tag := range tl.Tags {
		tagsList = append(tagsList, fmt.Sprintf("@%v", tag))
	}

	dueDate := time.UnixMilli(tl.DueDate).Format("2006-01-02 15:04:05 (UTC)")

	return fmt.Sprintf(format, doneStr, tl.Priority, tl.Description, dueDate, projectStr, strings.Join(tagsList, " "))
}
