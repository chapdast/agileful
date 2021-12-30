package db

import "fmt"

type Option func(o *option) error

func OptionLimit(limit int) Option {
	return func(o *option) error {
		if limit < 0 {
			return fmt.Errorf("invalid limit")
		}
		o.limit = limit
		return nil
	}
}
func OptionFilter(field string, condition Condition) Option {
	return func(o *option) error {
		if o.fillerBy == nil {
			o.fillerBy = map[string]Condition{}
		}
		o.fillerBy[field] = condition
		return nil
	}
}

func OptionOrder(field string, desc bool) Option {
	return func(o *option) error {
		o.orderBy = append(o.orderBy, []byte(field))
		o.orderDesc = desc
		return nil
	}
}

func OptionOffset(offset int) Option{
	return func(o *option) error {
		o.offset = offset
		return nil
	}
}