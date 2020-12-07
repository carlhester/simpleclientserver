package main

func main() {
	cfg := config{
		ip:   "0.0.0.0",
		port: 8123,
	}

	srv := newSimpleServer(cfg)
	srv.run()
}
