package flyweight

/*Factory is struct*/
type Factory struct {
	soldierMap map[string]ISoldier
}

/*NewFactory is function*/
func NewFactory() *Factory {
	return &Factory{

		soldierMap: make(map[string]ISoldier),
	}
}

/*CreateSoldier is function*/
func (f *Factory) CreateSoldier(name string) ISoldier {
	if val, ok := f.soldierMap[name]; ok {
		return val
	}

	f.soldierMap[name] = newSoldier(name)
	return f.soldierMap[name]
}

/*GetSize is function*/
func (f *Factory) GetSize() int {
	return len(f.soldierMap)
}
