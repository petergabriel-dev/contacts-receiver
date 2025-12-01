-- name: ListContacts :many
SELECT * FROM contacts;

-- name: GetContactByID :one
SELECT * FROM contacts WHERE id = ?;

-- name: CreateContact :exec
INSERT INTO contacts (name, email, message) VALUES (?, ?, ?);