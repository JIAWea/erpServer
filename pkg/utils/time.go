package utils

import "time"

func StrToTime(s string) (time.Time, error) {
	if len(s) > 10 {
		layout := "2006-01-02 15:04:05"
		t, err := time.ParseInLocation(layout, s, time.Local)
		if err != nil {
			layout = "2006/01/02 15:04:05"
			t, err := time.ParseInLocation(layout, s, time.Local)
			if err != nil {
				return time.Time{}, err
			}
			return t, nil
		}
		return t, nil
	}

	layout := "2006/01/02"
	t, err := time.ParseInLocation(layout, s, time.Local)
	if err != nil {
		layout = "2006-01-02"
		t, err := time.ParseInLocation(layout, s, time.Local)
		if err != nil {
			return time.Time{}, err
		}
		return t, nil
	}
	return t, nil
}
