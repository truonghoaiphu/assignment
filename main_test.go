package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	"assignment/api/controllers"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())

}

func TestRandomMessage(t *testing.T) {
	type args struct {
		in0 uint
	}
	tests := []struct {
		message string
		args    args
		want    bool
	}{
		{
			"test",
			args{0},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.message, func(t *testing.T) {
			if got := len(controllers.RandomMessage().Content); got < 1 {
				t.Errorf("realRand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func readLog( /*...*/ ) {
	// ...
	err := io.EOF // force an error
	if err != nil {
		fmt.Println("ERROR")
		log.Print("Couldn't read first byte")
		return
	}
	// ...
}

func TestReadLog(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	readLog()
	t.Log(buf.String())
}
