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

// provides incrementing ids
type idProvider id

// returns a new id that is one bigger than the previous one
func (this *idProvider) NewId() id {
    i := id(*this)
    *this++
    return i
}