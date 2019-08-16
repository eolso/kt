package labrador

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"sync"
	"text/tabwriter"

	"github.com/cheggaaa/pb/v3"

	log "github.com/sirupsen/logrus"
)

// ShowProgress : determines whether a progress bar is shown
var ShowProgress = true

// PodData : struct to store useful pod information
type PodData struct {
	podName   string
	namespace string
	node      string
	cpu       string
	memory    string
}

// Print : prints out the PodData to screen
func (pod PodData) Print() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintf(w, "Name\tCPU\tMemory\n----\t---\t------\n")
	fmt.Fprintf(w, pod.String())

	w.Flush()
}

// String : returns the PodData in string format
func (pod PodData) String() string {
	return fmt.Sprintf("%s\t%s\t%s\t%s\n", pod.podName, pod.cpu, pod.memory, pod.node)
}

// top : runs kubectl top pod on specified pod
func (pod *PodData) top() (err error) {
	argString := fmt.Sprintf("top pod %s --namespace %s --no-headers", pod.podName, pod.namespace)
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
		pod.cpu = lineFields[1]
		pod.memory = lineFields[2]
	}

	if errStr, _ := ioutil.ReadAll(stderr); len(errStr) != 0 {
		return errors.New(string(errStr))
	}

	cmd.Wait()

	return
}

// FetchPod : fetches a single pod
func FetchPod(podName string) (pod PodData, err error) {
	allPods, err := findAll()

	var podFound bool
	podFound = false
	for _, pod = range allPods {
		if pod.podName == podName {
			podFound = true
			break
		}
	}

	if !podFound {
		return (PodData{}), fmt.Errorf("no pods found with name %s", podName)
	}

	pod.top()

	return
}

// FetchPods : runs kubectl get pods across ALL namespaces and stores their data
func FetchPods() (pods []PodData, err error) {
	pods, err = findAll()

	if err != nil {
		return
	}

	topPods(pods)

	log.WithFields(log.Fields{
		"function": "FetchPods()",
	}).Debug("Fetching complete!")

	return
}

// FetchNode : runs kubectl get pods in a node and stores their data
func FetchNode(node string) (pods []PodData) {
	allPods, err := findAll()

	if err != nil {
		return
	}

	for _, pod := range allPods {
		if pod.node == node {
			pods = append(pods, pod)
		}
	}

	topPods(pods)

	log.WithFields(log.Fields{
		"function": "FetchNode()",
	}).Debug("Fetching complete!")

	return
}

// PrettyPrint : formats and prints a list of PodData
func PrettyPrint(pods []PodData) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintf(w, "Name\tCPU\tMemory\tNode\n----\t---\t------\t----\n")
	for _, pod := range pods {
		if pod.cpu != "" {
			fmt.Fprintf(w, pod.String())
		}
	}

	w.Flush()
}

// FindAll : finds all pods
func findAll() (pods []PodData, err error) {
	args := strings.Fields("get pods -Ao wide --no-headers")
	cmd := exec.Command("kubectl", args...)

	stdout, _ := cmd.StdoutPipe()

	if err = cmd.Start(); err != nil {
		return
	}

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		lineFields := strings.Fields(line)

		newPod := PodData{
			podName:   lineFields[1],
			namespace: lineFields[0],
			node:      lineFields[7],
		}

		pods = append(pods, newPod)
	}
	cmd.Wait()

	log.WithFields(log.Fields{
		"function": "findAll()",
	}).Debug("Finding complete!")

	return
}

// topPods : runs kubectl top pod over a slice of pods concurrently
func topPods(pods []PodData) {
	bar := pb.StartNew(len(pods))

	if !ShowProgress {
		bar.SetWriter(ioutil.Discard)
	}

	var wg sync.WaitGroup
	throttle := make(chan struct{}, 30)
	for index := range pods {
		throttle <- struct{}{}
		wg.Add(1)
		go func(index int, bar *pb.ProgressBar) {
			pods[index].top()
			<-throttle

			log.WithFields(log.Fields{
				"podName":   pods[index].podName,
				"namespace": pods[index].namespace,
				"node":      pods[index].node,
				"cpu":       pods[index].cpu,
				"memory":    pods[index].memory,
			}).Debug("pod stored!")

			bar.Increment()
			wg.Done()
		}(index, bar)
	}

	wg.Wait()
	bar.Finish()
}
