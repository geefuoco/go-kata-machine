package util

func QuickSort(arr []int) {
  quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, low, high int) {
  if low >= high {
    return
  }
  pivot := arr[(low+high)/2]
  index := partition(arr, low, high, pivot)
  quickSort(arr, low, index-1)
  quickSort(arr, index, high)
}

func partition(arr []int, low, high, pivot int) int {
  for low <= high {
    for arr[low] < pivot {
      low++
    }
    for arr[high] > pivot {
      high--
    }
    if low <= high {
      swap(arr, low, high)
      low++
      high--
    }
  }
  return low
}

func swap(arr []int, left, right int) {
  temp := arr[left]
  arr[left] = arr[right]
  arr[right] = temp
}
