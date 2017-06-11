package main 

import (
    "io"
    "log"
    "net/http"
    "os/exec"
)

func pullCode() {
	// Enter the directory
	cmd := exec.Command("cd", "~/webapp/")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()


	// Pull code
	cmd = exec.Command("git", "pull", "https://github.com/avenssi/webapp.git")
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
	log.Println("pulling code command finished.")
}

func reLaunch() {
	cmd := exec.Command("sh", "./test.sh")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}


func firstPage(w http.ResponseWriter, r *http.Request) {
	pullCode()
	reLaunch()
	io.WriteString(w, "<h1>Hello, I am avenssi! </h1>")
}


func main() {
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":5000", nil)
}
