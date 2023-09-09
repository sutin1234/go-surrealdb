db_start:
	surreal start --user root --pass root --bind 0.0.0.0:8000 file:mydatabase.db
run:
	air -c .air.toml
build:
	go build