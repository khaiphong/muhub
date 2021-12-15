/* 
 Copyright (c) 2017, 2018, 2019, 2020, 2021 KhaiPhong

 Licensed under the Apache License, Version 2.0 (the "License");

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  
SPAN Network and synchronization among (1) immediated relevant PersonaDB in Mu, (2)
its data in OmHub, (3) online versioning and offsite backup.
*/

package main

import "fmt"	
	
func main() {
  fmt.Println("Hello, 世界")
}

/*
  (1) HTTPS server pool from PersonaDB services: assigned server from the pool,

  (2) Distributed DB among user DBs in MuHub and OmHub
func Muhub() { //to enable Mu <=> MuHub
  fmt.Printf("Get connection to the user private MuHub")
}
func Omhub() { //to enable Mu <=> OmHub
  fmt.Printf("Get connection to the user data store at selected OmHub data centre")
}

func Bees() {
  fmt.Printf("Business Economics Engineering System")
}

func Prometheus() {
  fmt.Printf("Prometheus on Badger")
}

func GitOp() {
  fmt.Printf("GitOp on Badger")
}

*/
