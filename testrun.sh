#!/bin/bash
set -euo pipefail

TESTCASE=()
for OPT in "$@"; do
    case $OPT in
    -v) VERBOSE=1 ;;
    *)
        if [[ -z "${SOURCE:-}" ]]; then
            SOURCE=$OPT
        else
            TESTCASE+=("$OPT")
        fi
        ;;
    esac
done

if [[ -z "${SOURCE:-}" ]]; then
    echo "Error: Source file is not specified!"
    echo ""
    echo "Usage: $0 <source_file> [testcase_file(s)] [-v]"
    echo ""
    echo "  <source_file>      : The source code file to run (e.g., solution.go, solution.rs)."
    echo "  [testcase_file(s)] : Optional. One or more test case files containing input data."
    echo "  -v                 : Optional. Verbose mode, prints detailed information during test runs."
    echo ""
    echo "Example:"
    echo "  $0 solution.go test1.txt test2.txt -v"
    exit 1
fi

DIR=$(dirname "$SOURCE")
TASK=$(echo "$SOURCE" | cut -f 1 -d_)
RUN_FUNC=run_${SOURCE#*.}

run_go() {
    go run "$1"
}

run_rs() {
    rustc "$1" -o "$DIR/a.out"
    "$DIR"/a.out
}

run_testcase() {
    local input=${1:?}
    if [[ -z "${VERBOSE:-}" ]]; then
        printf "%s: " "$input"
    else
        printf "\n========> Testing %s" "$input"
        printf "\nInput:\n%s" "$(cat "$input")"
        printf "\n\nResult:\n"
    fi
    $RUN_FUNC "$SOURCE" <"$input"
}

# Main

if [[ -z "${TESTCASE:-}" ]]; then
    shopt -s nullglob
    TESTCASE=("$DIR"/"$TASK"*.txt)
    shopt -u nullglob
fi

if [[ ${#TESTCASE[@]} -gt 0 ]]; then
    for i in "${TESTCASE[@]}"; do
        run_testcase "$i"
    done
else
    echo "No test files are found. Waiting for input"
    $RUN_FUNC "$SOURCE"
fi
