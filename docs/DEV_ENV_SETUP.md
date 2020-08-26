# AKS/Docker Development Environment Setup

This document lays out the steps to get a Windows development evironment set up. We will install the following:

- Chocolatey
- Azure CLI for Windows
- Python for Windows
- Docker Desktop
- Windows Subsystem for Linux
- Ubuntu 18.04 LTS
- Azure CLI for Linux

## Chocolatey

Chocolatey is like apy-get for Windows. It helps you install the latest versions of Windows apps from the command line. We'll use it to reduce the time it takes to get setup and running.

Open an administrative Powershell command line and run:

`Set-ExecutionPolicy Bypass -Scope Process -Force; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))`

## Azure CLI for Windows

From that same administrative Powershell session run:

`Invoke-WebRequest -Uri https://aka.ms/installazurecliwindows -OutFile .\AzureCLI.msi; Start-Process msiexec.exe -Wait -ArgumentList '/I AzureCLI.msi /quiet'`

## Python for Windows

This gives you the latest bits for the Python runtime (currently 3.7).  From that same administrative Powershell session run:

`choco install python`

## Docker Desktop

From that same administrative Powershell session run:

`choco install docker-desktop`

## Windows Subsystem for Linux

First, enable WSL using that same Powershell command prompt:

`Enable-WindowsOptionalFeature -Online -FeatureName Microsoft-Windows-Subsystem-Linux`

You'll need to reboot.

Now, go to the Microsoft Store and [download Ubuntu 18.04 LTS](https://www.microsoft.com/en-us/p/ubuntu-1804-lts/9n9tngvndl3q?activetab=pivot:overviewtab). After the download completes, **you'll need to launch the 'app' one time to complete the Ubuntu setup**.

After that, you can load a command prompt and type `bash` to get a legit bash shell.

## Install Azure CLI for Linux

As stated above, open a command prompt, type `bash`, and then enter the following command (note: the curl option "-sL" is case sensitive):

`curl -sL https://aka.ms/InstallAzureCLIDeb | sudo bash`

Type `exit` to return to the Windows command prompt.

## Enable Kubernetes

Open Docker Desktop and click the gear icon at the top to go to settings.

Click `Kubernetes` on the left and check the box for `Enable Kubernetes`.
