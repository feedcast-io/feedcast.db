test:
	go clean -testcache && go test github.com/feedcast-io/feedcast.db
	go clean -testcache && go test github.com/feedcast-io/feedcast.db/models