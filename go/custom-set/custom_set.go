package stringset

import (
	"strconv"
	"strings"
)

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

func (s Set) String() string {
    var sb strings.Builder
    count := 0
    sb.WriteString("{")
    for k := range s.elements {
        quotedString := strconv.Quote(k)
        count++
        if count == 1{
            sb.WriteString(quotedString)
        } else if count > 1 {
            sb.WriteString(", " + quotedString)
        }     
    }
    sb.WriteString("}")
    return sb.String()
}

func (s Set) IsEmpty() bool {
    return len(s.elements) == 0
}

func (s Set) Has(elem string) bool {
    return s.elements[elem] == true
}

func (s Set) Add(elem string) {
    if !s.Has(elem) {
        s.elements[elem] = true
    }
}

func Subset(s1, s2 Set) bool {
    for k := range s1.elements {
        if !s2.Has(k) {
            return false
        }
    }
    return true
}

func Disjoint(s1, s2 Set) bool {
    for k := range s1.elements {
        if s2.Has(k) {
            return false
        }
    }
    return true
}

func Equal(s1, s2 Set) bool {
    if len(s1.elements) != len(s2.elements) {
        return false
    }
    for k := range s1.elements {
        if !s2.Has(k) {
            return false
        }
    }
    return true
}

func Intersection(s1, s2 Set) Set {
    s3 := New()
    for k := range s1.elements {
        if s2.Has(k) {
            s3.elements[k] = true
        }
    }
    return s3
}

func Difference(s1, s2 Set) Set {
    s3 := New()
    for k := range s1.elements {
        if !s2.Has(k) {
            s3.elements[k] = true
        }
    }
    return s3
}

func Union(s1, s2 Set) Set {
    s3 := New()
    for k := range s1.elements {
        s3.elements[k] = true
    }
    for k := range s2.elements {
        s3.elements[k] = true
    }
    return s3
}
