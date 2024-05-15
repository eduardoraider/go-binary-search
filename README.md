# Linear Search vs Binary Search in Go

Linear search and binary search are two commonly used algorithms for searching elements in an array. Here's a brief explanation of each:

## Linear Search
- **Method**: Sequentially scans each element of the array until the desired element is found or the end of the array is reached.
- **Performance**: O(n), where 'n' is the number of elements in the array. This means that the execution time increases linearly with the number of elements in the array. In terms of efficiency, the larger the array, the longer it will take to find the desired element, in the worst case scenario.


## Binary Search
- **Method**: Repeatedly divides the array in half and checks if the desired element is in the left or right half. This process continues until the element is found or the array part becomes empty.
- **Performance**: O(log n), where 'n' is the number of elements in the array. This means that the execution time increases logarithmically with the number of elements in the array. Binary search is significantly faster than linear search, especially for large arrays, because it discards half of the remaining elements in each iteration.

## Understanding Big O Notation
- **Big O notation**: It is a mathematical notation that describes the limiting behavior of a function when the argument approaches infinity. In the context of algorithms, it represents the worst-case time complexity or the upper bound of the time complexity of an algorithm. For example, O(n) indicates that the algorithm's performance grows linearly as the input size increases, while O(log n) signifies a logarithmic growth rate, indicating a more efficient algorithm compared to linear time complexity.

In summary, binary search has much better performance than linear search, especially for large arrays. However, binary search requires the array to be sorted, which can be a disadvantage if sorting is not an operation you already need to do anyway. If the array is not sorted, linear search may be the only viable option.


## How to Run

```
go run main.go
```

The program will output the results of both linear search and binary search on a sample array. You can modify the `main.go` file to test with different arrays or target values.


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

by **Eduardo Raider** - Software Engineer