#!/usr/bin/env bash

# https://github.com/fgeller/bilder/blob/master/config.json


env GOOS=linux GOARCH=amd64 go build -o linux.catbook /Users/Kitty/Go/src/github.com/fgeller/bilder/*.go

go build -o test.catbook /Users/Kitty/Go/src/github.com/fgeller/bilder/*.go

rsync -avz ./*  arew:/home/freyabison/sland.rotblauer.com/
rsync -avz /Volumes/Data/iceland/sharePic/*.jpg arew:/home/freyabison/sland.rotblauer.com/img/iceland/
rsync -avz /Volumes/Data/boots_bandit/*.jpg arew:/home/freyabison/sland.rotblauer.com/img/cats/


	ssh freya <<\EOI
cd sland.rotblauer.com 
./kickstart
exit
EOI
