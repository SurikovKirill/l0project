package apiserver

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"l0project/internal/cache"
	"l0project/internal/store"
	"net/http"
	"time"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
	cache  *cache.Cache
}

func New(config *Config, memoryCache *cache.Cache) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
		cache:  memoryCache,
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()
	if err := s.configureStore(); err != nil {
		return err
	}
	if err := s.configureCache(); err != nil {
		return err
	}
	s.logger.Info("Starting server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureCache() error {
	tmp, err := s.store.Order().AllOrders()
	if err != nil {
		return err
	}
	for i := 0; i < len(tmp); i++ {
		s.cache.Set(tmp[i].OrderUID, tmp[i], 10*time.Minute)
	}
	return nil
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
	s.router.HandleFunc("/order/{id}", s.getOrder()).Methods("GET")
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	return nil
}

func (s *APIServer) getOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//type View struct {
		//	Order    string
		//	Delivery string
		//	Payment  string
		//}
		vars := mux.Vars(r)
		key := vars["id"]
		tmp, _ := s.cache.Get(key)
		//o := model.Order{}
		//or, err := json.Marshal(tmp)
		//if err != nil {
		//	log.Println(err)
		//}
		//err = json.Unmarshal(or, &o)
		//if err != nil {
		//	log.Println(err)
		//}
		//data := View{
		//	Order:    o.OrderUID,
		//	Delivery: o.Delivery.Name,
		//	Payment:  o.Payment.Bank,
		//}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//tmpl := template.Must(template.New("data").Parse(`<div>
		//    <h1>{{ .Order}}</h1>
		//    <p>{{ .Delivery}}</p>
		//	<p>{{ .Payment}}</p>
		//</div>`))
		//tmpl.Execute(w, data)
		json.NewEncoder(w).Encode(tmp)
	}
}
