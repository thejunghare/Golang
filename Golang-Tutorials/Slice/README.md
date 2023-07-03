## Slice

Slices are similar to the arrays, but more powerful and flexible. They are used to store multiple values of the same type in a single variable.

_However, unlike arrays the length of a slice can grow and shrink as you see fit._

## Syntax for slice

```
slice_name := []int{} // create a slice

len(slice_name) // find length of slice

cap(slice_name) // find Capacity of slice

slice_name := make([]type, length, capacity) // create slice with make() function

slice_name[0] // access elements of slice

slice_name = append(slice_name, element1, element2, ...) // append elements to slice

slice_name = append(slice_name, slice_name...)
```
