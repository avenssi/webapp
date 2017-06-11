#! /bin/sh
kill -9 $(pgrep webserver)
cd ~/webapp/
git pull https://github.com/avenssi/webapp.git
cd webserver/
./webserver &

