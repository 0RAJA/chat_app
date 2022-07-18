package query

import (
	"context"
)

const KeyEmail = "KeyEmail"

// AddEmails adds emails to the set.
func (q *Queries) AddEmails(c context.Context, emails ...string) error {
	if len(emails) == 0 {
		return nil
	}
	data := make([]interface{}, len(emails))
	for i, email := range emails {
		data[i] = email
	}
	return q.rdb.SAdd(c, KeyEmail, data...).Err()
}

// ExistEmail checks if the email is in the set.
func (q *Queries) ExistEmail(c context.Context, email string) (bool, error) {
	return q.rdb.SIsMember(c, KeyEmail, email).Result()
}

// DeleteEmail deletes the email from the set.
func (q *Queries) DeleteEmail(c context.Context, email string) error {
	return q.rdb.SRem(c, KeyEmail, email).Err()
}

// UpdateEmail updates the email in the set.
func (q *Queries) UpdateEmail(c context.Context, oldEmail, newEmail string) error {
	if err := q.DeleteEmail(c, oldEmail); err != nil {
		return err
	}
	return q.rdb.SAdd(c, KeyEmail, newEmail).Err()
}

// ReloadEmails reloads the emails from the set.
func (q *Queries) ReloadEmails(c context.Context, emails ...string) error {
	if err := q.rdb.Del(c, KeyEmail).Err(); err != nil {
		return err
	}
	return q.AddEmails(c, emails...)
}
