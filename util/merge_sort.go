package util

func MergeSort(arr []int) {
  temp := make([]int, len(arr))
  merge(arr, 0, len(arr)-1, &temp)
}

func merge(arr []int, low, high int, temp *[]int) {
  if low >= high {
    return
  }
  middle := (low+high)/2
  merge(arr, low, middle, temp)
  merge(arr, middle+1, high, temp)
  mergeHalves(arr, low, high, temp)
}

func mergeHalves(arr []int, leftStart, rightEnd int, temp *[]int){
  leftEnd := (leftStart + rightEnd)/2
  rightStart := leftEnd+1
  size := rightEnd - leftStart+1

  left:= leftStart
  right := rightStart
  index := 0

  for left <= leftEnd && right <= rightEnd {
    if arr[left] <= arr[right] {
      (*temp)[index] = arr[left]
      left++
    } else {
      (*temp)[index] = arr[right]
      right++
    }
    index++
  }

  for left <= leftEnd {
    (*temp)[index] = arr[left]
    left++
    index++
  }

  for right <= rightEnd{
    (*temp)[index] = arr[right]
    right++
    index++
  }

  for i:=0; i < size; i++ {
    arr[leftStart +i] = (*temp)[i]
  }
}
