<h1 style="text-align: center;">
	Simple Go & Docker dev stack
</h1>

Run the stack

    docker-compose run --rm app

Run tests

    docker-compose run --rm app go test

Run tests with benchmarks

    docker-compose run --rm app go test -bench .

