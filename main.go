package main

type Interface interface {
	Run(string) (string, error)
}
