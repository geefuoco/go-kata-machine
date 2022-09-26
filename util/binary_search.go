package util


func BinarySearch(array []int, value int) bool {
  if(len(array) == 0){
    return false
  }
  middle := len(array)/2
  if value < array[middle] {
    return BinarySearch(array[:middle], value)
  } else if value > array[middle] {
    return BinarySearch(array[middle+1:], value)
  } else if value == array[middle]{
    return true
  }
  return false
}
