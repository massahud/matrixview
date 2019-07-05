# matrixview

matrixview is a library to create simple views of slice matrixes.

 
The intention of this package is to reuse a single big matrix without needing to 
allocate more matrixes when you need different matrix sizes. 
e.g. calculate edit distance of lots of pairs of strings   

## Example

```go
func ExampleEditDistance() {
	buf := make([][]int, 100)
	for i := range buf {
		buf[i] = make([]int, 100)
	}


	work := [][2]string{
		{"sitting", "kitten"},
		{"Sunday", "Saturday"},
	}

	for _, words := range work {
		from := words[0]
		to := words[1]

		view := FromInt(buf, 0, 0, len(from)+1, len(to)+1)
		if err := fillLevenshteinMatrix(view, from, to); err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			return
		}
		fmt.Printf("From: %s To: %s Distance: %d Ratio: %.2f\n", from, to, calculateDistance(view), calculateRatio(view))
	}

	// Output:
	//From: sitting To: kitten Distance: 5 Ratio: 0.62
	//From: Sunday To: Saturday Distance: 4 Ratio: 0.71
}

```

 