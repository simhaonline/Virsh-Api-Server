package main

import (
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
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		cmd := exec.Command("sudo", "virsh", "destroy", string(reqBody))
		output, _ := cmd.Output()
		log.Printf("/forceShutdownVM return -> %s\n", output)
		w.Write([]byte(output))
	}
}

func rebootVM(w http.ResponseWriter, r *http.Request) { // POST
	if r.Method == "POST" {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		cmd := exec.Command("sudo", "virsh", "reboot", string(reqBody))
		output, _ := cmd.Output()
		log.Printf("/rebootVM return -> %s\n", output)
		w.Write([]byte(output))
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
