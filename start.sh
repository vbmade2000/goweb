#/usr/bin/env bash

if [[ -f "goweb" || -f "goweb.exe" ]]; then
	export GIN_MODE=release
	./goweb
fi