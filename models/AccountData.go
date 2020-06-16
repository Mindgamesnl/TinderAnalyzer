package models

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"time"
)

type AccountData struct {
	Photos               []string `json:"Photos"`

	User struct {
		ActiveTime   time.Time `json:"active_time"`
		AgeFilterMax int       `json:"age_filter_max"`
		AgeFilterMin int       `json:"age_filter_min"`
		Bio          string    `json:"bio"`
		BirthDate    time.Time `json:"birth_date"`

		City         struct {
			Name   string `json:"name"`
			Region string `json:"region"`
			Coords struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			} `json:"coords"`
		} `json:"city"`

		CreateDate   time.Time `json:"create_date"`
		Email        string    `json:"email"`
		FullName     string    `json:"full_name"`
		Gender       string    `json:"gender"`
		GenderFilter string    `json:"gender_filter"`

		Instagram    struct {
			CompletedInitialFetch bool      `json:"completed_initial_fetch"`
			LastFetchTime         time.Time `json:"last_fetch_time"`
			MediaCount            int       `json:"media_count"`
			Photos                []struct {
				ID        string `json:"id"`
				Image     string `json:"image"`
				Thumbnail string `json:"thumbnail"`
				Ts        string `json:"ts"`
			} `json:"photos"`
			Username string `json:"username"`
		} `json:"instagram"`

		InterestedIn string `json:"interested_in"`
		IPAddress    string `json:"ip_address"`

		Jobs         []struct {
			Company struct {
				Displayed bool   `json:"displayed"`
				Name      string `json:"name"`
			} `json:"company"`
			Title struct {
				Displayed bool   `json:"displayed"`
				Name      string `json:"name"`
			} `json:"title"`
		} `json:"jobs"`

		Name string `json:"name"`

		Pos  struct {
			At  string  `json:"at"`
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"pos"`

		ClientRegistrationInfo struct {
			Platform   string `json:"platform"`
			AppVersion int    `json:"app_version"`
		} `json:"client_registration_info"`

		PhoneID   string        `json:"phone_id"`
		Education string        `json:"education"`
	} `json:"User"`

	Messages []struct {
		MatchID  string `json:"match_id"`
		Messages []struct {
			To       int    `json:"to"`
			From     string `json:"from"`
			Message  string `json:"message"`
			SentDate string `json:"sent_date"`
		} `json:"messages"`
	} `json:"Messages"`

	Spotify struct {
		SpotifyConnected     bool      `json:"spotify_connected"`
		SpotifyUsername      string    `json:"spotify_username"`
		SpotifyUserType      string    `json:"spotify_user_type"`
		SpotifyLastUpdatedAt time.Time `json:"spotify_last_updated_at"`
		SpotifyThemeTrack    struct {
			ID    string `json:"id"`
			Name  string `json:"name"`
			Album struct {
				ID     string `json:"id"`
				Name   string `json:"name"`
				Images []struct {
					Height int    `json:"height"`
					Width  int    `json:"width"`
					URL    string `json:"url"`
				} `json:"images"`
			} `json:"album"`
			Artists []struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"artists"`
			PreviewURL string `json:"preview_url"`
			URI        string `json:"uri"`
		} `json:"spotify_theme_track"`

		SpotifyTopArtists []struct {
			ID         string `json:"id"`
			Name       string `json:"name"`
			Popularity int    `json:"popularity"`
			Selected   bool   `json:"selected"`
			TopTrack   struct {
				ID    string `json:"id"`
				Name  string `json:"name"`
				Album struct {
					ID     string `json:"id"`
					Name   string `json:"name"`
					Images []struct {
						Height int    `json:"height"`
						Width  int    `json:"width"`
						URL    string `json:"url"`
					} `json:"images"`
				} `json:"album"`
				Artists []struct {
					ID   string `json:"id"`
					Name string `json:"name"`
				} `json:"artists"`
				PreviewURL string `json:"preview_url"`
				URI        string `json:"uri"`
			} `json:"top_track"`
		} `json:"spotify_top_artists"`
	} `json:"Spotify"`

	Usage struct {
		// the following fields are statistics linked to dates
		// format: YYYY-MM-DD
		AppOpens map[string]int `json:"app_opens"`
		SwipesLikes map[string]int `json:"swipes_likes"`
		SwipesPasses map[string]int `json:"swipes_passes"`
		Matches map[string]int `json:"matches"`
		MessagesSent map[string]int `json:"messages_sent"`
		MessagesReceived map[string]int `json:"messages_received"`
		AdvertisingID map[string]string `json:"advertising_id"`
		Idfa map[string]string `json:"idfa"`
	} `json:"Usage"`
}

type Day struct {
	Date   			 string
	AppOpens         int
	Likes            int
	Passes           int
	Matches          int
	MessagesSent     int
	MessagesReceived int
}

// parse days
func (account AccountData) GetDays() []Day {
	var output []Day
	for date := range account.Usage.AppOpens {
		opens := account.Usage.AppOpens[date]
		likes := account.Usage.SwipesLikes[date]
		passes := account.Usage.SwipesPasses[date]
		matches := account.Usage.Matches[date]
		sentMessages := account.Usage.MessagesSent[date]
		receivedMessages := account.Usage.MessagesReceived[date]

		day := Day{
			Date: date,
			AppOpens: opens,
			Likes: likes,
			Passes: passes,
			Matches: matches,
			MessagesSent: sentMessages,
			MessagesReceived: receivedMessages,
		}

		output = append(output, day)
	}
	return output
}

func LoadAccountData(filename string) AccountData {
	jsonFile, err := os.Open("data.json")
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var accountData AccountData
	json.Unmarshal(byteValue, &accountData)
	return accountData
}