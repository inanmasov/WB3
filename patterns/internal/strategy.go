package internal

import "fmt"

type FindShortestPath interface {
	FindShortPath(graphFileName string)
}

// Алгоритм Дейкстры
type AlgoritmDijkstra struct{}

// Реализация алгоритма Дейкстры
func (dij *AlgoritmDijkstra) FindShortPath(graphFileName string) {
	fmt.Println("Algoritm Dijkstra worked")
}

// Алгоритм Беллмана-Форда
type AlgoritmBellmanFord struct{}

// Реализация алгоритма Беллмана-Форда
func (bf *AlgoritmBellmanFord) FindShortPath(graphFileName string) {
	fmt.Println("Algoritm Bellman-Ford worked")
}

// Граф
type Graph struct {
	find FindShortestPath
}

// Устнавливаем нужный алгоритм
func (g *Graph) SetAlgoritm(find FindShortestPath) {
	g.find = find
}

// Выполняем алгоритм
func (g *Graph) FindPath(graphFileName string) {
	g.find.FindShortPath(graphFileName)
}

func StrategyPattern() {
	// Создаем граф с алгоритмом Дейкстры
	graph := &Graph{find: &AlgoritmDijkstra{}}
	// Ищем кратчайщий путь алгоритмом Дейкстры
	graph.FindPath("graph1.txt")
	// Устанавливаем алгоритм Беллмана-Форда
	graph.SetAlgoritm(&AlgoritmBellmanFord{})
	// Ищем кратчайщий путь алгоритмом Беллмана-Форда
	graph.FindPath("graph2.txt")
}
