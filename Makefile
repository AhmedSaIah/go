postgres:
	@docker run --name postgres-14 -p 5000:5000 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=01000 -d postgres:14-alpine

createDB:
	@docker exec -it exec postgres-14 createdb --username=root --owner=root bank

dropDB:
	@docker exec -it postgres-14 dropdb bank
<<<<<<< Updated upstream

.PHONY: format
format:
	@npm run format:prettier
=======
>>>>>>> Stashed changes
