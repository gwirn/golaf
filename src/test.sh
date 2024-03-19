#!/bin/bash

if [[ ! $(go run main.go "line" ../testFiles/* | grep line | wc -l) -eq 6 ]];then
    echo "FAILED: Pattern matching failed with default settings and word 'line'"
    exit 1
else
    echo "PASS: Pattern matching with default settings and word 'line' "
fi

if [[ ! $(go run main.go "world" ../testFiles/* | grep world | wc -l) -eq 6 ]];then
    echo "FAILED: Pattern matching failed with default settings and word 'world'"
    exit 1
else
    echo "PASS: Pattern matching with default settings and word 'world' "
fi

if [[ ! $(go run main.go -quality 0 "xyz" ../testFiles/* | wc -l) -eq 15 ]];then
    echo "FAILED: Pattern matching failed where it should match every line"
    exit 1
else
    echo "PASS: Pattern matching with quality 0 "
fi

if [[ ! $(go run main.go -quality 100 "xyz" ../testFiles/* | wc -l) -eq 3 ]];then
    echo "FAILED: Pattern matching failed where it should match nothing and print only the file paths"
    exit 1
else
    echo "PASS: Pattern matching for forced no result "
fi

if [[ ! $(go run main.go -quality 100 "world0" ../testFiles/* | grep world | wc -l) -eq 1 ]];then
    echo "FAILED: Pattern matching failed where it should match only one line"
    exit 1
else
    echo "PASS: Pattern matching for forced one line result "
fi

if [[ ! $(cat ../testFiles/testFile1.txt | go run main.go line | grep line | wc -l) -eq 6 ]];then
    echo "FAILED: Reading from StdIn through pipe failed"
    exit 1
else
    echo "PASS: Pattern matching for input from StdIn "
fi

if [[ ! $(go run main.go -quality 75 wold ../testFiles/* | grep world | wc -l) -eq 6 ]];then
    echo "FAILED: Deletion 'world' -> 'wold'"
    exit 1
else
    echo "PASS: Deletion 'world' -> 'wold'"
fi 

if [[ ! $(go run main.go -quality 75 wod ../testFiles/* | grep ../testFiles/testFile | wc -l) -eq 2 ]];then
    echo "FAILED: Deletion 'world' -> 'wod'"
    exit 1
else
    echo "PASS: Deletion 'world' -> 'wod'"
fi

if [[ ! $(go run main.go -quality 75 lixxne ../testFiles/testFile* | grep line | wc -l) -eq 6 ]];then
    echo "FAILED: Insertion 'line' -> 'lixxne'"
    exit 1
else
    echo "PASS: Insertion 'line' -> 'lixxne'"
fi

if [[ ! $(go run main.go -quality 75 lixxxne ../testFiles/* | grep ../testFiles/testFile | wc -l) -eq 2 ]];then
    echo "FAILED: Insertion 'line' -> 'lixxxne'"
    exit 1
else
    echo "PASS: Insertion 'line' -> 'lixxxne'"
fi

if [[ ! $(go run main.go -recursive=../testFiles hidden | grep testFile | wc -l) -eq 3 ]];then
    echo "FAILED: Only recursevly searching in not hidden files"
    exit 1
else
    echo "PASS: Only recursevly searching in not hidden files"
fi

if [[ ! $(go run main.go -recursive=../testFiles -recH hidden | grep hidden | wc -l) -eq 2 ]];then
    echo "FAILED: Recursevly searching also in hidden files"
    exit 1
else
    echo "PASS: Recursevly searching also in hidden files"
fi
