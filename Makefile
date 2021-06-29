include .envrc

# default will run DEVELOPMENT env
# `make run/api`
run/api:
	@go run ./api -jwt-secret=${JWT_SECRET}
