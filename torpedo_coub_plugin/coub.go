package torpedo_coub_plugin

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"net/url"

	common "github.com/tb0hdan/torpedo_common"
	"github.com/tb0hdan/torpedo_registry"
)

const (
	CoubAPIBase         = "https://coub.com/api/v2"
	CoubTimelineExplore = CoubAPIBase + "/timeline/explore/"
	CoubsSearch         = CoubAPIBase + "/search/coubs"
	TagsSearch          = CoubAPIBase + "/timeline/tag/"
)

type Coub struct {
	Flag                   string `json:"flag"`                      //: null,
	Abuses                 string `json:"abuses"`                    //: null,
	RecoubsByUsersChannels string `json:"recoubs_by_users_channels"` //: null,
	Recoub                 string `json:"recoub"`                    //": null,
	Like                   string `json:"like"`                      //": null,
	InMyBest2015           bool   `json:"in_my_best2015"`            //": false,
	ID                     int    `json:"id"`                        //: 54160735,
	Type                   string `json:"type"`                      //: "Coub::Simple",
	Permalink              string `json:"permalink"`                 //": "vspmy",
	Title                  string `json:"title"`                     //": "Skill",
	VisibilityType         string `json:"visibility_type"`           //": "public",
	OriginalVisibilityType string `json:"original_visibility_type"`  //": "public",
	ChannelID              int    `json:"channel_id"`                //": 2878646,
	CreatedAt              string `json:"created_at"`                //": "2017-07-10T19:14:50Z",
	UpdatedAt              string `json:"updated_at"`                //": "2017-07-12T13:05:44Z",
	IsDone                 bool   `json:"is_done"`                   //": true,
	ViewsCount             int    `json:"views_count"`               //": 36363,
	Cotd                   bool   `json:"cotd"`                      //": true,
	CotdAt                 string `json:"cotd_at"`                   //": "2017-07-11",
	Published              bool   `json:"published"`                 //": true,
	PublishedAt            string `json:"published_at"`              //": "2017-07-10T19:14:50Z",
	ViewsIncreaseCount     int    `json:"views_increase_count"`      //": 36245,
	SharesCount            int    `json:"shares_count"`              //": 0,
	Reversed               bool   `json:"reversed"`                  //": false,
	FromEditorV2           bool   `json:"from_editor_v2"`            //": true,
	IsEditable             bool   `json:"is_editable"`               //": true,
	OriginalSound          bool   `json:"original_sound"`            //": false,
	HasSound               bool   `json:"has_sound"`                 //": false,
	RecoubTo               string `json:"recoub_to"`                 //": null,
	AgeRestricted          bool   `json:"age_restricted"`            //": false,
	AgeRestrictedByAdmin   bool   `json:"age_restricted_by_admin"`   //": false,
	NotSafeForWork         bool   `json:"not_safe_for_work"`         //": false,
	AllowReuse             bool   `json:"allow_reuse"`               //": false,
	DontCrop               bool   `json:"dont_crop"`                 //": false,
	Banned                 bool   `json:"banned"`                    //": false,
	GlobalSafe             bool   `json:"global_safe"`               //": true,
	AudioFileURL           string `json:"audio_file_url"`            //": http://s4.storage.akamai.coub.com/get/b180/p/audio_track/cw_normalized_copy/4bec4efbc4a/595a2b3422db19de23bf0/low_1499719493_11wmes3_normalized_1499714009_audio.mp3,
	//ExternalDownload       bool    `json:"external_download"`         //": false,
	ExternalDownload     interface{} `json:"external_download"`
	Application          string      `json:"application"`             //": null,
	Pircture             string      `json:"picture"`                 //": http://s3.storage.akamai.coub.com/get/b180/p/coub/simple/cw_image/dd6b4d21fc3/55f3d520f9317ed7c62ee/med_1499719523_00032.jpg,
	TimelinePicture      string      `json:"timeline_picture"`        //": http://s1.storage.akamai.coub.com/get/b179/p/coub/simple/cw_timeline_pic/8629664d22c/5022b5e6b59bf841e1b94/ios_large_1499719523_image.jpg,
	SmallPicture         string      `json:"small_picture"`           //": http://s.storage.akamai.coub.com/get/b180/p/coub/simple/cw_image/dd6b4d21fc3/55f3d520f9317ed7c62ee/ios_mosaic_1499719523_00032.jpg,
	SharingPicture       string      `json:"sharing_picture"`         //": null,
	PercentDone          int         `json:"percent_done"`            //": 100,
	RecoubsCount         int         `json:"recoubs_count"`           //": 341,
	RemixesCount         int         `json:"remixes_count"`           //": 0,
	LikesCount           int         `json:"likes_count"`             //": 1095,
	RawVideoID           int64       `json:"raw_video_id"`            //": 9924768,
	UploadedByIOSApp     bool        `json:"uploaded_by_ios_app"`     //": false,
	UploadedByAndroidApp bool        `json:"uploaded_by_android_app"` //": false,
	RawVideoThumbnailURL string      `json:"raw_video_thumbnail_url"` //": http://s1.storage.akamai.coub.com/get/b189/p/raw_video/cw_image/1212c2fb09f/0224a9e6b11d2b54cf80f/coub_media_1499719118_image.jpg,
	RawVideoTitle        string      `json:"raw_video_title"`         //": "skills.mp4",
	VideoBlockBanned     bool        `json:"video_block_banned"`      //": false,
	Duration             float64     `json:"duration"`                //": 10,
	PromoWinner          bool        `json:"promo_winner"`            //": false,
	PromoWinnerRecoubers interface{} `json:"promo_winner_recoubers"`  //": null,
	TrackingPixelURL     string      `json:"tracking_pixel_url"`      //": null,
	PromoHint            string      `json:"promo_hint"`              //": null,
	BeelineBest2014      string      `json:"beeline_best_2014"`       //": null,
	FromWebEditor        bool        `json:"from_web_editor"`         //": true,
	NormalizeSound       bool        `json:"normalize_sound"`         //": true,
	LoopsCount           int         `json:"loops_count"`             //": 0,
	TotalViewsDuration   int         `json:"total_views_duration"`    //": null,
	Best2015Addable      bool        `json:"best2015_addable"`        //": false,
	AhmadPromo           string      `json:"ahmad_promo"`             //": null,
	PromoData            string      `json:"promo_data"`              //": null,
	AudioCopyrightClam   string      `json:"audio_copyright_claim"`   //": null,
	AdsSafe              string      `json:"ads_safe"`                //": null,
	PositionOnPage       int         `json:"position_on_page"`        //": 1
	FileVersions         struct {
		Web struct {
			Template string   `json:"template"`
			Types    []string `json:"types"`
			Versions []string `json:"versions"`
		} `json:"web"`
		HTML5 struct {
			Video struct {
				High struct {
					URL  string `json:"url"`
					Size int    `json:"size"`
				} `json:"high"`
				Med struct {
					URL  string `json:"url"`
					Size int    `json:"size"`
				} `json:"med"`
			} `json:"video"`
			Audio struct {
				High struct {
					URL  string `json:"url"`
					Size int    `json:"size"`
				} `json:"high"`
				Med struct {
					URL  string `json:"url"`
					Size int    `json:"size"`
				} `json:"med"`
			} `json:"audio"`
		} `json:"html5"`
		Mobile struct {
			GIFV        string   `json:"gifv"`
			LoopedAudio string   `json:"looped_audio"`
			Audio       []string `json:"audio"`
		} `json:"mobile"`
	} `json:"file_versions"`
	AudioVersions struct {
		Template string   `json:"template"`
		Versions []string `json:"versions"`
		Chunks   struct {
			Template string   `json:"template"`
			Versions []string `json:"versions"`
			Chunks   []int    `json:"chunks"`
		} `json:"chunks"`
	} `json:"audio_versions"`
	FLVAudioVersions struct {
		Template string   `json:"template"`
		Versions []string `json:"versions"`
		Chunks   struct {
			Template string   `json:"template"`
			Versions []string `json:"versions"`
			Chunks   []int    `json:"chunks"`
		} `json:"chunks"`
	} `json:"flv_audio_versions"`
	ImageVersions struct {
		Template string   `json:"template"`
		Versions []string `json:"versions"`
	} `json:"image_versions"`
	FirstFrameVersions struct {
		Template string   `json:"template"`
		Versions []string `json:"versions"`
	} `json:"first_frame_versions"`
	Dimensions struct {
		Big []int `json:"big"`
		Med []int `json:"med"`
	} `json:"dimensions"`
	SiteWH      []int `json:"site_w_h"`
	PageWH      []int `json:"page_w_h"`
	SiteWHSmall []int `json:"site_w_h_small"`
	Channel     struct {
		ID              int    `json:"id"`
		Permalink       string `json:"permalink"`
		Title           string `json:"title"`
		Description     string `json:"description"`
		FollowersCouint int    `json:"followers_count"`
		FollowingCount  int    `json:"following_count"`
		AvatarVersions  struct {
			Template string   `json:"template"`
			Versions []string `json:"versions"`
		} `json:"avatar_versions"`
	} `json:"channel"`
	Tags []struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
		Value string `json:"value"`
	} `json:"tags"`
	Categories []struct {
		ID        int    `json:"id"`
		Title     string `json:"title"`
		Permalink string `json:"permalink"`
		Visible   bool   `json:"visible"`
	} `json:"categories"`
	MediaBlocks struct {
		UploadedRawVideos []interface{} `json:"uploaded_raw_videos"`
		ExternalRawVideos []interface{} `json:"external_raw_videos"`
		RemixedFromCoubs  []interface{} `json:"remixed_from_coubs"`
	} `json:"media_blocks"`
	EditorialInfo struct {
	} `json:"editorial_info"`
	Suggestions []string `json:"suggestions"`
}

