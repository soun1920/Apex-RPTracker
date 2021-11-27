package main

import "time"

type Apextrackergg struct {
	Data struct {
		PlatformInfo struct {
			PlatformSlug           string      `json:"platformSlug"`
			PlatformUserId         string      `json:"platformUserId"`
			PlatformUserHandle     string      `json:"platformUserHandle"`
			PlatformUserIdentifier string      `json:"platformUserIdentifier"`
			AvatarUrl              string      `json:"avatarUrl"`
			AdditionalParameters   interface{} `json:"additionalParameters"`
		} `json:"platformInfo"`
		UserInfo struct {
			UserId          int           `json:"userId"`
			IsPremium       bool          `json:"isPremium"`
			IsVerified      bool          `json:"isVerified"`
			IsInfluencer    bool          `json:"isInfluencer"`
			IsPartner       bool          `json:"isPartner"`
			CountryCode     string        `json:"countryCode"`
			CustomAvatarUrl interface{}   `json:"customAvatarUrl"`
			CustomHeroUrl   interface{}   `json:"customHeroUrl"`
			SocialAccounts  []interface{} `json:"socialAccounts"`
			Pageviews       int           `json:"pageviews"`
			IsSuspicious    interface{}   `json:"isSuspicious"`
		} `json:"userInfo"`
		Metadata struct {
			CurrentSeason     int      `json:"currentSeason"`
			ActiveLegend      string   `json:"activeLegend"`
			ActiveLegendName  string   `json:"activeLegendName"`
			ActiveLegendStats []string `json:"activeLegendStats"`
		} `json:"metadata"`
		Segments []struct {
			Type       string `json:"type"`
			Attributes struct {
				Id string `json:"id,omitempty"`
			} `json:"attributes"`
			Metadata struct {
				Name             string `json:"name"`
				ImageUrl         string `json:"imageUrl,omitempty"`
				TallImageUrl     string `json:"tallImageUrl,omitempty"`
				BgImageUrl       string `json:"bgImageUrl,omitempty"`
				PortraitImageUrl string `json:"portraitImageUrl,omitempty"`
				LegendColor      string `json:"legendColor,omitempty"`
				IsActive         bool   `json:"isActive,omitempty"`
			} `json:"metadata"`
			ExpiryDate time.Time `json:"expiryDate"`
			Stats      struct {
				Level struct {
					Rank            interface{} `json:"rank"`
					Percentile      float64     `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        interface{} `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"level,omitempty"`
				Kills struct {
					Rank            interface{} `json:"rank"`
					Percentile      *float64    `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        interface{} `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"kills,omitempty"`
				RankScore struct {
					Rank            interface{} `json:"rank"`
					Percentile      float64     `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        interface{} `json:"category"`
					Metadata        struct {
						IconUrl  string `json:"iconUrl"`
						RankName string `json:"rankName"`
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"rankScore,omitempty"`
				ArenaRankScore struct {
					Rank            interface{} `json:"rank"`
					Percentile      interface{} `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        interface{} `json:"category"`
					Metadata        struct {
						IconUrl  string `json:"iconUrl"`
						RankName string `json:"rankName"`
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"arenaRankScore,omitempty"`
				Season5Wins struct {
					Rank            interface{} `json:"rank"`
					Percentile      float64     `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        interface{} `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"season5Wins,omitempty"`
				Season7Wins struct {
					Rank            interface{} `json:"rank"`
					Percentile      float64     `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        interface{} `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"season7Wins,omitempty"`
				Season7Kills struct {
					Rank            interface{} `json:"rank"`
					Percentile      float64     `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        interface{} `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"season7Kills,omitempty"`
				VoicesWarningsHeard struct {
					Rank            interface{} `json:"rank"`
					Percentile      float64     `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        interface{} `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"voicesWarningsHeard,omitempty"`
				GrappleTravelDistance struct {
					Rank            interface{} `json:"rank"`
					Percentile      float64     `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        interface{} `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"grappleTravelDistance,omitempty"`
				UltimateTrapsDestroyed struct {
					Rank            interface{} `json:"rank"`
					Percentile      float64     `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        interface{} `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"ultimateTrapsDestroyed,omitempty"`
				DomeDamageBlocked struct {
					Rank            interface{} `json:"rank"`
					Percentile      float64     `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        interface{} `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"domeDamageBlocked,omitempty"`
				GunShieldDamageBlocked struct {
					Rank            interface{} `json:"rank"`
					Percentile      float64     `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        interface{} `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"gunShieldDamageBlocked,omitempty"`
			} `json:"stats"`
		} `json:"segments"`
		AvailableSegments []struct {
			Type       string `json:"type"`
			Attributes struct {
			} `json:"attributes"`
			Metadata struct {
			} `json:"metadata"`
		} `json:"availableSegments"`
		ExpiryDate time.Time `json:"expiryDate"`
	} `json:"data"`
}
