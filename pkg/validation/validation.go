package validation

import (
	"errors"
	"strings"
)

var (
	ErrEmptyTitle   = errors.New("title cannot be empty")
	ErrEmptyContent = errors.New("content cannot be empty")
	ErrTitleTooLong = errors.New("title cannot exceed 255 characters")
	ErrInvalidUser  = errors.New("invalid user ID")
	ErrPostNotFound = errors.New("post not found")
)

func ValidatePostTitle(title string) error {
	if strings.TrimSpace(title) == "" {
		return ErrEmptyTitle
	}
	if len(title) > 255 {
		return ErrTitleTooLong
	}
	return nil
}

func ValidatePostContent(content string) error {
	if strings.TrimSpace(content) == "" {
		return ErrEmptyContent
	}
	return nil
}

func ValidateUserID(userID int32) error {
	if userID <= 0 {
		return ErrInvalidUser
	}
	return nil
}
