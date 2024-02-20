package main

import (
	"cards/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestNewDeck(t *testing.T) {
	d := new_deck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected: Ace of Spades | Found: %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected: Four of Clubs | Found: %v", d[len(d)-1])
	}
}

func TestDeckSaveToFileTwo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFileWriter := mocks.NewMockFileWriter(ctrl)

	// Set expectations on the mock
	mockFileWriter.EXPECT().WriteFile(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	// Create an instance of your deck and call the method you want to test
	d := new_deck()
	err := d.saveToFileTwo(mockFileWriter, "my_cards.txt")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
