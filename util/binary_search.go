package util


func BinarySearch(array []int, value int) bool {
  middle := len(array)/2
  if value < array[middle] {
    return BinarySearch(array[:middle], value)
  } else if value > array[middle] {
    return BinarySearch(array[middle:], value)
  } else if value == array[middle]{
    return true
  }
  return false
}
