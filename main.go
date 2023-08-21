package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/guoyk93/rg"
)

func main() {
	var err error
	defer func() {
		if err == nil {
			return
		}
		log.Println("exited with error:", err.Error())
		os.Exit(1)
	}()
	defer rg.Guard(&err)

	var (
		optConf string
		optDry  bool
	)
	flag.StringVar(&optConf, "conf", "/etc/dellog.d", "directory to load configuration from")
	flag.BoolVar(&optDry, "dry", false, "dry run")
	flag.Parse()

	rules := rg.Must(LoadRules(optConf))
	ExpandRules(&rules)

	for _, rule := range rules {
		for _, matched := range rg.Must(filepath.Glob(rule.Pattern)) {
			log.Println("found:", matched)
			switch EvaluateDays(matched, rule.Days) {
			case ActionUnknownDate:
				log.Println("unknown date:", matched)
			case ActionSkip:
				log.Println("skip:", matched)
			case ActionRemove:
				log.Println("remove:", matched)
				if !optDry {
					_ = os.Remove(matched)
				}
			}
		}
	}
}
