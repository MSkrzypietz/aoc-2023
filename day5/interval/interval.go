package interval

type Interval struct {
	Min int
	Max int
}

func (i Interval) Contains(other Interval) bool {
	return i.Min <= other.Min && other.Max <= i.Max
}

func (i Interval) Overlaps(other Interval) (Interval, bool) {
	if i.Min > other.Min || other.Min > i.Max {
		return Interval{Min: 0, Max: 0}, false
	}
	return Interval{Min: other.Min, Max: i.Max}, true
}
