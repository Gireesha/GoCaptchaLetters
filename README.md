# GoWithPostgreSQL
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/gomods/athens)
[![Open in Visual Studio Code](https://open.vscode.dev/badges/open-in-vscode.svg)](https://open.vscode.dev/Naereen/badges)
[![Ask Me Anything !](https://img.shields.io/badge/Ask%20me-anything-1abc9c.svg)](https://GitHub.com/Naereen/ama)
________________________

 Go code to generate Captcha letters for any TrueType font format files. The code can be modified for the background and font colors. The code can be easily upgraded for distorted and rotated letters. These generated lettes can be stiched together to make captcha string.

## Softwares required
Go compiler --> https://go.dev/doc/install  
Visual Studio code --> https://code.visualstudio.com/download  

## Step 1
Open a command prompt and cd to your home directory.  
  
  On Linux or Mac:  
    
    cd

  On Windows:  

    cd %HOMEPATH%

## Step 2
Create a GoCaptchaLetters directory.
    For example, from your home directory use the following commands:

    mkdir GoCaptchaLetters
    cd GoCaptchaLetters

## Step 3
Start your module using the **go mod init** command.

    go mod init gocaptchaletters

## Step 4
In your text editor, create **main.go** and paste the content and save.

## Step 5
Build the current module's packages and dependencies.

    go mod tidy

## Step 6
Download and install the font. Copy the .ttf file to the GoCaptchaLetters folder. Make chnages in the main.go to read the newly installed .ttf file. 

## Step 7
Build the repo and execute the code in the terminal

    go build .\main.go
    .\main.exe
.png files will be created upon executing the code.

## References
https://go.dev/blog/using-go-modules  
https://go.dev/doc/tutorial/create-module  
https://github.com/fogleman/gg  
https://fonts.google.com/  
https://go.dev/play/p/KcuJ_2c_NDj  

