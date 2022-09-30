package util

func ShortestPath(edges [][2]rune, start rune, end rune) int{
  visited := Set{}
  visited[start] = true
  shortest := int(^uint(0)>>1)
  graph := ConvertEdgelistToGraph(edges)
  for _, node := range graph[start] {
    var count int
    exploreEdges(graph, node, end, visited, &count)
    if count < shortest{
      shortest = count
    }
  }
  return shortest 
}

func exploreEdges(graph Graph, node, end rune, visited Set, count *int) {
  if _, ok := visited[node]; ok{
    return
  }
  if(node == end){
    *count++
    return
  }
  visited[node] =true
  *count++
  for _, neighbor := range graph[node] {
    exploreEdges(graph, neighbor, end, visited, count)
  }
}

func ConvertEdgelistToGraph(edgelist [][2]rune) Graph {
  graph := make(map[rune][]rune)
  for _, edges := range edgelist {
    a, b := edges[0], edges[1]
    if _, ok := graph[a]; !ok {
      graph[a] = []rune{}
    }
    if _, ok := graph[b]; !ok {
      graph[b] = []rune{}
    }
    graph[a] = append(graph[a], b)
    graph[b] = append(graph[b], a)
  }

  return graph
}
