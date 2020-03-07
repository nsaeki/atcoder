#!/bin/bash

for OPT in "$@"; do
    case $OPT in
        -v) VERBOSE=1;;
        *)
            if [[ -z "$SOURCE" ]]; then
                SOURCE=$OPT
            else
                TESTCASE="$TESTCASE $OPT"
            fi
            ;;
    esac
done

if [[ -z "$SOURCE" ]]; then
    echo "Error: specify a source file"
    echo "Usage: $0 source [testcase_file] [-v]"
    exit 1
fi

DIR=$(dirname $SOURCE)
TASK=$(echo $SOURCE | cut -f 1 -d_)
RUN_FUNC=run_${SOURCE#*.}

if [[ -z "$TESTCASE" ]]; then
    TESTCASE=$(ls $DIR/$TASK*.txt)
fi

run_go() {
    go run $1
}

run_rs() {
    rustc $1 -o $DIR/a.out
    $DIR/a.out
}

run_testcase() {
    local input=${1:?}
    if [[ -z "$VERBOSE" ]]; then
        echo -n "$input: "
    else
        echo -e "\n========> Testing file ${input}"
        echo "Input:"
        cat $input
        echo -e "\n\nResult:"
    fi
    ${RUN_FUNC} ${SOURCE} < $input
}

# Main

if [[ -n "$TESTCASE" ]]; then
   for i in $TESTCASE; do
        run_testcase $i        
    done
else
    echo "No test files are found. Waiting input"
    cat | ${run_func} ${SOURCE}
fi
