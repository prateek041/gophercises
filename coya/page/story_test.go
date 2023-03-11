package page

import (
	"testing"
)

func TestReadAll(t *testing.T) {
	got := ReadAll()
	t.Run("Returned data not empty", func(t *testing.T) {
		if len(got) == 0 {
			t.Errorf("Expected non-empty data, got empty")
		}
	})

}

//func TestCreateStoryMap(t *testing.T) {
//	testJson := `{"a":"apple", "b":"ball"}`
//	got := createStoryMap([]byte(testJson))
//	t.Run("Returned data not empty", func(t *testing.T) {
//		if len(got) == 0 {
//			t.Errorf("Expected non empty data, got empty")
//		}
//	})
//	t.Run("Testing individual key access", func(t *testing.T) {
//		value, found := got["a"]
//		if !found {
//			t.Errorf("Error in accesing the value of map")
//		}
//
//		if value != "apple" {
//			t.Errorf("Expected %s got %s \n", "apple", value)
//		}
//	})
//}

func TestRunGame(t *testing.T) {
	StartGame()
}
