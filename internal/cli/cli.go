package cli

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"uniqstr/internal/service"
)

func Run() {
	cfg, err := argParse()
	if err != nil {
		log.Fatalf("err: parse arg err: %v", err)
	}
	var str string
	if cfg.Input == "" {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("err: read from stdin err: %v", err)
		}
		str = string(data)
	} else {
		file, err := os.ReadFile(cfg.Input)
		if err != nil {
			log.Fatalf("err: read input file err: %v", err)
		}
		str = string(file)
	}
	sl := strings.Split(str, "\n")
	uniq := service.NewUniqService()
	result, err := uniq.Process(sl, cfg)
	if err != nil {
		log.Fatalf("err: process err: %v", err)
	}
	if cfg.Output == "" {
		fmt.Println(result)
	} else {
		if err := os.WriteFile(cfg.Output, []byte(result), 0644); err != nil {
			log.Fatalf("err: write output file err: %v", err)
		}
	}
}
