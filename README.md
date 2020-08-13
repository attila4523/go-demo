# Go demo

## How to install

* Change or Check the list directory files.
* Run the `bash deploy.sh` command.
(It will replace the placeholder in the main.go)

## How to enable/disable a module

Add or Remove it in the list/enabled-modules.list

## How to create a new module

1. You need create a new Github repository and release it. Example: <https://github.com/attila4523/go-demo-module>
1. Release it - you need a tag version
1. Add it to the go.mod
    1. Add the import things to the list/repo.list
    1. Add the call things to the list/call.list
    1. Add the name to the list/enabled-modules.list
