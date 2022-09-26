package wrecon

type Site struct {
	Name    string
	Address string
}

func (s Site) GetName() string {
	name := s.Address
	if s.Name != "" {
		name = s.Name
	}
	return name
}
