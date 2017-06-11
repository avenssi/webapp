#! /bin/sh
kill -9 $(pgrep webserver)
cd /webapp/
./webserver &

