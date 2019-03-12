//menunjukan bahwa file ini berada di package mana
package main

//menujukan bahwa file ini menggunakan file yang ada di dalam package cmd
import "github.com/fahlepyrizal01/payment-app/cmd"

//fungsi utama yang pertama kali dijalankan pada saat running, perintahnya adalah
//menjalankan perintah pada package cmd
func main() {
	cmd.Execute()
}
