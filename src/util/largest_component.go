package util

func LargestComponent(graph Graph) int {
  visited := Set{}
  var maxCount int
  for i := range graph {
    var count int
    graphSearch(graph, i, visited, &count)
    if count > maxCount {
      maxCount =count 
    }
  }
  return maxCount
}

func graphSearch(graph Graph, node rune, visited Set, count *int) {
  if _, ok := visited[node]; ok {
    return 
  }
  visited[node] = true
  *count++
  for _, neighbor := range graph[node] {
    graphSearch(graph, neighbor, visited, count)
  }
}
