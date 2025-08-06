start:
	docker-compose up -d

stop:
	docker-compose down

logs:
	docker-compose logs -f app

test:
	go test -v ./...

migrate:
	docker-compose exec postgres psql -U user -d url_shortener -c "SELECT * FROM urls;"

clean:
	docker-compose down -v
	rm -f url-shortener