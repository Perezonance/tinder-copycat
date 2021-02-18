package models

import "mime/multipart"

type (
	//Profile provides the full structure for containing a profile's data
	Profile struct {
		ProfileID       string   `json:"profileID"`
		UserID          string   `json:"userID"`
		Bio             string   `json:"bio"`
		InterestTags    []string `json:"interestTags"`
		Occupation      string   `json:"occupation"`
		Company         string   `json:"company"`
		School          string   `json:"school"`
		LivingIn        string   `json:"livingIn"`
		DisplayAge      bool     `json:"displayAge"`
		DisplayDistance bool     `json:"displayDistance"`
	}

	/*ProfilePreview1 provides a hybrid referance and storage structure that displays the minimum
	and age, distance, work
	It is used for cutting down real-time serving of profiles to prioritize on rendering each full profile*/
	ProfilePreview1 struct {
		PreviewID string `json:"previewID"`
		Base      Base   `json:"base"`
		Album     Album  `json:"album"`

		Age      int `json:"age"`
		Distance int `json:"distance"`
	}

	//ProfilePreview2 provides a hybrid referance and storage structure that displays the minimum
	ProfilePreview2 struct {
		Name string `json:"name"`
	}

	//Base is the base info needed for all ProfilePreview types
	Base struct {
		Name string `json:"name"`
		Bio  string `json:"bio"`

		//ReferenceMetaData provides a reference struct for other profile data
		ReferenceMetaData struct {
			UserID    string `json:"userID"`
			ProfileID string `json:"profileID"`
			AlbumID   string `json:"albumID"`
		} `json:"refMetaData"`
	}

	//Album is the reference structure that contains the collection of images as well as the id for retrieval
	Album struct {
		AlbumID string           `json:"albumID"`
		Images  []multipart.File `json:"images"`
	}

	//TODO: Tech Debt -> determine if data model structure is efficient enough or if there needs to be a better solution
	// //Age is the reference structure that contains the value as well as whether it should be displayed or not
	// Age struct {
	// 	AgeValue   int  `json:"ageValue"`
	// 	DisplayAge bool `json:"diplayAge"`
	// }

	// //Distance provides the masking structure that contains the Distance value as well
	// Distance struct {
	// 	DistanceValue   int  `json:"distanceValue"`
	// 	DisplayDistance bool `json:"diplayDistance"`
	// }
)
