package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/Rammurthy5/algos/src"
)

func runPythonScript(args ...string) {
	cmd := exec.Command("python3", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("result: %s", strings.TrimSpace(string(output)))
}

func main() {
	src.LinkedListRun()
	src.HashMapRun()
	src.BinaryTreeRun()
	src.BinarySearchTreeRun()
	src.StacksQueuesRun()
	src.TrieRun()
	src.CompressedRadixTrieRun()
	runPythonScript("src/kadane_algorithm.py", "1", "-2", "3", "4", "-1", "2", "1", "-5", "4")
	runPythonScript("src/recursion_backtracking.py", "subsequences", "abc")
	runPythonScript("src/recursion_backtracking.py", "print_sequences_target", "10", "1", "2", "3", "4", "5")
	runPythonScript("src/recursion_backtracking.py", "print_one_sequence_target", "10", "1", "2", "3", "4", "5")
	runPythonScript("src/recursion_backtracking.py", "count_sequences_target", "10", "1", "2", "3", "4", "5")
}
