#!/bin/bash -xe

exec_subpackages(){
    find . -maxdepth 1 -mindepth 1 \! -name vendor -type d -print0 | xargs -I % -0 $*
}

go get -u github.com/golang/lint/golint

exec_subpackages golint -set_exit_status '%/...'
