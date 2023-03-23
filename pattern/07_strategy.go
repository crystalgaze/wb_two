package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

type Strategy interface {
	Route(startPoint int, endPoint int)
}

type RouteStrategy struct {
}

func (r *RouteStrategy) Route(startPoint int, endPoint int) {
	avgSpeed := 30
	trafficJam := 2
	total := endPoint - startPoint
	totalTime := total * 40 * trafficJam
	fmt.Printf("Road A:[%d] to B: [%d] Avg speed: [%d] Traffic jam: [%d] Total: [%d] Total time: [%d] min\n",
		startPoint, endPoint, avgSpeed, trafficJam, total, totalTime)
}

type PublicTransportStrategy struct {
}

func (p *PublicTransportStrategy) Route(startPoint int, endPoint int) {
	avgSpeed := 40
	total := endPoint - startPoint
	totalTime := total * 40
	fmt.Printf("PublicTransport A:[%d] to B: [%d] Avg speed: [%d] Total: [%d] Total time: [%d] min\n",
		startPoint, endPoint, avgSpeed, total, totalTime)
}

type WalkStrategy struct {
}

func (p *WalkStrategy) Route(startPoint int, endPoint int) {
	avgSpeed := 4
	total := endPoint - startPoint
	totalTime := total * 60
	fmt.Printf("Walk A:[%d] to B: [%d] Avg speed: [%d] Total: [%d] Total time: [%d] min\n",
		startPoint, endPoint, avgSpeed, total, totalTime)
}

type Navigator struct {
	Strategy
}

func (n *Navigator) SetStrategy(str Strategy) {
	n.Strategy = str
}

var (
	start      = 10
	end        = 100
	strategies = []Strategy{
		&PublicTransportStrategy{},
		&RouteStrategy{},
		&WalkStrategy{},
	}
)

func main() {
	// Контекст
	nav := Navigator{}
	for _, strategy := range strategies {
		nav.SetStrategy(strategy)
		nav.Route(start, end)
	}
}
