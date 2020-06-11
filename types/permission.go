package types

const (
	ReadPerm PermValue = iota
	TransferPerm
)

type PermValue int32

func (p PermValue) IsValid() bool {
	if p != ReadPerm && p != TransferPerm {
		return false
	}
	return true
}
