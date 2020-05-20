# Requirements Traceability Matrix Tool

Installing to your GOPATH or GOBIN path:

    $ go get github.com/simon3z/tmatrix

Building from sources:

    $ go build

Running from sources:

    $ go run --matrix test-example.yaml
    Components-1 Factor-2, Components-2 Factor-1 (Score: 200)
    Components-1 Factor-1, Components-2 Factor-1 (Score: 150)
    Components-1 Factor-2, Components-2 Factor-2 (Score: 150)
    Components-1 Factor-1, Components-2 Factor-2 (Score: 100)

A test-example.yaml file is provided to illustrate the functionality of the tool.
