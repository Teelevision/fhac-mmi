package graph

// interface for objects that expose and ID
type idInterface interface {
    GetId() uint
}

// alias for an ID, which is just an unsigned integer
type id uint

// returns the id as unsigned integer
func (this id) GetId() uint {
    return uint(this)
}

// provides unique ids
type idProvider uint

// returns a new id that was not previously used
func (this idProvider) NewId() uint {
    id := uint(this)
    this++
    return id
}