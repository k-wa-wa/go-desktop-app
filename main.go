package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os/exec"
	"sync"
)

//go:embed web/out/*
var static embed.FS

func main() {
	public, err := fs.Sub(static, "web/out")
	if err != nil {
		panic(err)
	}
	http.Handle("/", http.FileServer(http.FS(public)))

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		log.Fatal(http.ListenAndServe("localhost:8080", nil))
	}()

	args := []string{
		"-n",
		"-a", "Google Chrome",
		"--args",
		"--app=http://localhost:8080",
		"--disable-plugins",
		"--disable-popup-blocking",
		"--disable-dev-tools",
		"--disable-extensions",
		"--disable-sync",
		"--disable-hang-monitor",
		"--disable-prompt-on-repost",
		"--disable-print-preview",
	}
	err = exec.Command("open", args...).Run()
	if err != nil {
		log.Fatal(err)
	}
	wg.Wait()
}
