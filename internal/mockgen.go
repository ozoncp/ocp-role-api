package internal

//go:generate mockgen -destination=./mocks/flusher_mock.go  -package=mocks github.com/ozoncp/ocp-role-api/internal/flusher Flusher
//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ozoncp/ocp-role-api/internal/repo Repo
//go:generate mockgen -destination=./mocks/saver_mock.go -package=mocks github.com/ozoncp/ocp-role-api/internal/saver Saver
//go:generate mockgen -destination=./mocks/ticker_mock.go -package=mocks github.com/ozoncp/ocp-role-api/internal/ticker Ticker
//go:generate mockgen -destination=./mocks/server_api_mock.go -package=mocks github.com/ozoncp/ocp-role-api/internal/api ApiServer
