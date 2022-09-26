package wrecon

import "testing"

func TestSite_GetName(t *testing.T) {
	address := "www.example.com"
	site := Site{Address: address}
	if site.GetName() != address {
		t.Errorf("Site GetName should return %s, got %s", address, site.GetName())
	}

	name := "Example"
	otherSite := Site{Address: address, Name: name}
	if otherSite.GetName() != name {
		t.Errorf("Site GetName should return %s, got %s", name, otherSite.GetName())
	}
}
