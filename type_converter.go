package copier

import (
	"errors"
	"time"
)

var (
	Int64        int64
	PInt64       *int64 = nil
	Time         time.Time
	PTime        *time.Time = nil
	TimeToPInt64            = TypeConverter{
		SrcType: Time,
		DstType: PInt64,
		Fn: func(src interface{}) (interface{}, error) {
			s, ok := src.(time.Time)
			if !ok {
				return nil, errors.New("src type not matching")
			}
			p := s.UnixMilli()
			return &p, nil
		},
	}
	TimeToInt64 = TypeConverter{
		SrcType: Time,
		DstType: Int64,
		Fn: func(src interface{}) (interface{}, error) {
			s, ok := src.(time.Time)
			if !ok {
				return nil, errors.New("src type not matching")
			}
			return s.UnixMilli(), nil
		},
	}

	Int64ToTime = TypeConverter{
		SrcType: Int64,
		DstType: Time,
		Fn: func(src interface{}) (interface{}, error) {
			s, ok := src.(int64)
			if !ok {
				return nil, errors.New("src type not matching")
			}
			tm := time.UnixMilli(s)
			return tm, nil
		},
	}
	Int64ToPTime = TypeConverter{
		SrcType: Int64,
		DstType: PTime,
		Fn: func(src interface{}) (interface{}, error) {
			s, ok := src.(int64)
			if !ok {
				return nil, errors.New("src type not matching")
			}
			tm := time.UnixMilli(s)
			return &tm, nil
		},
	}

	PInt64ToTime = TypeConverter{
		SrcType: PInt64,
		DstType: Time,
		Fn: func(src interface{}) (interface{}, error) {
			s, ok := src.(*int64)
			if !ok {
				return nil, errors.New("src type not matching")
			}
			tm := time.UnixMilli(*s)
			return tm, nil
		},
	}
	PInt64ToPTime = TypeConverter{
		SrcType: PInt64,
		DstType: PTime,
		Fn: func(src interface{}) (interface{}, error) {
			s, ok := src.(*int64)
			if !ok {
				return nil, errors.New("src type not matching")
			}
			tm := time.UnixMilli(*s)
			return &tm, nil
		},
	}
)
