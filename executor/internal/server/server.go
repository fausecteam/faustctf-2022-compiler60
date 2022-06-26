package server

import (
	"bytes"
	"encoding/json"
	"executor/internal/file_pruner"
	"executor/internal/sandbox_runner"
	"executor/pkg/compiler60_shared"
	"executor/pkg/signed_elf"
	"fmt"
	"net/http"
	"os"
	"time"
)

const (
	corsOrigin  = "Access-Control-Allow-Origin"
	corsHeaders = "Access-Control-Allow-Headers"
)

func Start() {
	largeWhenDebug := time.Duration(1)
	if sandbox_runner.IsDebug() {
		fmt.Println("Starting executor in DEBUG mode")
		largeWhenDebug = 99999
	}

	var executeHandler http.Handler
	executeHandler = http.HandlerFunc(HandleExecute)
	executeHandler = ContentTypeFilter(executeHandler, compiler60_shared.HttpJsonType)
	executeHandler = CORSAwareRequestMethodFilter(executeHandler, "POST", "*")
	executeHandler = http.TimeoutHandler(executeHandler, 8*time.Second*largeWhenDebug, "Retry later\n")

	http.Handle("/execute", executeHandler)

	srv := &http.Server{
		Addr:              ":" + compiler60_shared.ExecutorPort,
		ReadHeaderTimeout: 1 * time.Second,
		ReadTimeout:       3 * time.Second,
		WriteTimeout:      6 * time.Second * largeWhenDebug,
		IdleTimeout:       3 * time.Second,
	}

	fmt.Printf("Starting server on %s...\n", srv.Addr)
	err := srv.ListenAndServe()
	fmt.Printf("listen err %v\n", err)
	os.Exit(1)
}

func HandleExecute(w http.ResponseWriter, req *http.Request) {
	// TODO reorder this for better efficiency

	file_pruner.SchedulePruning()

	req.Body = http.MaxBytesReader(w, req.Body, signed_elf.MaxEncodedSize)

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()

	var elfToRun signed_elf.SignedElf
	err := dec.Decode(&elfToRun)
	if err != nil {
		http.Error(w, "illegal json: "+err.Error(), http.StatusBadRequest)
		return
	}

	if !sandbox_runner.IsDebug() {
		err = elfToRun.VerifyBinarySignature()
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
	}

	runner := &sandbox_runner.Runner{
		ExecFile: elfToRun.Binary,
	}

	// ctx will have timeout because of http.TimeoutHandler
	ctx := req.Context()
	run := runner.Run(ctx)

	if ctx.Err() != nil {
		return // request already canceled
	}

	// encode to buffer, json does not stream anyway
	var buffer bytes.Buffer
	enc := json.NewEncoder(&buffer)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(true)
	err = enc.Encode(run)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(compiler60_shared.HttpContentType, compiler60_shared.HttpJsonType)
	w.WriteHeader(http.StatusOK)
	// ignore err, cannot change header now
	_, _ = w.Write(buffer.Bytes())
}

func ContentTypeFilter(handler http.Handler, neededType string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch req.Header.Get(compiler60_shared.HttpContentType) {
		case neededType:
			handler.ServeHTTP(w, req)
		default:
			http.Error(w, "Expected "+compiler60_shared.HttpContentType+": "+neededType, http.StatusUnsupportedMediaType)
		}
	})
}

func CORSAwareRequestMethodFilter(handler http.Handler, method string, corsConfig string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "OPTIONS":
			w.Header().Set(corsOrigin, corsConfig)
			w.Header().Set(corsHeaders, compiler60_shared.HttpContentType)
			w.WriteHeader(http.StatusNoContent)
		case method:
			w.Header().Set(corsOrigin, corsConfig)
			w.Header().Set(corsHeaders, compiler60_shared.HttpContentType)
			handler.ServeHTTP(w, req)
		default:
			http.Error(w, "Only "+method+" allowed", http.StatusMethodNotAllowed)
		}
	})
}
