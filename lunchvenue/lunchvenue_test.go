package lunchvenue_test

import "testing"
import "github.com/treychua/beatricethetelegrambot/lunchvenue"

import "strconv"

func TestAdd(t *testing.T) {
	lvs := lunchvenue.LunchVenues{}
	err := lvs.Add("location1")

	if 1 != len(lvs) {
		t.Errorf("Expected length of LunchVenues to be %v but got %v", 1, len(lvs))
	}

	if "location1" != lvs[0].Location {
		t.Errorf("Expected %v but got %v", "location1", lvs[0].Location)
	}

	err = lvs.Add("location1")
	if err == nil {
		t.Errorf("Expected an error but got nil instead")
	}

	err = lvs.Add("location2")
	if 2 != len(lvs) {
		t.Errorf("Expected length of LunchVenues to be %v but got %v", 2, len(lvs))
	}
}

func TestHas(t *testing.T) {
	lvs := lunchvenue.LunchVenues{}

	if lvs.Has("l1") {
		t.Errorf("Expected no venue found but received true")
	}

	lvs.Add("l1")

	if false == lvs.Has("l1") {
		t.Error("Expected venue to be fonud but received false")
	}

}

func TestDelete(t *testing.T) {
	lvs := lunchvenue.LunchVenues{}

	for i := 1; i <= 5; i++ {
		lvs.Add("lunch venue " + strconv.Itoa(i))
	}

	result, err := lvs.Delete("lunch venue 1")
	if !result || nil != err {
		t.Errorf("Expected %v, %v but got %v, %v instead", true, nil, result, err)
	}

	if lvs.Has("lunch venue 1") {
		t.Error("Expected lunch venue 1 to be removed but it is still inside")
	}

	result, err = lvs.Delete("lunch venue 3")
	if !result || nil != err {
		t.Errorf("Expected %v, %v but got %v, %v instead", true, nil, result, err)
	}

	if lvs.Has("lunch venue 3") {
		t.Error("Expected lunch venue 3 to be removed but it is still inside")
	}

	result, err = lvs.Delete("lunch venue 1")
	if result || nil == err {
		t.Errorf("Expected %v, %v but got %v, %v instead", true, nil, result, err)
	}

}
