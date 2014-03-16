package main

import (
	"flag"
	"github.com/coopernurse/gorp"
	"github.com/gorilla/sessions"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"github.com/stretchr/goweb/handlers"
	log "github.com/umisama/golog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"time"
)

var (
	logger        log.Logger
	flagAddress   *string = flag.String("addr", ":8080", "listen addr")
	flagDirPath   *string = flag.String("dir", "ui/", "path to static files dir")
	dbmap         *gorp.DbMap
	session_store = sessions.NewCookieStore([]byte("secure"))
	session_name  = "twittor-session"
)

type ResponseMeta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseBase struct {
	Meta    ResponseMeta `json:"meta"`
	Content interface{}  `json:"content"`
}

func main() {
	flag.Parse()
	logger, _ = log.NewLogger(os.Stdout, log.TIME_FORMAT_MILLISEC, log.LOG_FORMAT_POWERFUL, log.LogLevel_Debug)

	err := wMain()
	if err != nil {
		logger.Info(err)
		os.Exit(2)
		return
	}

	os.Exit(0)
	return
}

func wMain() (err error) {
	dbmap, err = connectToDb()
	if err != nil {
		return err
	}

	goweb.API.SetStandardResponseObjectTransformer(formatApi)
	err = mapRoutes()
	if err != nil {
		return
	}

	err = listenAndServe(*flagAddress)
	return
}

func mapRoutes() (err error) {
	goweb.MapBefore(beforeHandler)
	goweb.MapAfter(afterHandler)

	// static files
	goweb.MapStatic("/", *flagDirPath, func(c context.Context)(handlers.MatcherFuncDecision, error){
		if regexp.MustCompile(`^api`).MatchString(c.Path().RawPath) {
			return handlers.NoMatch, nil
		}
		return handlers.Match, nil
	})

	// apis
	goweb.MapController("/api/session", &SessionController{})
	goweb.MapController("/api/statuses", &StatusesController{})
	goweb.MapController("/api/user", &UserController{})
	goweb.MapController("/api/search", &SearchController{})

	return nil
}

func listenAndServe(addr string) (err error) {
	s := &http.Server{
		Addr:           addr,
		Handler:        goweb.DefaultHttpHandler(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Info("Could not listen", err)
	}

	go func() {
		for _ = range c {
			listener.Close()
		}
	}()

	err = s.Serve(listener)
	return
}

func formatApi(ctx context.Context, obj interface{}) (ret interface{}, err error) {
	obj_map := obj.(map[string]interface{})

	stat := obj_map["s"].(int)
	message, ok := obj_map["e"].([]string)
	if !ok {
		message = []string{"complete"}
	}

	ret = ResponseBase{
		Meta: ResponseMeta{
			Code:    stat,
			Message: message[0],
		},
		Content: obj_map["d"],
	}
	return
}
