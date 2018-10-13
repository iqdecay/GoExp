// treesort uses a binary tree to sort int slices in-place
package main

import (
    "fmt"
    )

type tree struct{
    value int
    left, right *tree
    }

// Sort sorts values in place
func Sort(values []int) {
    var root *tree
    for _, v := range values {
        root = add(root, v)
    }
    appendValues(values[:0], root)
}

// appendValues append the elements of t to values
// in order and returns the resulting slice
func appendValues(values []int, t *tree) []int {
   if t != nil {
        values = appendValues(values, t.left)
        values = append(values, t.value)
        values = appendValues(values, t.right)
   }
   return values
}

func add(t *tree, value int) *tree {
    if t == nil {
        // Equivalent to return &tree{value: value}
        t = new(tree)
        t.value = value
        return t
   }
   if value < t.value {
    t.left = add(t.left, value)
   } else {
    t.right = add(t.right, value)
   }
   return t
}

func main() {
    a := []int{10,9,7,5,6}
    fmt.Println(a)
    Sort(a)
    fmt.Println(a)
    }
