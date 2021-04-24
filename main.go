package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func getVmStatus(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("bash", "/Virsh-Api-Server/getVmStatus.sh")
	output, _ := cmd.Output()
	fmt.Fprintf(w, string(output))
	fmt.Println("getVmStatus")
}

func startVM(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("sudo", "virsh", "start", "Windows10")
	output, _ := cmd.Output()
	fmt.Fprintf(w, string(output))
	fmt.Println("startVM")
}

func shutdownVM(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("sudo", "virsh", "shutdown", "Windows10")
	output, _ := cmd.Output()
	fmt.Fprintf(w, string(output))
	fmt.Println("shutdownVM")
}

func forceShutdownVM(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("sudo", "virsh", "destroy", "Windows10")
	output, _ := cmd.Output()
	fmt.Fprintf(w, string(output))
	fmt.Println("forceShutdownVM")
}

func rebootVM(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("sudo", "virsh", "reboot", "Windows10")
	output, _ := cmd.Output()
	fmt.Fprintf(w, string(output))
	fmt.Println("rebootVM")
}

func handleRequests() {
	http.HandleFunc("/getVmStatus", getVmStatus)
	http.HandleFunc("/startVM", startVM)
	http.HandleFunc("/shutdownVM", shutdownVM)
	http.HandleFunc("/forceShutdownVM", forceShutdownVM)
	http.HandleFunc("/rebootVM", rebootVM)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequests()
}
