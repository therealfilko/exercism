package stringset

import "strings"

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set struct {
    elements map[string]bool
}

func New() Set {
    return Set{elements: make(map[string]bool)}
}

func NewFromSlice(l []string) Set {
    mySet := New()
    for _, v := range l {
        mySet.elements[v] = true
    }
    return mySet
}

// muss schauen wie wir das machen
// string muss { elment1, element2 } ausgeben
// vielleicht len - 1 um die länge zu haben minus eins und so oft durch gehen 
// mit einer schleife und letzte ist dann halt ohne komma
func (s Set) String() string {
    var builder strings.Builder
    builder.WriteString("{")
    for k := range s.elements {
        builder.WriteString(k)
    }
    builder.WriteString("}")
    return builder.String()
}

func (s Set) IsEmpty() bool {
    if len(s.elements) == 0 {
        return true
    }
    return false
}

func (s Set) Has(elem string) bool {
	panic("Please implement the Has function")
}

func (s Set) Add(elem string) {
	panic("Please implement the Add function")
}

func Subset(s1, s2 Set) bool {
	panic("Please implement the Subset function")
}

func Disjoint(s1, s2 Set) bool {
	panic("Please implement the Disjoint function")
}

func Equal(s1, s2 Set) bool {
	panic("Please implement the Equal function")
}

func Intersection(s1, s2 Set) Set {
	panic("Please implement the Intersection function")
}

func Difference(s1, s2 Set) Set {
	panic("Please implement the Difference function")
}

func Union(s1, s2 Set) Set {
	panic("Please implement the Union function")
}
