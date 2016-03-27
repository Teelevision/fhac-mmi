package graph

import "testing"

// tests that the id provider returns incrementing ids
func TestIdProvider(t *testing.T) {

    var id uint
    id = 5
    provider := idProvider(id)

    for ; id < 20; id++ {
        newId := provider.NewId()
        if id != newId.GetId() {
            t.Errorf("Expected new ID to be %d, got %d.", id, newId)
        }
    }
}
