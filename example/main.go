package main

import (
	"fmt"
	"github.com/sertangulveren/routarder"
	"github.com/sertangulveren/routarder/example/economy"
	"github.com/sertangulveren/routarder/example/science"
	"net/http"
)

func main() {
	r := routarder.New()
	r.HandleFunc("/", mainHandler)

	categoryGroup := r.Group("/kategori")

	mountEconomy(categoryGroup)
	mountScience(categoryGroup)
	r.HandleFunc("/p", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, r.RouteTree())
	})
	routarder.Serve(":3000")
}

func mountEconomy(r *routarder.Router) {
	economyGroup := r.Group("/ekonomi")

	economyGroup.HandleFunc("/", economy.MainHandler)

	economyGroup.HandleFunc("/enflasyon", economy.InflationHandler)
	economyGroup.HandleFunc("/asgari-ucret", economy.MinWageHandler)
}

func mountScience(r *routarder.Router) {
	scienceGroup := r.Group("/bilim")

	scienceGroup.HandleFunc("/", science.MainHandler)

	scienceGroup.HandleFunc("/yazilim", science.SoftwareHandler)
	scienceGroup.HandleFunc("/evrim", science.EvolutionHandler)
}

func mainHandler(w http.ResponseWriter, _ *http.Request)  {
	fmt.Fprint(w, "Kainatın Gazetesi!")
}