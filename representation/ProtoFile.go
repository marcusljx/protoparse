package representation

type ProtoFile struct {
	Syntax string
	Package string
	Options Options
	Services []*Service
}