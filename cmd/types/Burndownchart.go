package types

type Burndownchart struct {
	Project      Project
	Developer    Developer
	TasksCostSum int64
	WorkingHours map[string]int64
}
