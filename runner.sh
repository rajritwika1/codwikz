#!/bin/bash
LANGUAGE=$1
CODE_FILE=$2

case $LANGUAGE in
    "python")
        python3 $CODE_FILE
        ;;
    "cpp")
        g++ $CODE_FILE -o code.out && ./code.out
        ;;
    "go")
        go run $CODE_FILE
        ;;
    "java")
        javac $CODE_FILE && java $(basename $CODE_FILE .java)
        ;;
    *)
        echo "Unsupported language"
        exit 1
        ;;
esac
