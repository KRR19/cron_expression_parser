package cronparser

type CronFields struct {
    Minutes     []int
    Hours       []int
    DayOfMonth  []int
    Month       []int
    DayOfWeek   []int
    Command     *string
}
