package pagination

type OrderDir int32

type Pagination struct {
	Limit    int
	Offset   int
	OrderBy  string
	OrderDir OrderDir
}

const (
	ASC  OrderDir = 1
	DESC OrderDir = 2
)

func (p OrderDir) String() string {
	switch p {
	case ASC:
		return "ASC"
	case DESC:
		return "DESC"
	default:
		return "ASC"
	}
}
