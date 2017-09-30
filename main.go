package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n+1)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	b[n] = '\n'
	return string(b)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// max log count or write duration
	var max, duration int64
	var speed int64 = 100
	var size int64 = 256
	var logCount int64
	var totalStartAt, totalEndAt int64
	var start time.Time
	var e error
	var debug bool
	var w *bufio.Writer

	s := os.Getenv("MAX")
	if s != "" {
		max, e = strconv.ParseInt(s, 10, 64)
		check(e)
		duration = math.MaxInt64
	} else {
		s = os.Getenv("DURATION")
		if s == "" {
			panic("must use DURATION or MAX")
		} else {
			duration, e = strconv.ParseInt(s, 10, 64)
			check(e)
			duration = duration * 1e9
			max = math.MaxInt64
		}
	}

	s = os.Getenv("SPEED")
	if s != "" {
		speed, e = strconv.ParseInt(s, 10, 64)
		check(e)
	}

	s = os.Getenv("SIZE")
	if s != "" {
		size, e = strconv.ParseInt(s, 10, 64)
		check(e)
	}

	s = os.Getenv("DEBUG")
	if s != "" {
		debug = true
	}

	s = os.Getenv("OUTPUT")
	if s == "" {
		w = bufio.NewWriter(os.Stdout)
	} else {
		if debug {
			fmt.Printf(fmt.Sprintf("\nWrite to file: %s\n", s))
		}
		path := filepath.Dir(s)
		_ = os.MkdirAll(path, os.ModePerm)

		f, err := os.Create(s)
		check(err)

		w = bufio.NewWriter(f)
		defer f.Close()
	}

	logLine := randStringRunes(int(size))
	totalStartAt = time.Now().UnixNano()
	start = time.Now()

	for logCount < max && (time.Now().UnixNano()-totalStartAt) <= duration {
		logCount++

		_, err := w.WriteString(fmt.Sprintf("%d %d %s", time.Now().UnixNano(), logCount, logLine))
		check(err)

		if logCount%speed == 0 && logCount > 0 {
			w.Flush()
			x := time.Now().Sub(start)

			if x < 1e9 {
				if debug {
					fmt.Printf(fmt.Sprintf("\nWrote: %d, Used %.6f, Sleep %.6f\n", logCount, float64(x)/1e9, float64((1e9-x))/1e9))
				}
				time.Sleep(1e9 - x)
			}
			start = time.Now()
		}
	}

	w.Flush()
	totalEndAt = time.Now().UnixNano()
	if debug {
		fmt.Printf("\nExit. Total %d, Used %.6f, Avg speed %.6f\n", logCount, float64((totalEndAt-totalStartAt)/1e9), float64(logCount/((totalEndAt-totalStartAt)/1e9)))
	}

}
