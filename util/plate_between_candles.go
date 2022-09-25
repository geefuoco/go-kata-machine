package util

func NumberOfItems(input string, startIndices []int, endIndices []int) []int {
  k := len(input)
  memo := make([]int, k)
  var count int
  for i, c := range input {
    if c == '|' {
      memo[i] = count
    } else {
      count++
    }
  }
  result := []int{}
  for i := range startIndices {
    start := startIndices[i]-1
    end := endIndices[i]-1

    for input[start] != '|' {
      start++
    }
    for input[end] != '|' {
      end--
    }
    if start < end {
      result = append(result ,memo[end] - memo[start])
    } else {
      result = append(result, 0)
    }
  }
  return result
}
