package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func getVmStatus(w http.ResponseWriter, r *http.Request) { // POST
	if r.Method == "POST" {
		cmd := exec.Command("bash", "/Virsh-Api-Server/getVmStatus.sh")
		output, _ := cmd.Output()
		fmt.Fprintf(w, string(output))
		log.Println("getVmStatus")
	}
}

func startVM(w http.ResponseWriter, r *http.Request) { // POST
	if r.Method == "POST" {
		cmd := exec.Command("sudo", "virsh", "start", "Windows10")
		output, _ := cmd.Output()
		fmt.Fprintf(w, string(output))
		log.Println("startVM")
	}
}

func shutdownVM(w http.ResponseWriter, r *http.Request) { // POST
	if r.Method == "POST" {
		cmd := exec.Command("sudo", "virsh", "shutdown", "Windows10")
		output, _ := cmd.Output()
		fmt.Fprintf(w, string(output))
		log.Println("shutdownVM")
	}
}

func forceShutdownVM(w http.ResponseWriter, r *http.Request) { // POST
	if r.Method == "POST" {
		cmd := exec.Command("sudo", "virsh", "destroy", "Windows10")
		output, _ := cmd.Output()
		fmt.Fprintf(w, string(output))
		log.Println("forceShutdownVM")
	}
}

func rebootVM(w http.ResponseWriter, r *http.Request) { // POST
	if r.Method == "POST" {
		cmd := exec.Command("sudo", "virsh", "reboot", "Windows10")
		output, _ := cmd.Output()
		fmt.Fprintf(w, string(output))
		log.Println("rebootVM")
	}
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
