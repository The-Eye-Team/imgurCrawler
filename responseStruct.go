package main


type AdConfig struct {
	SafeFlags   []string      `json:"safeFlags"`
	ShowsAds    bool          `json:"showsAds"`
	UnsafeFlags []interface{} `json:"unsafeFlags"`
}

type ImageFile struct {
	Animated    bool        `json:"animated"`
	Datetime    string      `json:"datetime"`
	Description interface{} `json:"description"`
	Ext         string      `json:"ext"`
	HasSound    bool        `json:"has_sound"`
	Hash        string      `json:"hash"`
	Height      int         `json:"height"`
	Looping     bool        `json:"looping"`
	PreferVideo bool        `json:"prefer_video"`
	Size        int         `json:"size"`
	Title       string      `json:"title"`
	Width       int         `json:"width"`
}

type ImageFiles []ImageFile

type AlbumImages struct {
	Count  int        `json:"count"`
	Images ImageFiles `json:"images"`
}

type Metadata struct {
	Accent             string      `json:"accent"`
	Description        string      `json:"description"`
	IsPromoted         string      `json:"is_promoted"`
	LogoDestinationURL interface{} `json:"logo_destination_url"`
	LogoHash           interface{} `json:"logo_hash"`
	TagID              string      `json:"tag_id"`
	Title              string      `json:"title"`
}

type GalleryTag struct {
	AccountID      string `json:"account_id"`
	Animated       string `json:"animated"`
	BackgroundHash string `json:"background_hash"`
	Blocked        string `json:"blocked"`
	Display        string `json:"display"`
	Downs          string `json:"downs"`
	Hash           string `json:"hash"`
	ID             string `json:"id"`
	Image          struct {
		Animated string `json:"animated"`
	} `json:"image"`
	Images      string   `json:"images"`
	IsPromoted  string   `json:"is_promoted"`
	Metadata    Metadata `json:"metadata"`
	Nsfw        string   `json:"nsfw"`
	Score       string   `json:"score"`
	Spam        string   `json:"spam"`
	Subscribers string   `json:"subscribers"`
	Tag         string   `json:"tag"`
	TagID       string   `json:"tag_id"`
	Thumbnail   struct {
		Animated interface{} `json:"animated"`
	} `json:"thumbnail"`
	ThumbnailAnimated interface{} `json:"thumbnail_animated"`
	ThumbnailHash     interface{} `json:"thumbnail_hash"`
	Timestamp         string      `json:"timestamp"`
	Ups               string      `json:"ups"`
}

type GalleryTags []GalleryTag

type Image struct {
	AccountID        string      `json:"account_id"`
	AccountURL       string      `json:"account_url"`
	AdConfig         AdConfig    `json:"adConfig"`
	AdType           int         `json:"ad_type"`
	AdURL            string      `json:"ad_url"`
	AlbumCover       string      `json:"album_cover"`
	AlbumCoverHeight int         `json:"album_cover_height"`
	AlbumCoverWidth  int         `json:"album_cover_width"`
	AlbumImages      AlbumImages `json:"album_images"`
	AlbumLayout      string      `json:"album_layout"`
	AlbumPrivacy     string      `json:"album_privacy"`
	Animated         bool        `json:"animated"`
	Bandwidth        interface{} `json:"bandwidth"`
	CommentCount     int         `json:"comment_count"`
	Description      string      `json:"description"`
	Downs            int         `json:"downs"`
	Ext              string      `json:"ext"`
	FavoriteCount    int         `json:"favorite_count"`
	Favorited        bool        `json:"favorited"`
	GalleryTags      GalleryTags `json:"galleryTags"`
	GalleryDatetime  string      `json:"gallery_datetime"`
	Hash             string      `json:"hash"`
	Height           int         `json:"height"`
	HotDatetime      string      `json:"hot_datetime"`
	ID               int         `json:"id"`
	InGallery        bool        `json:"in_gallery"`
	IsAlbum          bool        `json:"is_album"`
	IsHot            bool        `json:"is_hot"`
	Looping          bool        `json:"looping"`
	MemeBottom       interface{} `json:"meme_bottom"`
	MemeName         interface{} `json:"meme_name"`
	MemeTop          interface{} `json:"meme_top"`
	Mimetype         interface{} `json:"mimetype"`
	Nsfw             bool        `json:"nsfw"`
	NumImages        int         `json:"num_images"`
	Pending          string      `json:"pending"`
	Platform         string      `json:"platform"`
	Points           int         `json:"points"`
	PreferVideo      bool        `json:"prefer_video"`
	Readonly         bool        `json:"readonly"`
	Reddit           interface{} `json:"reddit"`
	Score            int         `json:"score"`
	Section          string      `json:"section"`
	Size             int         `json:"size"`
	Spam             string      `json:"spam"`
	StartingScore    int         `json:"starting_score"`
	Subtype          string      `json:"subtype"`
	Tags             []string    `json:"tags"`
	Timestamp        string      `json:"timestamp"`
	Title            string      `json:"title"`
	Topic            string      `json:"topic"`
	TopicID          int         `json:"topic_id"`
	Ups              int         `json:"ups"`
	VideoHost        interface{} `json:"video_host"`
	VideoSource      interface{} `json:"video_source"`
	Views            string      `json:"views"`
	Virality         float64     `json:"virality"`
	Vote             interface{} `json:"vote"`
	Weight           int         `json:"weight"`
	Width            int         `json:"width"`
}

type Caption struct {
	AlbumCover  string      `json:"album_cover"`
	Author      string      `json:"author"`
	AuthorID    string      `json:"author_id"`
	BestComment int         `json:"best_comment"`
	BestScore   string      `json:"best_score"`
	Caption     string      `json:"caption"`
	Datetime    string      `json:"datetime"`
	Deleted     bool        `json:"deleted"`
	DeletedMeta int         `json:"deleted_meta"`
	Downs       string      `json:"downs"`
	Hash        string      `json:"hash"`
	ID          string      `json:"id"`
	OnAlbum     bool        `json:"on_album"`
	ParentID    interface{} `json:"parent_id"`
	Platform    string      `json:"platform"`
	Points      string      `json:"points"`
	Ups         string      `json:"ups"`
}

type Captions []Caption

type Response struct {
	Data struct {
		Captions Captions `json:"captions"`
		Image Image `json:"image"`
	} `json:"data"`
	Status  int  `json:"status"`
	Success bool `json:"success"`
}