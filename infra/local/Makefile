default:
	docker-compose up

delete_all_data:
	docker-compose down
	docker-compose rm -f
	docker volume rm -f local_postgres_data

.PHONY: delete_all_data