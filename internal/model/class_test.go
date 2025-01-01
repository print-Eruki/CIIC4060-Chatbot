package model

import (
	"encoding/json"
	"testing"
)

func TestClassSerialization(t *testing.T) {
	class := Class{
		Cid:       1,
		Ccode:     "CIIC3015",
		Cname:     "Introduction to Programming",
		Cred:      3,
		Cdesc:     "Learn basic programming concepts.",
		Csyllabus: "http://example.com/syllabus",
		Term:      "Fall",
		Years:     "2024",
	}

	// Serialize to JSON
	data, err := json.Marshal(class)
	if err != nil {
		t.Fatalf("Failed to serialize Class to JSON: %v", err)
	}

	// Check the JSON output
	expected := `{"cid":1,"ccode":"CIIC3015","cname":"Introduction to Programming","cred":3,"cdesc":"Learn basic programming concepts.","csyllabus":"http://example.com/syllabus","term":"Fall","years":"2024"}`
	if string(data) != expected {
		t.Errorf("JSON serialization mismatch.\nExpected: %s\nGot: %s", expected, string(data))
	}

	// Deserialize back to a Class struct
	var deserializedClass Class
	err = json.Unmarshal(data, &deserializedClass)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON to Class: %v", err)
	}

	// Verify the deserialized struct matches the original
	if class != deserializedClass {
		t.Errorf("Struct mismatch after deserialization.\nExpected: %+v\nGot: %+v", class, deserializedClass)
	}
}

func TestCreateClassFromJSON(t *testing.T) {
	incomingJSON := `{
		"cid": 1,
		"ccode": "CIIC3015",
		"cname": "Introduction to Programming",
		"cred": 3,
		"cdesc": "Learn basic programming concepts.",
		"csyllabus": "http://example.com/syllabus",
		"term": "Fall",
		"years": "2024"
	}`

	// Deserialize into a Class struct
	var class Class
	err := json.Unmarshal([]byte(incomingJSON), &class)
	if err != nil {
		t.Fatalf("Failed to bind JSON to Class: %v", err)
	}

	// Expected struct
	expectedClass := Class{
		Cid:       1,
		Ccode:     "CIIC3015",
		Cname:     "Introduction to Programming",
		Cred:      3,
		Cdesc:     "Learn basic programming concepts.",
		Csyllabus: "http://example.com/syllabus",
		Term:      "Fall",
		Years:     "2024",
	}

	// Compare the result with the expected
	if class != expectedClass {
		t.Errorf("Class struct mismatch.\nExpected: %+v\nGot: %+v", expectedClass, class)
	}
}
