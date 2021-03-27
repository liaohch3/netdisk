package cache

import "time"

func UpdateSessionMap(session string, userID int64, maxAge time.Duration) error {
	return client.Set(session, userID, maxAge).Err()
}

func GetUserIDFromSession(session string) (int64, error) {
	return client.Get(session).Int64()
}
