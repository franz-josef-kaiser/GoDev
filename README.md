<h1 align="center">
	Simple Go & Docker dev stack
</h1>

<p align="center">
	Basic intro to _Go_ [with examples](https://gobyexample.com/).
</p>

Run the stack

    docker-compose run --rm app

Run tests

    docker-compose run --rm app go test

Run tests with benchmarks

    docker-compose run --rm app go test -bench .

