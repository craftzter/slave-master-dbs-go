package psql

import (
	"context"
	"database/sql"

	"learn-api/internal/db"
	"learn-api/internal/entity"
	"learn-api/internal/repository"
)

type notificationRepo struct {
	master *sql.DB
	slave  *sql.DB
}

func NewNotificationRepoPG(master, slave *sql.DB) repository.NotificationRepository {
	return &notificationRepo{
		master: master,
		slave:  slave,
	}
}

func (r *notificationRepo) Create(notification *entity.Notifications) error {
	q := db.New(r.master)
	params := db.CreateNotificationParams{
		UserID:    notification.UserID,
		Type:      notification.Type,
		Message:   notification.Message,
		RelatedID: notification.RelatedID,
	}
	return q.CreateNotification(context.Background(), params)
}

func (r *notificationRepo) GetByUserID(userID int32) ([]*entity.Notifications, error) {
	q := db.New(r.slave)
	notifications, err := q.GetNotificationsByUserID(context.Background(), userID)
	if err != nil {
		return nil, err
	}
	var result []*entity.Notifications
	for _, n := range notifications {
		result = append(result, &entity.Notifications{
			ID:        n.ID,
			UserID:    n.UserID,
			Type:      n.Type,
			Message:   n.Message,
			RelatedID: n.RelatedID,
			IsRead:    n.IsRead,
			CreatedAt: n.CreatedAt,
		})
	}
	return result, nil
}

func (r *notificationRepo) MarkAsRead(id int32) error {
	q := db.New(r.master)
	return q.MarkNotificationAsRead(context.Background(), id)
}
