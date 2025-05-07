package enrichment_service

import (
	"service/pkg/models/domain/config"
	"service/pkg/models/domain/user"
	"testing"
)

func setupEnrichmentService() EnrichmentService {
	configEnrich := config.Enrich{
		UrlAge:    "https://api.agify.io/",
		UrlGender: "https://api.genderize.io/",
		UrlNation: "https://api.nationalize.io/",
	}
	var enrichmentService EnrichmentService = NewEnrichmentServiceManager(configEnrich)
	return enrichmentService
}

func TestEnrichmentServiceManager_EnrichmentUser(t *testing.T) {
	enrichmentService := setupEnrichmentService()

	tests := []struct {
		name    string
		user    *user.User
		wantErr bool
	}{
		{
			name: "Успешный тест",
			user: &user.User{
				Name: "Joe",
			},
			wantErr: false,
		},
		{
			name: "Неверное имя",
			user: &user.User{
				Name: "unknown",
			},
			wantErr: true,
		},
		{
			name: "Успешный тест",
			user: &user.User{
				Name: "Donald",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := enrichmentService.EnrichmentUser(tt.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("EnrichmentUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
