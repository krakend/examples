package main

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/url"
	"strings"
	"time"
)

var (
	errUnkownRequestType  = errors.New("unknown request type")
	errUnkownResponseType = errors.New("unknown response type")
)

type registerer string

func handleResponseStream(respw ResponseWrapper) interface{} {
	pr, pw := io.Pipe()
	go func() {
		defer pw.Close()
		in := bufio.NewReader(respw.Io())
		for {
			line, err := in.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					return
				}
				pw.CloseWithError(err)
				return
			}

			trimmedline := strings.TrimSpace(line)
			if trimmedline == "" {
				_, err := pw.Write([]byte(line))
				if err != nil {
					pw.CloseWithError(err)
					return
				}
				continue
			}

			if !strings.HasPrefix(trimmedline, "data:") {
				_, err := pw.Write([]byte(line))
				if err != nil {
					pw.CloseWithError(err)
					return
				}
				continue
			}
			var m map[string]interface{}

			data := strings.TrimSpace(trimmedline[5:])
			if data == "[DONE]" {
				_, err := pw.Write([]byte(line))
				if err != nil {
					pw.CloseWithError(err)
				}
				return
			}

			err = json.Unmarshal([]byte(data), &m)
			if err != nil {
				pw.CloseWithError(err)
				return
			}
			// example to manipulate incoming messages adding a processed_time key
			m["processed_time"] = time.Now().Format(time.RFC3339Nano)

			b, err := json.Marshal(m)
			if err != nil {
				pw.CloseWithError(err)
				return
			}

			_, err = pw.Write(append([]byte("data: "), b...))
			if err != nil {
				pw.CloseWithError(err)
				return
			}
			_, err = pw.Write([]byte("\n"))
			if err != nil {
				pw.CloseWithError(err)
				return
			}
		}
	}()

	return streamResponseWrapper{
		ResponseWrapper: respw,
		in:              pr,
	}
}

func main() {}

type RequestWrapper interface {
	Context() context.Context
	Params() map[string]string
	Headers() map[string][]string
	Body() io.ReadCloser
	Method() string
	URL() *url.URL
	Query() url.Values
	Path() string
}

type ResponseWrapper interface {
	Context() context.Context
	Request() interface{}
	Data() map[string]interface{}
	IsComplete() bool
	Headers() map[string][]string
	StatusCode() int
	Io() io.Reader
}

type streamResponseWrapper struct {
	ResponseWrapper
	in io.Reader
}

func (s streamResponseWrapper) Io() io.Reader {
	return s.in
}
