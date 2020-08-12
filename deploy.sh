#!/bin/bash
cd /app/app

MODULES=$(cat list/enabled-modules.list | tr '\n' '\|')
echo "enabled modules: ${MODULES}"

REPOS=$(cat list/repo.list | grep -E "${MODULES}")
MAIN_CALL=$(cat list/call.list | grep -E "${MODULES}")

function replace(){
    REPLACE_STRING="$1"
    #first match line
    FIRST_MATCH=$(grep -n -m 1 ${REPLACE_STRING} main.go)
    echo $?

    FIRST_MATCH=$(echo ${FIRST_MATCH} | cut -d':' -f1)
    echo $?
    #insert after
    if [ "${REPLACE_STRING}" = "DEPLOY_imports" ];
    then
        echo "Add repos"
        echo "${REPOS}" | sed -i "${FIRST_MATCH}r /dev/stdin" main.go
    else
        echo "Add main call function"
        echo "${MAIN_CALL}" | sed -i "${FIRST_MATCH}r /dev/stdin" main.go
    fi
    #remote firts match
    sed -i "0,/${REPLACE_STRING}/{/${REPLACE_STRING}/d;}" main.go
}

replace "DEPLOY_imports"
replace "DEPLOY_functions"
