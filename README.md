# git

git remote add origin https://github.com/irvinespectrum/stucco.git
git push -u origin master

## Quickstart for Go App Engine Standard Environment 
https://cloud.google.com/appengine/docs/standard/go/quickstart

## install go debugger on Mac
https://github.com/derekparker/delve/blob/master/Documentation/installation/osx/install.md
xcode-select --install
go get -u github.com/derekparker/delve/cmd/dlv
export PATH=$PATH:/users/cc/go/bin
dlv

## Development server 
https://stackoverflow.com/questions/49329710/how-do-i-setup-vscode-debug-session-for-golang-and-appengine
dev_appserver.py app.yaml --go_debugging=true


find pid
ps aux | grep _go_app

start Delve server
dlv --headless -l "localhost:2345" attach $GO_APP_PID


web
http://localhost:8080

admin
http://localhost:8000

## Deploy app

gcloud app deploy


## view application in the web browser

gcloud app browse
https://counter-sa.appspot.com

## production
https://apc.salesamount.com


## vscode debug
F5 "debug go file"
http://localhost:8080/