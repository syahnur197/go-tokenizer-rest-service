package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/consumer-one", consumerOne).Methods("GET")
	r.HandleFunc("/consumer-two", consumerTwo).Methods("GET")
	r.HandleFunc("/consumer-three", consumerThree).Methods("GET")

	log.Fatal(http.ListenAndServe(":8001", r))
}

func consumerOne(w http.ResponseWriter, r *http.Request) {
	text := "You can only define methods on a type defined in that same package. Your DB type, in this case, is defined within your dbconfig package, so your entity package can't define methods on it.In this case, your options are to make GetContracts a function instead of a method and hand it the *dbconfig.DB as an argument, or to invert the dependency by importing your entity package in dbconfig and write GetContracts there (as a method or function, works either way). This second one may actually be the better option, because, from a design perspective, it breaks abstraction to have packages other than your database package creating SQL query strings."

	consumer(w, r, text)
}

func consumerTwo(w http.ResponseWriter, r *http.Request) {
	text := "Brunei, officially the Nation of Brunei, the Abode of Peace[10] (Malay: Negara Brunei Darussalam) (also Brunei Darussalam) is a country located on the north coast of the island of Borneo in Southeast Asia. Apart from its South China Sea coast, it is completely surrounded by the Malaysian state of Sarawak. It is separated into two parts by the Sarawak district of Limbang. Brunei is the only sovereign state entirely on Borneo; the remainder of the island is divided between Malaysia and Indonesia. As of 2020, its population was 460,345,[5] of whom about 100,000 live in the capital and largest city, Bandar Seri Begawan. The government is an absolute monarchy ruled by its Sultan, entitled the Yang di-Pertuan, and implements a combination of English common law and sharia law, as well as general Islamic practices.	At the peak of the Bruneian Empire, Sultan Bolkiah (reigned 1485-1528) is claimed to have had control over most regions of Borneo, including modern-day Sarawak and Sabah, as well as the Sulu Archipelago off the northeast tip of Borneo, and the islands off the northwest tip of Borneo. Claims also state that they had control over Seludong (or the Kingdom of Maynila, where the modern-day Philippine capital Manila now stands) but Southeast Asian scholars believe this refers to a settlement Mount Selurong in Indonesia.[11] The maritime state of Brunei was visited by Spain's Magellan Expedition in 1521 and fought against Spain in the 1578 Castilian War. During the 19th century, the Bruneian Empire began to decline. The Sultanate ceded Sarawak (Kuching) to James Brooke and installed him as the White Rajah, and it ceded Sabah to the British North Borneo Chartered Company. In 1888, Brunei became a British protectorate and was assigned a British resident as colonial manager in 1906. After the Japanese occupation during World War II, in 1959 a new constitution was written. In 1962, a small armed rebellion against the monarchy was ended with the help of the British.Brunei gained its independence from the United Kingdom on 1 January 1984. Economic growth during the 1990s and 2000s, with the GDP increasing 56% from 1999 to 2008, transformed Brunei into an industrialised country. It has developed wealth from extensive petroleum and natural gas fields. Brunei has the second-highest Human Development Index among the Southeast Asian nations, after Singapore, and is classified as a developed country. According to the International Monetary Fund (IMF), Brunei is ranked fifth in the world by gross domestic product per capita at purchasing power parity. The IMF estimated in 2011 that Brunei was one of two countries (the other being Libya) with a public debt at 0% of the national GDP."

	consumer(w, r, text)
}

func consumerThree(w http.ResponseWriter, r *http.Request) {
	text := "The Go language has built-in facilities, as well as library support, for writing concurrent programs. Concurrency refers not only to CPU parallelism, but also to asynchrony: letting slow operations like a database or network read run while the program does other work, as is common in event-based servers. The primary concurrency construct is the goroutine, a type of light-weight process. A function call prefixed with the go keyword starts a function in a new goroutine. The language specification does not specify how goroutines should be implemented, but current implementations multiplex a Go process's goroutines onto a smaller set of operating-system threads, similar to the scheduling performed in Erlang While a standard library package featuring most of the classical concurrency control structures (mutex locks, etc.) is available, idiomatic concurrent programs instead prefer channels, which send messages between goroutines.[78] Optional buffers store messages in FIFO orderand allow sending goroutines to proceed before their messages are received. Channels are typed, so that a channel of type chan T can only be used to transfer messages of type T. Special syntax is used to operate on them; <-ch is an expression that causes the executing goroutine to block until a value comes in over the channel ch, while ch <- x sends the value x (possibly blocking until another goroutine receives the value). The built-in switch-like select statement can be used to implement non-blocking communication on multiple channels; see below for an example. Go has a memory model describing how goroutines must use channels or other operations to safely share data. The existence of channels sets Go apart from actor model-style concurrent languages like Erlang, where messages are addressed directly to actors (corresponding to goroutines). The actor style can be simulated in Go by maintaining a one-to-one correspondence between goroutines and channels, but the language allows multiple goroutines to share a channel or a single goroutine to send and receive on multiple channels.From these tools one can build concurrent constructs like worker pools, pipelines (in which, say, a file is decompressed and parsed as it downloads), background calls with timeout, 'fan-out' parallel calls to a set of services, and others.[81] Channels have also found uses further from the usual notion of interprocess communication, like serving as a concurrency-safe list of recycled buffers,[82] implementing coroutines (which helped inspire the name goroutine),[83] and implementing iterators. Concurrency-related structural conventions of Go (channels and alternative channel inputs) are derived from Tony Hoare's communicating sequential processes model. Unlike previous concurrent programming languages such as Occam or Limbo (a language on which Go co-designer Rob Pike worked),[85] Go does not provide any built-in notion of safe or verifiable concurrency.[86] While the communicating-processes model is favored in Go, it is not the only one: all goroutines in a program share a single address space. This means that mutable objects and pointers can be shared between goroutines; see § Lack of race condition safety, below."

	consumer(w, r, text)

}

func consumer(w http.ResponseWriter, r *http.Request, text string) {
	resp, err := http.PostForm("http://localhost:8000/top-ten-words/", url.Values{"text": {text}})

	if err != nil {
		responseJson(w, Response{
			Success: false,
			Message: "fail",
			Data:    nil,
		})
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		responseJson(w, Response{
			Success: false,
			Message: "fail",
			Data:    nil,
		})
	}

	var target Response

	json.Unmarshal(body, &target)

	responseJson(w, Response{
		Success: true,
		Message: "success",
		Data:    target.Data,
	})
}

func responseJson(w http.ResponseWriter, response Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
