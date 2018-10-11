package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/barnex/chess"
)

var flagAddr = flag.String("http", "localhost:26355", "HTTP address")

func main() {
	http.Handle("/play", http.StripPrefix("/play", http.HandlerFunc(handlePlay)))
	fmt.Println("visit http://" + *flagAddr + "/play")
	log.Fatal(http.ListenAndServe(*flagAddr, nil))
}

func handlePlay(w http.ResponseWriter, r *http.Request) {

	d := &data{
		I:     []int{7, 6, 5, 4, 3, 2, 1, 0},
		J:     []int{0, 1, 2, 3, 4, 5, 6, 7},
		Board: chess.NewBoard(),
	}

	err := templ.Execute(w, d)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type data struct {
	I     []int
	J     []int
	Board *chess.Board
}

func (d *data) At(i, j int) chess.Piece {
	return d.Board.At(chess.RC(i, j))
}

func (d *data) Oddness(i, j int) string {
	if (i+j)%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}

var templ = template.Must(template.New("play").Parse(templText))

const templText = `
<html>
<head>
	<style>
		.odd{
			background-color: #759455;	
		}	
		.even{
			background-color: #efefd3;	
		}	
		td{
			width: 2em;	
			height: 2em;	
			font-size: 2em;
			text-align: center;
		}
	</style>
</head>
<body>

<table>
{{range $x,$i := $.I}}
	<tr>
	{{range $x,$j := $.J}}
	<td class="{{$.Oddness $i $j}}"> {{$.At $i $j}} </td>
	{{end}}
	</tr>
{{end}}
</table>

</body>
</html>
`
