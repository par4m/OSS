package main

import (
	"flag"
	"golang.org/x/net/atom"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"strings"
)

type Link struct {
	url   string
	text  string
	depth int
}

type HttpError struct {
}

func main() {

}
