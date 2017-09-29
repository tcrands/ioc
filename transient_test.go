package main

import "testing"

type testStruct struct {
	val int
}

func TestTransientContainer_Register_item_does_not_exist(t *testing.T) {

	container := newTransientContainer()

	testVal := &testStruct{val: 12}
	container.Register("new", testVal)

	if !container.Exists("new") {
		t.Fatalf("Expected true, got false")
	}
}

func TestTransientContainer_Register_item_does_exist(t *testing.T) {

	container := newTransientContainer()

	testVal := &testStruct{val: 12}
	container.Mapper["new"] = testVal
	val := container.Register("new", testVal)
	expected := "'new' is already in IoC container."

	if val.Error() != expected {
		t.Fatalf("Expected %s, got %s", expected, val.Error())
	}
}

func TestTransientContainer_resolve(t *testing.T) {

	container := newTransientContainer()

	testVal := &testStruct{val: 12}
	container.Mapper["new"] = testVal
	val, _ := container.Resolve("new")

	if val != testVal {
		t.Fatalf("Expected")
	}
}

func TestTransientContainer_release(t *testing.T) {

	container := newTransientContainer()

	testVal := &testStruct{val: 12}
	container.Mapper["new"] = testVal
	container.Release("new")

	if container.Exists("new") {
		t.Fatalf("Expected false got true")
	}
}

func TestTransientContainer_exists(t *testing.T) {

	container := newTransientContainer()

	testVal := &testStruct{val: 12}
	container.Mapper["new"] = testVal
	val := container.Exists("new")

	if val != true {
		t.Fatalf("Expected false got true")
	}
}

func TestTransientContainer_exists_not(t *testing.T) {

	container := newTransientContainer()

	val := container.Exists("new")

	if val == true {
		t.Fatalf("Expected false got true")
	}
}
