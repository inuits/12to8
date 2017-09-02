#!/bin/bash -xe

exec_subpackages(){
    find . -maxdepth 1 -mindepth 1 \! -name tests \! -name vendor -type d -print -exec $* ';'
}

exec_subpackages go test -v '{}/...'
exec_subpackages go vet '{}/...'
