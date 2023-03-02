package apiserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/JohnnyJa/http-rest-api/internal/app/apiclient"
	_ "github.com/JohnnyJa/http-rest-api/internal/app/model"
	"github.com/JohnnyJa/http-rest-api/internal/app/store"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	APIclient *apiclient.APIClient
	store  *store.Store
}

var (
	validate *validator.Validate
)

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()
	s.configureClient()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("Starting API Server")

	validate = validator.New()

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/getMax", s.handleGetMaxPost()).Methods("POST")
	s.router.HandleFunc("/getMax", s.handleGetMaxGet()).Methods("GET")
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)

	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

func (s *APIServer) configureClient(){
	s.APIclient = apiclient.New(s.config.Client)
}

func (s *APIServer) handleGetMaxPost() http.HandlerFunc {
	type request struct {
		Request_id  int    `json:"request_id"`
		Url_package []int  `json:"url_package" validate:"required,dive,required"`
		Ip          string `json:"ip" validate:"required,ipv4"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, 204, err)
			return
		}

		if err := validate.Struct(req); err != nil {
			s.error(w, r, 204, err)
			return
		}
		
		max, err := s.GetMaxSize(req.Url_package)

		if err !=nil {
			s.error(w, r, 204, err)
		}

		s.respond(w, r, 200, map[string]float64{"MaxPrice": max})
		fmt.Println(max)
	}

}

func (s *APIServer) handleGetMaxGet() http.HandlerFunc {
	type request struct {
		Request_id  int    
		Url_package []int  `validate:"required,dive,required"`
		Ip          string `validate:"required,ipv4"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		req.Request_id, _ = strconv.Atoi(r.URL.Query().Get("request_id"))
		req.Url_package = toIntArray(r.URL.Query().Get("url_package"))
		req.Ip = r.URL.Query().Get("ip")

		if err := validate.Struct(req); err != nil {
			s.error(w, r, 204, err)
			return
		}
		
		max, err := s.GetMaxSize(req.Url_package)

		if err !=nil {
			s.error(w, r, 204, err)
		}

		s.respond(w, r, 200, map[string]float64{"MaxPrice": max})
	}

}

func toIntArray(str string) []int {
    chunks := strings.Split(str, ",")

    var res []int
    for _, c := range chunks {
        i, _ := strconv.Atoi(c) // error handling ommitted for concision
        res = append(res, i)
    }

    return res
}

func (s *APIServer) error(w http.ResponseWriter, r *http.Request, code int, err error) {

	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *APIServer) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *APIServer) GetMaxSize(ids []int) (float64, error){
	maxPrice := 0.0
	ch := make(chan apiclient.PriceResult)
	

	for _, id := range ids {
		url, err := s.store.UrlPackage().FindById(id)
		if err != nil{
			return 0, err
		}
		go s.APIclient.GetPrice(url.UrlString, ch)
	}

	for i := 0; i < len(ids); i++ {
		res := <-ch
		
		if res.Error !=nil {
			return 0, res.Error
		}
		
		if res.Price > maxPrice{
			maxPrice = res.Price
		}
	}
	return maxPrice, nil
}