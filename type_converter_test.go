package copier_test

import (
	"github.com/jinzhu/copier"
	"testing"
	"time"
)

func TestConvertersTime(t *testing.T) {
	type SrcStruct struct {
		Field1 time.Time
		Field2 *time.Time
		Field3 time.Time
		Field4 *time.Time
	}

	type DestStruct struct {
		Field1 int64
		Field2 int64
		Field3 *int64
		Field4 *int64
	}

	t1 := time.Now()
	t2 := t1.Add(time.Minute)
	t3 := t2.Add(time.Minute)
	t4 := t3.Add(time.Minute)
	src := SrcStruct{
		Field1: t1,
		Field2: &t2,
		Field3: t3,
		Field4: &t4,
	}
	var dst DestStruct
	err := copier.Copy(&dst, &src)
	if err != nil {
		t.Error("copyError", err)
		return
	}

	if v := src.Field1.UnixMilli(); v != dst.Field1 {
		t.Fatalf("got %q, wanted %q", v, dst.Field1)
	}

	if v := src.Field2.UnixMilli(); v != dst.Field2 {
		t.Fatalf("got %q, wanted %q", v, dst.Field2)
	}
	if v := src.Field3.UnixMilli(); v != *dst.Field3 {
		t.Fatalf("got %q, wanted %q", v, *dst.Field3)
	}

	if v := src.Field4.UnixMilli(); v != *dst.Field4 {
		t.Fatalf("got %q, wanted %q", v, *dst.Field4)
	}

	newSrc := SrcStruct{}

	err = copier.Copy(&newSrc, &dst)
	if err != nil {
		t.Error("copyError", err)
		return
	}
	if v := dst.Field1; v != newSrc.Field1.UnixMilli() {
		t.Fatalf("got %q, wanted %q", v, newSrc.Field1.UnixMilli())
	}

	if v := dst.Field2; v != newSrc.Field2.UnixMilli() {
		t.Fatalf("got %q, wanted %q", v, newSrc.Field2)
	}
	if v := dst.Field3; *v != newSrc.Field3.UnixMilli() {
		t.Fatalf("got %q, wanted %q", *v, newSrc.Field3.UnixMilli())
	}

	if v := dst.Field4; *v != newSrc.Field4.UnixMilli() {
		t.Fatalf("got %q, wanted %q", *v, newSrc.Field4.UnixMilli())
	}
}
