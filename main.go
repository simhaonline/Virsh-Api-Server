package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

var Auth_Token = "virsh_api_server"

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Auth_Token")
}

func getVmStatus(w http.ResponseWriter, r *http.Request) { // POST
	enableCors(&w)
	if r.Method == "POST" {
		if r.Header.Get("Auth_Token") == Auth_Token {
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Fatal(err)
			}
			cmd := exec.Command("bash", "/Virsh-Api-Server/getVmStatus.sh", string(reqBody))
			output, _ := cmd.Output()
			log.Printf("/getVmStatus return -> %s\n", output)
			w.Write([]byte(output))
		} else {
			w.Write([]byte("Error Bad Auth-Token"))
		}
	}
}

func startVM(w http.ResponseWriter, r *http.Request) { // POST
	enableCors(&w)
	if r.Method == "POST" {
		if r.Header.Get("Auth_Token") == Auth_Token {
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Fatal(err)
			}
			cmd := exec.Command("sudo", "virsh", "start", string(reqBody))
			output, _ := cmd.Output()
			log.Printf("/startVM return -> %s\n", output)
			w.Write([]byte(output))
		} else {
			w.Write([]byte("Error Bad Auth-Token"))
		}
	}
}

func shutdownVM(w http.ResponseWriter, r *http.Request) { // POST
	enableCors(&w)
	if r.Method == "POST" {
		if r.Header.Get("Auth_Token") == Auth_Token {
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Fatal(err)
			}
			cmd := exec.Command("sudo", "virsh", "shutdown", string(reqBody))
			output, _ := cmd.Output()
			log.Printf("/shutdownVM return -> %s\n", output)
			w.Write([]byte(output))
		} else {
			w.Write([]byte("Error Bad Auth-Token"))
		}
	}
}

func forceShutdownVM(w http.ResponseWriter, r *http.Request) { // POST
	enableCors(&w)
	if r.Method == "POST" {
		if r.Header.Get("Auth_Token") == Auth_Token {
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Fatal(err)
			}
			cmd := exec.Command("sudo", "virsh", "destroy", string(reqBody))
			output, _ := cmd.Output()
			log.Printf("/forceShutdownVM return -> %s\n", output)
			w.Write([]byte(output))
		} else {
			w.Write([]byte("Error Bad Auth-Token"))
		}
	}
}

func rebootVM(w http.ResponseWriter, r *http.Request) { // POST
	enableCors(&w)
	if r.Method == "POST" {
		if r.Header.Get("Auth_Token") == Auth_Token {
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Fatal(err)
			}
			cmd := exec.Command("sudo", "virsh", "reboot", string(reqBody))
			output, _ := cmd.Output()
			log.Printf("/rebootVM return -> %s\n", output)
			w.Write([]byte(output))
		} else {
			w.Write([]byte("Error Bad Auth-Token"))
		}
	}
}

func updateXML(w http.ResponseWriter, r *http.Request) { // POST
	enableCors(&w)
	if r.Method == "POST" {
		if r.Header.Get("Auth_Token") == Auth_Token {
			cmd := exec.Command("sudo", "virsh", "define", "/var/lib/libvirt/xml/windows10_vm_template.xml")
			output, _ := cmd.Output()
			log.Printf("/updateXML return -> %s\n", output)
			w.Write([]byte(output))
		} else {
			w.Write([]byte("Error Bad Auth-Token"))
		}
	}
}

func handleRequests() {
	http.HandleFunc("/getVmStatus", getVmStatus)
	http.HandleFunc("/startVM", startVM)
	http.HandleFunc("/shutdownVM", shutdownVM)
	http.HandleFunc("/forceShutdownVM", forceShutdownVM)
	http.HandleFunc("/rebootVM", rebootVM)
	http.HandleFunc("/updateXML", updateXML)
	log.Println("server started on port 8080")
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequests()
}
