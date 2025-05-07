package enrichment_service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"service/pkg/models/domain/config"
	"service/pkg/models/domain/user"
)

type EnrichmentServiceManager struct {
	config config.Enrich
}

type EnrichmentService interface {
	EnrichmentUser(user *user.User) error
}

func NewEnrichmentServiceManager(configEnrich config.Enrich) *EnrichmentServiceManager {
	return &EnrichmentServiceManager{
		configEnrich,
	}
}

// EnrichmentUser - обогащение возрастом, полом, национальностью
func (e *EnrichmentServiceManager) EnrichmentUser(user *user.User) error {
	urlAge := fmt.Sprintf("%s?name=%s", e.config.UrlAge, user.Name)
	urlGender := fmt.Sprintf("%s?name=%s", e.config.UrlGender, user.Name)
	urlNation := fmt.Sprintf("%s?name=%s", e.config.UrlNation, user.Name)

	if err := e.enrichmentUserAge(user, urlAge); err != nil {
		return err
	}

	if err := e.enrichmentUserGender(user, urlGender); err != nil {
		return err
	}

	if err := e.enrichmentUserNation(user, urlNation); err != nil {
		return err
	}

	return nil
}

func (e *EnrichmentServiceManager) enrichmentUserAge(user *user.User, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("EnrichmentUserAge status code: %d", resp.StatusCode)
	}

	var data map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return err
	}

	if age, ok := data["age"]; !ok || age == nil {
		return fmt.Errorf("EnrichmentUser age not found in response")
	} else if ageFloat, ok := age.(float64); !ok {
		return fmt.Errorf("EnrichmentUser must be a integer")
	} else {
		user.Age = int(ageFloat)
	}

	return nil
}

func (e *EnrichmentServiceManager) enrichmentUserGender(user *user.User, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("EnrichmentUserGender status code: %d", resp.StatusCode)
	}

	var data map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return err
	}

	if age, ok := data["gender"]; !ok || age == nil {
		return fmt.Errorf("EnrichmentUser gender not found in response")
	} else if gender, ok := age.(string); !ok {
		return fmt.Errorf("EnrichmentUser must be a string")
	} else {
		user.Gender = gender
	}

	return nil
}

func (e *EnrichmentServiceManager) enrichmentUserNation(user *user.User, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("EnrichmentUserNation status code: %d", resp.StatusCode)
	}

	type Country struct {
		CountryID   string  `json:"country_id"`
		Probability float64 `json:"probability"`
	}

	type Response struct {
		Count   int       `json:"count"`
		Name    string    `json:"name"`
		Country []Country `json:"country"`
	}

	var response Response

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}

	// поиск максимального probability для получения национальности
	maxProbability := float64(-1 << 30)
	countryID := ""
	for _, country := range response.Country {
		if country.Probability > maxProbability {
			maxProbability = country.Probability
			countryID = country.CountryID
		}
	}

	if countryID == "" {
		return fmt.Errorf("EnrichmentUserNation not found in response")
	}

	user.Nation = countryID

	return nil
}
