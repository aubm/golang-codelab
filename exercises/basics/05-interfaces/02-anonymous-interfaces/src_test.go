package basics_test

import (
	"reflect"
	"testing"

	. "github.com/aubm/golang-codelab/exercises/basics/05-interfaces/02-anonymous-interfaces"
)

func TestIntroducerIsAnAnonymousInterface(t *testing.T) {
	meeting := Meeting{}
	typeOfMeeting := reflect.TypeOf(meeting)
	introducerField, fieldExists := typeOfMeeting.FieldByName("Introducer")
	if !fieldExists {
		t.Error(`Type Meeting has no field named Introducer.`)
		return
	}

	if introducerField.Type.Kind() != reflect.Interface {
		t.Error(`Field Introducer must be an interface.`)
		return
	}

	if introducerField.Type.Name() != "" {
		t.Error(`Field Introducer must have an unnamed type.`)
		return
	}

	meeting.Introducer = Capuchin{}
	meeting.Introducer = Baboon{}
	meeting.Introducer = Mandrill{}
}
