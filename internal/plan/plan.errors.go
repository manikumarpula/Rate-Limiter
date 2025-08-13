package plan

// ErrPlanNotFound is returned when a plan is not found
type ErrPlanNotFound struct{}

func (e ErrPlanNotFound) Error() string {
	return "plan not found"
}
