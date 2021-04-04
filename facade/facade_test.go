package facade

import (
	"testing"
)

func TestFacade(t *testing.T) {
	facadeImp := createFacade()
	appleContent := facadeImp.getAppleAttribute()
	bananaContent := facadeImp.getBananaAttribute()
	if appleContent != "name: apple,price: 1" {
		t.Error("apple content error")
	}
	if bananaContent != "name: banana,price: 2" {
		t.Error("banana content error")
	}
}
