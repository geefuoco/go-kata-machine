package util

type Set map[rune]bool
type Graph map[rune][]rune

func HasCycle(graph Graph) bool {
  visited := Set{}
  currentStack := Set{}
  for node := range graph {
    if(search(graph, node, &visited, &currentStack)){
      return true
    }
  }
  return false
}

func search(graph Graph, node rune, visited, stack *Set) bool {
  if _, has := (*visited)[node]; has {
    if _, ok := (*stack)[node]; ok {
      return true
    }
  }
  (*visited)[node] = true
  for _, neighbor := range graph[node] {
    (*stack)[node] = true
    if(search(graph, neighbor, visited, stack)) {
      return true
    }
    delete(*stack, node)
  }
  delete(*visited, node)
  return false
}

