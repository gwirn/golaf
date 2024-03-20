#!/bin/bash

go build -o ../testFiles/binfile
buildsuc="$?"
if [[ ! "$buildsuc" -eq "0" ]];then
    echo "Build failed"
    exit 1
fi

if [[ ! $(../testFiles/binfile "line" ../testFiles/* | grep line | wc -l) -eq 6 ]];then
    echo "FAILED: Pattern matching failed with default settings and word 'line'"
    exit 1
else
    echo "PASS: Pattern matching with default settings and word 'line' "
fi

if [[ ! $(../testFiles/binfile "world" ../testFiles/* | grep world | wc -l) -eq 6 ]];then
    echo "FAILED: Pattern matching failed with default settings and word 'world'"
    exit 1
else
    echo "PASS: Pattern matching with default settings and word 'world' "
fi

if [[ ! $(../testFiles/binfile -quality 0 "xyz" ../testFiles/* | wc -l) -eq 17 ]];then
    echo "FAILED: Pattern matching failed where it should match every line"
    exit 1
else
    echo "PASS: Pattern matching with quality 0 "
fi

if [[ ! $(../testFiles/binfile -quality 100 "xyz" ../testFiles/* | wc -l) -eq 3 ]];then
    echo "FAILED: Pattern matching failed where it should match nothing and print only the file paths"
    exit 1
else
    echo "PASS: Pattern matching for forced no result "
fi

if [[ ! $(../testFiles/binfile -quality 100 "world0" ../testFiles/* | grep world | wc -l) -eq 1 ]];then
    echo "FAILED: Pattern matching failed where it should match only one line"
    exit 1
else
    echo "PASS: Pattern matching for forced one line result "
fi

if [[ ! $(cat ../testFiles/testFile1.txt | ../testFiles/binfile line | grep line | wc -l) -eq 6 ]];then
    echo "FAILED: Reading from StdIn through pipe failed"
    exit 1
else
    echo "PASS: Pattern matching for input from StdIn "
fi

if [[ ! $(../testFiles/binfile -quality 75 wold ../testFiles/* | grep world | wc -l) -eq 6 ]];then
    echo "FAILED: Deletion 'world' -> 'wold'"
    exit 1
else
    echo "PASS: Deletion 'world' -> 'wold'"
fi 

if [[ ! $(../testFiles/binfile -quality 75 wod ../testFiles/* | grep ../testFiles/testFile | wc -l) -eq 2 ]];then
    echo "FAILED: Deletion 'world' -> 'wod'"
    exit 1
else
    echo "PASS: Deletion 'world' -> 'wod'"
fi

if [[ ! $(../testFiles/binfile -quality 75 lixxne ../testFiles/testFile* | grep line | wc -l) -eq 6 ]];then
    echo "FAILED: Insertion 'line' -> 'lixxne'"
    exit 1
else
    echo "PASS: Insertion 'line' -> 'lixxne'"
fi

if [[ ! $(../testFiles/binfile -quality 75 lixxxne ../testFiles/* | grep ../testFiles/testFile | wc -l) -eq 2 ]];then
    echo "FAILED: Insertion 'line' -> 'lixxxne'"
    exit 1
else
    echo "PASS: Insertion 'line' -> 'lixxxne'"
fi

if [[ ! $(../testFiles/binfile -recursive=../testFiles FileContent | grep FileContent | wc -l) -eq 3 ]];then
    echo "FAILED: Only recursevly searching in not hidden files"
    exit 1
else
    echo "PASS: Only recursevly searching in not hidden files"
fi

if [[ ! $(../testFiles/binfile -recursive=../testFiles -recH FileContent | grep FileContent | wc -l) -eq 6 ]];then
    echo "FAILED: Recursevly searching also in hidden files"
    exit 1
else
    echo "PASS: Recursevly searching also in hidden files"
fi

if [[ ! $(../testFiles/binfile -recursive=../testFiles -type=n test | wc -l) -eq 3 ]];then
    echo "FAILED: Recursive file name search"
    exit 1
else
    echo "PASS: Recursive file name search"
fi

if [[ ! $(../testFiles/binfile -recursive="../testFiles" -type="n" -recH test | wc -l) -eq 6 ]];then
    echo "FAILED: Recursive file name search including hidden files"
    exit 1
else
    echo "PASS: Recursive file name search including hidden files"
fi

if [[ ! $(../testFiles/binfile -recursive="../testFiles" -type="n" file | grep -v bin | wc -l) -eq 3 ]];then
    echo "FAILED: Recursive file name search excluding binary files"
    exit 1
else
    echo "PASS: Recursive file name search excluding binary files"
fi

if [[ ! $(../testFiles/binfile -recursive="../testFiles" -type="n" -binary file | grep bin | wc -l) -eq 1 ]];then
    echo "FAILED: Recursive file name search including binary files"
    exit 1
else
    echo "PASS: Recursive file name search including binary files"
fi

rm ../testFiles/binfile
