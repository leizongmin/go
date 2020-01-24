#!/bin/sh

for f in `ls .`
do
    if [ -d $f ]
    then
        p="github.com/leizongmin/go/$f"
        echo "-------------"
        echo $p
        go test $p
        echo "-------------"
    fi
done
