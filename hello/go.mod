module github.com/mabrarov/hello-golang/hello

go 1.23.2

replace github.com/mabrarov/hello-golang/greetings => ../greetings

require (
	github.com/mabrarov/hello-golang/greetings v0.0.0-00010101000000-000000000000
	rsc.io/quote v1.5.0
)

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)
