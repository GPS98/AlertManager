package main

import (
	"testing"
)

func TestSaveAlert(t *testing.T) {
	// Initialize the InMemory storage
	storage, _ := InMemoryStorage()

	// Create a new alert
	alert := Alert{AlertType: "test", Time: 123456}

	// Save the alert
	err := storage.SaveAlert(alert)

	// Check if there was an error
	if err != nil {
		t.Errorf("Error saving alert: %s", err)
	}
}

func TestGetAlertByType(t *testing.T) {
	// Initialize the InMemory storage
	storage, _ := InMemoryStorage()

	// setup
	alert1 := Alert{AlertType: "test", Time: 123456}
	alert2 := Alert{AlertType: "test", Time: 123457}
	alert3 := Alert{AlertType: "test", Time: 123458}

	// Save the alert
	storage.SaveAlert(alert1)
	storage.SaveAlert(alert2)
	storage.SaveAlert(alert3)

	// Get the alert by type
	alerts, err := storage.GetAlertByType("test", 0, 2)

	// Check if there was an error
	if err != nil {
		t.Errorf("Error getting alert by type: %s", err)
	}

	// Check if the returned alerts are correct
	if len(alerts) != 2 || alerts[0].AlertType != "test" {
		t.Errorf("Incorrect alerts returned: %v", alerts)
	}
}

func TestGetAlertByTypeAndInRange(t *testing.T) {
	// Initialize the InMemory storage
	storage, _ := InMemoryStorage()

	// Create a new alert
	alert := Alert{AlertType: "test", Time: 123456}
	alert1 := Alert{AlertType: "test", Time: 123459}
	alert2 := Alert{AlertType: "test", Time: 123457}
	alert3 := Alert{AlertType: "test", Time: 123458}

	// Save the alert
	storage.SaveAlert(alert)
	storage.SaveAlert(alert1)
	storage.SaveAlert(alert2)
	storage.SaveAlert(alert3)

	// Get the alert by type and in range
	alerts, err := storage.GetAlertByTypeAndInRange("test", 123450, 123460, 0, 10)

	// Check if there was an error
	if err != nil {
		t.Errorf("Error getting alert by type and in range: %s", err)
	}

	// Check if the returned alerts are correct
	if len(alerts) != 4 || alerts[0].AlertType != "test" || alerts[0].Time != 123456 {
		t.Errorf("Incorrect alerts returned: %v", alerts)
	}
}
