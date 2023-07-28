package internal

type SupportLevel int

const (
	Basic SupportLevel = iota
	Advanced
	Expert
)

type SupportRequest struct {
	Level   SupportLevel
	Message string
}
