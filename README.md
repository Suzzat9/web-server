# A simple web server in Go

The original idea for this project is from [Coding Challenges](https://codingchallenges.fyi/challenges/challenge-webserver).

In this project I built a HTTP web server in [Go](https://go.dev). I used the [net package](https://pkg.go.dev/net) for socket programming, and made use of [WaitGroups](https://gobyexample.com/waitgroups) in Go to run concurrent threads to test the server. 

To see the web server in action, clone this repo and navigate to the web-server directory. Then in a new terminal, run `sudo go run server.go` to start the web server. 

You can then run `sudo go run client-test.go` in a separate terminal to run multiple requests on the server, or curl the server by running `curl http://localhost/` in a separate terminal. 

