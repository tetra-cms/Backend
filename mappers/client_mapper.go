package mapper

import (
	"tetra-server/dto"
	"tetra-server/models"
)

func ToClientResponse(client *models.Client) dto.ClientResponse {
	return dto.ClientResponse{
		ID:      client.ID,
		UserID:  client.UserID,
		FCS:     client.FCS,
		City:    client.City,
		Address: client.Address,
		Phone:   client.Phone,
	}
}

func ToClientResponses(clients []models.Client) []dto.ClientResponse {
	result := make([]dto.ClientResponse, len(clients))

	for i := range clients {
		result[i] = ToClientResponse(&clients[i])
	}

	return result
}
