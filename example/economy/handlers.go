package economy

import (
	"fmt"
	"net/http"
)

func MainHandler(w http.ResponseWriter, _ *http.Request)  {
	fmt.Fprint(w, "Borç yiğidin değil yoksulun kamçısıdır.\n- David Ricardo")
}

func InflationHandler(w http.ResponseWriter, _ *http.Request)  {
	fmt.Fprint(w, "Bu enflasyon önümüzü görmeyi engelliyor. Herkes bir değer söylüyor. Bence %51\n- Robert Solow")
}

func MinWageHandler(w http.ResponseWriter, _ *http.Request)  {
	fmt.Fprint(w, "Kuantum iyi ama anlatması zor.\n- Albert Einstein")
}
