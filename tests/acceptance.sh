#!/bin/bash -xe

PATH="..:$PATH" go test ./tests -parallel 2 -test.v
