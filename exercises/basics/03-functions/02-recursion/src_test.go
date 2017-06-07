package basics_test

import (
	"encoding/json"
	"testing"

	. "github.com/aubm/golang-codelab/exercises/basics/03-functions/02-recursion"
)

var family = `{
   "Name":"Socorro",
   "Children":[
      {
         "Name":"Michael",
         "Children":[
            { "Name":"Cynthia", "Children":[] }
         ]
      },
      {
         "Name":"Darrin",
         "Children":[
            {
               "Name":"Carroll",
               "Children":[
                  {
                     "Name":"Rodney",
                     "Children":[
                        { "Name":"Sandra", "Children":[] },
                        { "Name":"Jesus", "Children":[] },
                        { "Name":"Lashawn", "Children":[] }
                     ]
                  },
                  {
                     "Name":"Jason",
                     "Children":[
                        { "Name":"Margie", "Children":[] },
                        {
                           "Name":"Doreen",
                           "Children":[
                              { "Name":"Frank", "Children":[] }
                           ]
                        }
                     ]
                  }
               ]
            }
         ]
      }
   ]
}`

func TestNumberOfDescendants(t *testing.T) {
	tree := MemberOfTheFamily{}
	json.Unmarshal([]byte(family), &tree)
	expected := 12
	if actual := NumberOfDescendants(tree); actual != expected {
		t.Errorf("when family is: %v\nexpected number of descendants to equal %v, but got %v", family, expected, actual)
	}
}
