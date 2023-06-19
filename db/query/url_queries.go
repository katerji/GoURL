package query

const FetchUserURLs = "SELECT id, user_id, short_url, original_url FROM url WHERE user_id = ?"
const InsertURL = "INSERT INTO url (user_id, short_url, original_url) VALUES (?, ?, ?)"
const DeleteURL = "DELETE FROM url WHERE id = ?"
const FetchURL = "SELECT id, user_id, short_url, original_url FROM url WHERE short_url = ?"
