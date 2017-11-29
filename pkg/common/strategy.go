package common

type Strategy interface {
	Serve(h string, p string)
	RegisterBackend(h string, p string)
}

type StrategyMap map[Protocol]Strategy
