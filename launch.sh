#!/usr/bin/env bash

# https://github.com/fgeller/bilder/blob/master/config.json


env GOOS=linux GOARCH=amd64 go build -o linux.catbook /Users/Kitty/Go/src/github.com/fgeller/bilder/*.go

go build -o test.catbook /Users/Kitty/Go/src/github.com/fgeller/bilder/*.go

rsync -avz ./*  arew:/home/freyabison/sland.rotblauer.com/


	ssh freya <<\EOI
cd sland.rotblauer.com 
./kickstart
exit
EOI
