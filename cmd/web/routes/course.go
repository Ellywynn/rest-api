package routes

import "net/http"

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("All courses"))
}
