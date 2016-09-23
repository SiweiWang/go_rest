package models

type (
  // User represents the structure of our resources 
  // We also use struct tags for each field
  User struct {
    Name string `json:"name"`
    Gender string `json:"gender"`
    Age int `json:"age"`
    Id  string `json:"id"`
  }
)