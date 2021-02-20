package models

//User is the structure that contains all data involoving a User
type User struct {
	UserID string `json:"userID"`

	Name string `json:"name"`

	BirthDate string `json:"birthDate"`

	Location string `json:"location"`

	Gender string `json:"gender"`

	PhoneNumber string `json:"phoneNumber`

	SuggestionPreferences struct {
		Age struct {
			Max int `json:"maxAge"`

			Min int `json:"minAge"`
		} `json:"ageRange"`

		AllowGlobal string `json:"allowGlobal"`

		MaxDistance int `json:"maxDistance"`

		GenderPreference string `json:"genderPreference"`
	} `json:"suggestionPreferences"`
}
