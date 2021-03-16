# Quasar Fire API

**Developed by** Juan Pablo Franco Berrones </br>
**Creation Date:** 16/03/2020  </br>

***

## Summary

REST API for clients CRU operations

## Installation

1. ### Install Go

First you need golang to be installed, for that you can follow the [download and install tutorial](https://golang.org/doc/install). Make sure that you add to your PATH the bin file.(step 2 if you are Linux user).

2. ### Set ssh key for Github
Next you need to set your ssh key because go modules need it in turns to download the necessary dependencies. [This tutorial](https://docs.github.com/es/github/authenticating-to-github/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent) teach you how to do that

3. ### Install docker
If you want to use level3 functions you need to create a mariadb database for that. One way is using a mariadb docker image to run on your computer and connect the app to it. For that you need to install docker with this [tutorial](https://docs.docker.com/engine/install/ubuntu/)

4. ### Download and run mariadb
After installing docker run this command on your terminal:
```bash
    sudo docker run -e MYSQL_ROOT_PASSWORD=test --name mymariadb -d mariadb
```
5. ### Get mariadb host
Exec the next command and search for the IPAddress of the container:
```bash
    sudo docker inspect mymariadb
```
Once you get the IP address copy and paste it onthe .env file of the project into the enviroment variables called MARIADB_HOST

6. ### Create Database
To create the needed database you need to enter into the container running, exec:
```bash
    sudo docker exec -it mymariadb
```
When asking you the password, enter "test". After that you can create a database running:
```bash
    create database quasarfire;
```
7. ### Export env variables
On the root of the project run:
```bash
   source ./.env
```
to set all the enviroment variables.

8. ### run go modules an run the project
Now you can run the project:
```bash
    go mod tidy
   go run main.go
```
