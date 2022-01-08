package routarder

type Router struct {
	prefix   string
	routes   []string
	children []*Router
}
