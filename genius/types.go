package genius

// Artist is how Genius represents the creator of one or more songs (or other
// documents hosted on Genius). It's usually a musician or group of musicians.
type Artist struct {
	AlternateNames        []string               `json:"alternate_names"`
	APIPath               string                 `json:"api_path"`
	CurrentUserMetadata   UserMetadata           `json:"current_user_metadata"`
	Description           map[string]interface{} `json:"description"`
	DescriptionAnnotation map[string]interface{} `json:"description_annotation"`
	FacebookName          string                 `json:"facebook_name"`
	FollowersCount        int                    `json:"followers_count"`
	HeaderImageURL        string                 `json:"header_image_url"`
	ID                    int                    `json:"id"`
	ImageURL              string                 `json:"image_url"`
	InstagramName         string                 `json:"instagram_name"`
	IQ                    int                    `json:"iq"`
	IsMemeVerified        bool                   `json:"is_meme_verified"`
	IsVerified            bool                   `json:"is_verified"`
	Name                  string                 `json:"name"`
	TranslationArtist     bool                   `json:"translation_artist"`
	TwitterName           string                 `json:"twitter_name"`
	URL                   string                 `json:"url"`
	User                  User                   `json:"user"`
}

type Meta struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Object struct {
	Meta     `json:"meta"`
	Response `json:"response"`
}

type Response struct {
	Artist   *Artist     `json:"artist"`
	Hits     []SearchHit `json:"hits"`
	NextPage int         `json:"next_page"`
	Songs    []Song      `json:"songs"`
}

type SearchHit struct {
	Highlights []interface{} `json:"highlights"`
	Index      string        `json:"index"`
	Type       string        `json:"song"`
	Result     SearchResult  `json:"result"`
}

type SearchResult Song

// Song is a document hosted on Genius. It's usually music lyrics.
//
// Data for a song includes details about the document itself and information
// about all the referents that are attached to it, including the text to which
// they refer.
type Song struct {
	AnnotationCount          int    `json:"annotation_count"`
	APIPath                  string `json:"api_path"`
	FullTitle                string `json:"full_title"`
	HeaderImageThumbnailURL  string `json:"header_image_thumbnail_url"`
	HeaderImageURL           string `json:"header_image_url"`
	ID                       int    `json:"id"`
	LyricsOwnerID            int    `json:"lyrics_owner_id"`
	LyricsState              string `json:"lyrics_state"`
	Path                     string `json:"path"`
	PrimaryArtist            Artist `json:"primary_artist"`
	PyongsCount              int    `json:"pyongs_count"`
	SongArtImageThumbnailURL string `json:"song_art_image_thumbnail_url"`
	Stats                    Stat   `json:"stats"`
	Title                    string `json:"title"`
	TitleWithFeatured        string `json:"title_with_featured"`
	URL                      string `json:"url"`
}

type Stat struct {
	Hot                   bool `json:"hot"`
	UnreviewedAnnotations int  `json:"unreviewed_annotations"`
	PageViews             int  `json:"pageviews"`
}

type User struct {
	APIPath                     string       `json:"api_path"`
	Avatar                      interface{}  `json:"avatar"`
	CurrentUserMetadata         UserMetadata `json:"current_user_metadata"`
	HeaderImageURL              string       `json:"header_image_url"`
	HumanReadableRoleForDisplay string       `json:"human_readable_role_for_display"`
	ID                          int          `json:"id"`
	IQ                          int          `json:"iq"`
	Login                       string       `json:"login"`
	Name                        string       `json:"name"`
	RoleForDisplay              string       `json:"role_for_display"`
	URL                         string       `json:"url"`
}

type UserMetadata map[string]interface{}
