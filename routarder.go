package routarder

type Router struct {
	prefix   string
	routes   []string
	children []*Router
}

func New() *Router {
	return &Router{prefix: "/"}
}

func Serve(addr string) {
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}