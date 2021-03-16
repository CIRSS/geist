package geist

type Client interface {
	Select(query string) (rs *ResultSet, err error)
}
