# Auth CLI

Have created this tool for the CLI tool for the gRPC auth service I'm writing.
This CLI has been built using the [Cobra framework](https://github.com/spf13/cobra).

<!-- vim-markdown-toc GFM -->

* [Test coverage](#test-coverage)

<!-- vim-markdown-toc -->

## Test coverage

Provided that you have installed [gocov](https://github.com/axw/gocov), it should
be possible to run the test suite in the `sessions` package can be run with the
following command:

    gocov test -count=1 ./sessions | gocov report
