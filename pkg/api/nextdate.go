package api

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func afterNow(date, now time.Time) bool {
	if date.Year() > now.Year() {
		return true
	}
	if date.Year() == now.Year() && date.Month() > now.Month() {
		return true
	}
	if date.Year() == now.Year() && date.Month() == now.Month() &&
		date.Day() > now.Day() {
		return true
	}
	return false
}

func equalDate(dx, dy time.Time) bool {
	return dx.Year() == dy.Year() && dx.Month() == dy.Month() &&
		dx.Day() == dy.Day()
}

func lastDay(month, year int) int {
	var daysMonth = []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	last := daysMonth[month]
	if month == 2 {
		if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
			last = 29
		}
	}
	return last
}

func NextDate(now time.Time, d string, format string) (string, error) {
	errMsg := "wrong repeat format %s"

	date, err := time.Parse(formatDate, d)
	if err != nil {
		return ``, err
	}
	if len(format) == 0 {
		return ``, fmt.Errorf(errMsg, format)
	}
	pars := strings.Split(format, ` `)
	if format[0] == 'y' {
		if date.Year() >= now.Year() {
			date = date.AddDate(1, 0, 0)
		} else {
			next := time.Date(now.Year(), date.Month(), date.Day(), 12, 0, 0, 0, time.Local)
			if equalDate(date, next) || afterNow(now, next) {
				next = next.AddDate(1, 0, 0)
			}
			date = next
		}
		return date.Format(formatDate), nil
	}
	if len(pars) < 2 {
		return d, fmt.Errorf(errMsg, format)
	}
	switch format[0] {
	case 'm': // m 1,3,4,7,15,-1,-2 1,3,4,12
		var day [32]bool
		var month [13]bool

		days := strings.Split(pars[1], `,`)
		var (
			is, isMonth   bool
			last, prelast bool
		)
		for _, v := range days {
			mday, err := strconv.ParseInt(strings.TrimSpace(v), 10, 32)
			if err != nil {
				return d, fmt.Errorf(errMsg, format)
			}
			if mday >= 1 && mday <= 31 {
				is = true
				day[mday] = true
			} else if mday == -1 {
				last = true
			} else if mday == -2 {
				prelast = true
			} else {
				return ``, fmt.Errorf(errMsg, format)
			}
		}
		if len(pars) > 2 {
			months := strings.Split(pars[2], `,`)
			for _, v := range months {
				m, err := strconv.ParseUint(strings.TrimSpace(v), 10, 32)
				if err != nil {
					return d, fmt.Errorf(errMsg, format)
				}
				if m >= 1 && m <= 12 {
					isMonth = true
					month[m] = true
				} else {
					return ``, fmt.Errorf(errMsg, format)
				}
			}
		}
		for {
			var next bool
			for !next {
				date = date.AddDate(0, 0, 1)
				if isMonth && !month[date.Month()] {
					continue
				}
				if is {
					next = day[date.Day()]
				}
				if !next && (prelast || last) {
					lastDay := lastDay(int(date.Month()), date.Year())
					next = (last && date.Day() == lastDay) ||
						(prelast && date.Day() == lastDay-1)
				}
			}
			if afterNow(date, now) {
				break
			}
		}
	case 'w': // w 1,4,5,7
		var week [7]bool

		weeks := strings.Split(pars[1], `,`)
		var is bool
		for _, v := range weeks {
			wday, err := strconv.ParseUint(strings.TrimSpace(v), 10, 32)
			if err != nil {
				return d, fmt.Errorf(errMsg, format)
			}
			if wday < 1 || wday > 7 {
				return ``, fmt.Errorf(errMsg, format)
			}
			is = true
			if wday == 7 {
				week[0] = true
			} else {
				week[wday] = true
			}
		}
		if !is {
			date = now.AddDate(0, 0, 1)
			break
		}
		for {
			var next bool
			for !next {
				date = date.AddDate(0, 0, 1)
				next = week[date.Weekday()]
			}
			if afterNow(date, now) {
				break
			}
		}
	case 'd': // d 14
		interval, err := strconv.ParseUint(pars[1], 10, 32)
		if err != nil {
			return d, fmt.Errorf(errMsg, format)
		}
		if interval > 400 {
			return d, fmt.Errorf(errMsg, format)
		}
		for {
			date = date.AddDate(0, 0, int(interval))
			if afterNow(date, now) {
				break
			}
		}
	default:
		return "", fmt.Errorf(errMsg, format)
	}
	return date.Format(formatDate), nil
}

func nextDayHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	if v := r.FormValue("now"); len(v) > 0 {
		var err error
		now, err = time.Parse(`20060102`, v)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
	}
	out, err := NextDate(now, r.FormValue("date"), r.FormValue("repeat"))
	if err != nil {
		out = err.Error()
	}
	io.WriteString(w, out)
}
