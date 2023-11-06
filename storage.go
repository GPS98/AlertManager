package main

import "log"

type Storage interface {
	SaveAlert(alert Alert) error
	GetAlertByType(alertType string, pageNo, pageSize int64) ([]Alert, error)
	GetAlertByTypeAndInRange(alertType string, fromTime, endTime, pageNo, pageSize int64) ([]Alert, error)
}

type InMemory struct {
	alerts       []Alert
	alertsByType map[string][]Alert
}

func InMemoryStorage() (*InMemory, error) {
	return &InMemory{
		alerts:       []Alert{},
		alertsByType: make(map[string][]Alert),
	}, nil
}

func (im *InMemory) SaveAlert(alert Alert) error {
	im.alerts = append(im.alerts, alert)
	im.alertsByType[alert.AlertType] = append(im.alertsByType[alert.AlertType], alert)
	return nil
}

func (im *InMemory) GetAlertByType(alertType string, pageNo int64, pageSize int64) ([]Alert, error) {
	log.Println("JSON API GetAlertByType: ", alertType)
	startIndex := pageNo * pageSize
	if startIndex >= int64(len(im.alertsByType[alertType])) {
		return []Alert{}, nil
	}
	endIndex := startIndex + pageSize
	if endIndex > int64(len(im.alertsByType[alertType])) {
		endIndex = int64(len(im.alertsByType[alertType]))
	}
	return im.alertsByType[alertType][startIndex:endIndex], nil
}

func (im *InMemory) GetAlertByTypeAndInRange(alertType string, fromTime, endTime, pageNo, pageSize int64) ([]Alert, error) {
	log.Println("JSON API GetAlertByTypeRange ")
	var alertsInRange []Alert
	for _, alert := range im.alertsByType[alertType] {
		if alert.Time >= fromTime && alert.Time <= endTime {
			alertsInRange = append(alertsInRange, alert)
		}
	}
	startIndex := pageNo * pageSize
	if startIndex >= int64(len(alertsInRange)) {
		return []Alert{}, nil
	}
	endIndex := startIndex + pageSize
	if endIndex > int64(len(alertsInRange)) {
		endIndex = int64(len(alertsInRange))
	}
	return alertsInRange[startIndex:endIndex], nil
}
