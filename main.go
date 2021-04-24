package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"io/ioutil"
)

func getVmStatus(w http.ResponseWriter, r *http.Request) { // POST
	if r.Method == "POST" {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		cmd := exec.Command("bash", "/Virsh-Api-Server/getVmStatus.sh", string(reqBody))
		output, _ := cmd.Output()
		log.Printf("/getVmStatus return -> %s\n", output)
		w.Write([]byte(output))
	}
}

func startVM(w http.ResponseWriter, r *http.Request) { // POST
	if r.Method == "POST" {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		cmd := exec.Command("sudo", "virsh", "start", string(reqBody))
		output, _ := cmd.Output()
		log.Printf("/startVM return -> %s\n", output)
		w.Write([]byte(output))
	}
}

func shutdownVM(w http.ResponseWriter, r *http.Request) { // POST
	if r.Method == "POST" {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		cmd := exec.Command("sudo", "virsh", "shutdown", string(reqBody))
		output, _ := cmd.Output()
		log.Printf("/shutdownVM return -> %s\n", output)
		w.Write([]byte(output))
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