type CoubResponse struct {
	Page       int    `json:"page"`
	TotalPages int    `json:"total_pages"`
	PerPage    int    `json:"per_page"`
	Coubs      []Coub `json:"coubs"`
}

type CoubClient struct {
	cu     *common.Utils
	logger *log.Logger
}

func (cc *CoubClient) ParseBigCoub(request_url string) (result string) {
	coubresponse := CoubResponse{}

	err := cc.cu.GetURLUnmarshal(request_url, &coubresponse)
	if err != nil {
		result = fmt.Sprintf("%+v\n", err)

	}
	coub := coubresponse.Coubs[0]
	template := coub.FileVersions.Web.Template
	vtype := ""
	vsize := ""
	for _, name := range coub.FileVersions.Web.Types {
		if name == "mp4" {
			vtype = name
			break
		}
	}
	for _, name := range coub.FileVersions.Web.Versions {
		if name == "big" {
			vsize = name
			break
		}
	}
	if vtype != "" && vsize != "" {
		result = strings.Replace(template, `%{type}`, vtype, -1)
		result = strings.Replace(result, `%{version}`, vsize, -1)
	} else {
		result = "No usable formats available"
	}
	return
}

func (cc *CoubClient) GetCoub(command string) (result string, err error) {
	var request_url string
	command = strings.TrimSpace(command)
	switch strings.Split(command, " ")[0] {
	case "random":
		request_url = CoubTimelineExplore + "random?page=1&per_page=1"
	case "newest":
		request_url = CoubTimelineExplore + "newest?page=1&per_page=1"
	case "dujour":
		request_url = CoubTimelineExplore + "coub_of_the_day?page=1&per_page=1"
	case "tag":
		if len(strings.Split(command, " ")) >= 2 {
			query := strings.TrimSpace(strings.TrimPrefix(command, "tag"))
			request_url = TagsSearch + fmt.Sprintf("%s?order_by=newest_popular&page=1", url.QueryEscape(query))
		} else {
			result = fmt.Sprintf("tag requires string argument")
		}
	case "search":
		if len(strings.Split(command, " ")) >= 2 {
			query := strings.TrimSpace(strings.TrimPrefix(command, "search"))
			request_url = CoubsSearch + fmt.Sprintf("?q=%s&order_by=likes_count", url.QueryEscape(query))
		} else {
			result = "search requires string argument"
		}
	}
	if request_url == "" {
		if result != "" {
			err = errors.New(result)
		} else {
			result = "Usage: coub [random|newest|dujour|search|tag]"
		}
	} else {
		cc.logger.Printf(request_url)
		result = cc.ParseBigCoub(request_url)
	}
	return
}

func CoubProcessMessage(api *torpedo_registry.BotAPI, channel interface{}, incoming_message string) {
	var message string
	cu := &common.Utils{}
	logger := cu.NewLog("coub-process-message")
	_, command, _ := common.GetRequestedFeature(incoming_message)
	client := &CoubClient{cu: cu, logger: logger}
	message, err := client.GetCoub(command)
	if err != nil {
		message = fmt.Sprintf("An error occured while processing Coub request: %+v\n", err)
	}
	api.Bot.PostMessage(channel, message, api)
	return
}

func init() {
	torpedo_registry.Config.RegisterHelpAndHandler("coub", "Get Coub.", CoubProcessMessage)
}
