package delivery

// import (
// 	"html/template"
// 	"net/http"
// )

// func ErrorPage(w http.ResponseWriter, error string, code int) {
// 	response := struct {
// 		ErrorCode int
// 		ErrorText string
// 	}{
// 		ErrorCode: code,
// 		ErrorText: error,
// 	}
// 	// w.WriteHeader(code)
// 	tmpl, _ := template.ParseFiles("ui/templates/error.html")
// 	tmpl.Execute(w, response)
// }
