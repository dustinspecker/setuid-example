# setuid-example

## Usage

1. Clone this repository and navigate to root of this project
1. Run `go build main.go`
1. Run `./main`
   - You'll be shown a sad face and error message
1. Run `sudo ./main` to verify everything works.
1. Run `sudo chown root:root ./main`
1. Run `sudo chmod u+s ./main`
1. Run `./main` and it will run as root now
