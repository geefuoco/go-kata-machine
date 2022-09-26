package util

func ServedBuildings(buildings, routerLocations, routerRanges []int) int {
  serveCounts := make(map[int]int)
  for i := range routerLocations {
    router := routerLocations[i]-1
    routerRange := routerRanges[i]

    for j := range buildings {
      if _, ok := serveCounts[j]; !ok {
        serveCounts[j]=0
      }
      if j >= router-routerRange && j <= router+routerRange {
        serveCounts[j]++
      }
    }
  }
  var totalServed int
  for i, v := range buildings {
    if serveCounts[i] >= v {
      totalServed++
    }
  }
  return totalServed
}
