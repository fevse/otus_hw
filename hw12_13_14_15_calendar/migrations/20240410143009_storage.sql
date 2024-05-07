-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS events (
    id    		SERIAL PRIMARY KEY,
	title 		VARCHAR(255) NOT NULL,
	date 		DATE,
	duration 	INT,
	description TEXT,
	userid 		INT,
	reminder 	DATE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS events;
-- +goose StatementEnd