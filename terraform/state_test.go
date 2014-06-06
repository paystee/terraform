package terraform

import (
	"reflect"
	"testing"
)

func TestResourceState_MergeDiff(t *testing.T) {
	rs := ResourceState{
		ID: "foo",
		Attributes: map[string]string{
			"foo": "bar",
		},
	}

	diff := map[string]*ResourceAttrDiff{
		"foo": &ResourceAttrDiff{
			Old: "bar",
			New: "baz",
		},
		"bar": &ResourceAttrDiff{
			Old: "",
			New: "foo",
		},
		"baz": &ResourceAttrDiff{
			Old:         "",
			New:         "foo",
			NewComputed: true,
		},
	}

	rs2 := rs.MergeDiff(diff, "what")

	expected := map[string]string{
		"foo": "baz",
		"bar": "foo",
		"baz": "what",
	}

	if !reflect.DeepEqual(expected, rs2.Attributes) {
		t.Fatalf("bad: %#v", rs2.Attributes)
	}
}

func TestResourceState_MergeDiff_nil(t *testing.T) {
	var rs *ResourceState = nil

	diff := map[string]*ResourceAttrDiff{
		"foo": &ResourceAttrDiff{
			Old: "",
			New: "baz",
		},
	}

	rs2 := rs.MergeDiff(diff, "computed")

	expected := map[string]string{
		"foo": "baz",
	}

	if !reflect.DeepEqual(expected, rs2.Attributes) {
		t.Fatalf("bad: %#v", rs2.Attributes)
	}
}