mpdx [![Build Status](https://travis-ci.org/mdlayher/mpdx.svg?branch=master)](https://travis-ci.org/mdlayher/mpdx) [![GoDoc](http://godoc.org/github.com/mdlayher/mpdx?status.svg)](http://godoc.org/github.com/mdlayher/mpdx) [![Report Card](https://goreportcard.com/badge/github.com/mdlayher/mpdx)](https://goreportcard.com/report/github.com/mdlayher/mpdx)
====

Package `mpdx` is an extension of package [mpd](https://github.com/fhs/gompd).
MIT Licensed.

Why?
----

Package `mpd` is excellent, but it returns data in a wrapped `map[string]string`,
instead of a unique `struct` per response type.  This is intentional, for the
reasons described [in this issue](https://github.com/fhs/gompd/issues/31).

Package `mpdx` provides these wrapped `struct` types, for better type safety
and ease of use.
