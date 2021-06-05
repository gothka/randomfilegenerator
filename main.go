package main

import (
	"context"
	// "fmt"
	"os"
	"log"
	"time"
)

func doEvery(ctx context.Context, d time.Duration, f func(time.Time)) error {
	ticker := time.Tick(d)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case x := <-ticker:
			f(x)
		}
	}
}

func generaterandomfiles(t time.Time) {
	// fmt.Printf("%v: Start writing files!\n", t)
	data := make([]byte, int(1e7), int(1e7))
	f, err := os.CreateTemp("", "test.*.txt")
	if err != nil {
		log.Fatal(err)
	}
	// defer os.Remove(f.Name()) // clean up

	if _, err := f.Write(data); err != nil {
		f.Close()
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	println(f.Name())
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Hour)
	defer cancel() // This cancel doesn't really do anything, but you could register a signal channel that could result in canceling the context.
	doEvery(ctx, 30*time.Second, generaterandomfiles)
}
