package common

type Servable interface {
	Use(p Protocol)
	Serve(h string, p string)
	RegisterBackend(h string, p string)
}

type ServeBase struct {
	currentStrategy Strategy
	strategies      StrategyMap
}

func (s *ServeBase) GetCurrent() Strategy {
	return s.currentStrategy
}

func (s *ServeBase) SetCurrent(protocol Protocol) {
	s.currentStrategy = s.strategies[protocol]
}

func (s *ServeBase) SetStrategies(stmap StrategyMap) {
	s.strategies = stmap
}
