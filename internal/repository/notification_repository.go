package repository

import "learn-api/internal/entity"

type NotificationRepository interface {
	Create(Notification *entity.Notifications) error
	GetByUserID(userID int32) ([]*entity.Notifications, error)
	MarkAsRead(id int32) error
}
