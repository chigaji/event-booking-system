#!/bin/bash
cd cmd/api-gateway
go mod init github.com/chigaji/event-booking-system/cmd/api-gateway

cd ../booking-service
go mod init github.com/chigaji/event-booking-system/cmd/booking-service

cd ../event-service
go mod init github.com/chigaji/event-booking-system/cmd/event-service

cd ../notification-service
go mod init github.com/chigaji/event-booking-system/cmd/notification-service
