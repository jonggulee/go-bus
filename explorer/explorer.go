package explorer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/jonggulee/gbis/bus"
)

const (
	port        int    = 4000
	templateDir string = "explorer/templates/"
)

var templates *template.Template

type homeData struct {
	PageTitle   string
	Buses       []bus.Bus
	StationName string
	NowTime     string
}

func getNowTime() string {
	now := time.Now()
	kst, _ := time.LoadLocation("Asia/Seoul")
	kstTime := now.In(kst)
	return kstTime.Format("2006-01-02 15:04:05 KST")
}

func home(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	kstTime := getNowTime()
	arsId := "48626" // 위례중앙중학교 정류소 번호
	data := homeData{"Home", bus.GetArrivalBus(arsId), "위례중앙중학교", kstTime}
	templates.ExecuteTemplate(rw, "home", data)
}

func homeJake(rw http.ResponseWriter, r *http.Request) {
	kstTime := getNowTime()
	arsId := "28532" // 위례중앙중학교 정류소 번호
	data := homeData{"Home", bus.GetArrivalBus(arsId), "하남스타필드시티", kstTime}
	templates.ExecuteTemplate(rw, "home", data)
}

func health(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
}

func Start() {
	handler := http.NewServeMux()
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	handler.HandleFunc("/", home)
	handler.HandleFunc("/jake", homeJake)
	handler.HandleFunc("/health", health)
	fmt.Printf("Listening on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
