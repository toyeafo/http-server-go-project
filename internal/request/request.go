package request

import (
	"errors"
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

const crlf = "\r\n"

func RequestFromReader(reader io.Reader) (*Request, error) {

	requestString, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}

	httpRequestLine := parseRequestLine(requestString)
	requestLine := strings.Split(httpRequestLine, " ")

	if len(requestLine) != 3 {
		return nil, fmt.Errorf("invalid number of parts in request line: %s", httpRequestLine)
	}

	method, target, version := requestLine[0], requestLine[1], requestLine[2]

	if !regexp.MustCompile(`^[a-zA-Z]*$`).MatchString(method) {
		return nil, errors.New("invalid method in request line")
	}

	httpVersion := strings.Split(version, "/")
	if len(httpVersion) != 2 {
		return nil, errors.New("invalid http version in request line")
	}

	if httpVersion[0] != "HTTP" {
		return nil, errors.New("invalid http version in request line")
	}

	newRequestLine := &RequestLine{
		Method:        method,
		HttpVersion:   httpVersion[1],
		RequestTarget: target,
	}

	return &Request{
		RequestLine: *newRequestLine,
	}, nil
}

func parseRequestLine(requestString []byte) string {
	line := strings.Split(string(requestString), crlf)
	return line[0]
}
