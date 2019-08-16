package labrador

import (
	"bufio"
	"errors"
	"io/ioutil"
	"os/exec"
	"strings"
	"testing"
)

func TestFetchPod(t *testing.T) {
	badPodName := "definitelynotapod"
	goodPod, err := getGoodPod()

	if err != nil {
		t.Fatalf("Failed to get good pod name: %s", err.Error())
	}

	badPod, err := FetchPod(badPodName)
	if badPod != (PodData{}) {
		t.Errorf("Expected badPod PodData to be empty, got:\n\t%s", badPod.String())
	}
	if err == nil {
		t.Errorf("Expected FetchPod(badPodName) to return an error")
	}

	t.Logf("testing with pod %s", goodPod.podName)
	goodPod, err = FetchPod(goodPod.podName)
	if goodPod == (PodData{}) {
		t.Errorf("Expected goodPod PodData to contain values")
	}
	if err != nil {
		t.Errorf("FetchPod(goodPodName) encountered an error: %s", err.Error())
	}
}

func TestTop(t *testing.T) {
	var err error

	badPod := PodData{
		podName:   "definitelynotapod",
		namespace: "definitelynotans",
		node:      "definitelynotanode",
	}

	goodPod, err := getGoodPod()
	if err != nil {
		t.Fatalf("Failed to get good pod name: %s", err.Error())
	}

	err = badPod.top()
	if err == nil {
		t.Errorf("Error expected for badPod.top()")
	}

	t.Logf("testing top with pod %s", goodPod.podName)
	err = goodPod.top()
	if err != nil {
		t.Errorf("goodPod.top() encountered an error: %s", err.Error())
	}
}

func getGoodPod() (goodPod PodData, err error) {
	argString := "get pods -A --field-selector=status.phase==Running --no-headers"
	args := strings.Fields(argString)
	cmd := exec.Command("kubectl", args...)

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	if err = cmd.Start(); err != nil {
		return
	}

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		lineFields := strings.Fields(line)
		goodPod = PodData{
			podName:   lineFields[1],
			namespace: lineFields[0],
		}
		break
	}

	if errStr, _ := ioutil.ReadAll(stderr); len(errStr) != 0 {
		return (PodData{}), errors.New(string(errStr))
	}

	cmd.Wait()

	return
}
