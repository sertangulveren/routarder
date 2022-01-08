package science

import (
	"fmt"
	"net/http"
)

func MainHandler(w http.ResponseWriter, _ *http.Request)  {
	fmt.Fprint(w, "Kuantum iyi ama anlatması zor.\n- Albert Einstein")
}

func SoftwareHandler(w http.ResponseWriter, _ *http.Request)  {
	fmt.Fprint(w, "Rust'ın bu kadar yayılacağını bilmiyordum. Valla ben de çok şaşırdım.\n- Graydon Hoare")
}

func EvolutionHandler(w http.ResponseWriter, _ *http.Request)  {
	fmt.Fprint(w, "Tavuklar T-Rex'ten gelmedi. Ataları ortak.\n- Çağrı Mert Bakırcı")
}


