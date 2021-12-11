package main

import "fmt"

func main () {
  tampil_mahasiswa()
}

func tampil_mahasiswa(){
  var tinggi_mahasiswa = map[string]int{"Aldo":182,"Yosep":178}
  fmt.Println("Aldo :",tinggi_mahasiswa["Aldo"],"CM")
  fmt.Println("Yosep :",tinggi_mahasiswa["Yosep"],"CM")
}
