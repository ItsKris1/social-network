package utils

import uuid "github.com/satori/go.uuid"

// Create unique ID
func UniqueId() string {
	return uuid.NewV4().String()
}
