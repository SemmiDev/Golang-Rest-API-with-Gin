package main

func main() {
	router := routes()
	router.Run(":9000")
}