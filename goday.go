package goday

import (
	"time"
)

type Goday struct {
	time time.Time
}

// Now returns a new Goday instance initialized with the current time.
func Now() *Goday {
	return &Goday{time: time.Now()}
}

// StartOf sets the time to the start of the given unit (e.g., 'month').
func (d *Goday) StartOf(unit string) *Goday {
	switch unit {
	case "year":
		d.time = time.Date(d.time.Year(), time.January, 1, 0, 0, 0, 0, d.time.Location())
	case "month":
		d.time = time.Date(d.time.Year(), d.time.Month(), 1, 0, 0, 0, 0, d.time.Location())
	case "day":
		d.time = time.Date(d.time.Year(), d.time.Month(), d.time.Day(), 0, 0, 0, 0, d.time.Location())
	}
	return d
}

// Add adds the given duration to the current time.
func (d *Goday) Add(value int, unit string) *Goday {
	var dur time.Duration
	switch unit {
	case "year":
		dur = time.Duration(value) * 365 * 24 * time.Hour
	case "month":
		dur = time.Duration(value) * 30 * 24 * time.Hour
	case "day":
		dur = time.Duration(value) * 24 * time.Hour
	}
	d.time = d.time.Add(dur)
	return d
}

// Set sets the specified component of the time to the given value.
func (d *Goday) Set(unit string, value int) *Goday {
	switch unit {
	case "year":
		d.time = time.Date(value, d.time.Month(), d.time.Day(), d.time.Hour(), d.time.Minute(), d.time.Second(), d.time.Nanosecond(), d.time.Location())
	case "month":
		d.time = time.Date(d.time.Year(), time.Month(value), d.time.Day(), d.time.Hour(), d.time.Minute(), d.time.Second(), d.time.Nanosecond(), d.time.Location())
	case "day":
		d.time = time.Date(d.time.Year(), d.time.Month(), value, d.time.Hour(), d.time.Minute(), d.time.Second(), d.time.Nanosecond(), d.time.Location())
	}
	return d
}

// Format returns the formatted time string using the given layout.
func (d *Goday) Format(layout string) string {
	return d.time.Format(layout)
}
