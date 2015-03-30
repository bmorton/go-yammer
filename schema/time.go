package schema

import "time"

type Time time.Time

func (t Time) UnmarshalJSON(b []byte) error {
	const shortForm = "\"2006/01/02 15:04:05 -0700\"" // yyyymmdd date format

	parse, err := time.Parse(shortForm, string(b))
	if err != nil {
		return err
	}

	t = Time(parse)
	return nil
}
