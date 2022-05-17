package bothook

import (
	"fmt"
	"net/http"
)

func PostCall(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	//fmt.Println(params)

	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/x-www-form-urlencoded" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	r.ParseForm()

	fmt.Println("----------------post form below--------------------------")
	fmt.Println("request.PostForm::")
	for key, value := range r.PostForm {
		fmt.Printf("Key:%s, Value:%s\n", key, value)
	}

	//w.Header().Add("Location", "http://openknowledge.network/id/ldn/1234/inbox/1")
	//w.WriteHeader(201)
	w.WriteHeader(200)
	fmt.Fprintf(w, "Zoltan is on it!")

	Reply(r.PostForm["response_url"], r.PostForm["text"])

}
