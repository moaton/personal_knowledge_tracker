package usecases

import (
	"personal_knowledge_tracker/internal/dto"
	"personal_knowledge_tracker/internal/entity"
)

func convertResourceToEntity(r *dto.Resource) *entity.Resource {
	return &entity.Resource{
		Title:     r.Title,
		Type:      r.Type,
		Content:   r.Content,
		Tags:      r.Tags,
		Metadata:  r.Metadata,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}

func convertResourceToDTO(r *entity.Resource) *dto.Resource {
	return &dto.Resource{
		ID:        r.ID.Hex(),
		Title:     r.Title,
		Type:      r.Type,
		Content:   r.Content,
		Tags:      r.Tags,
		Metadata:  r.Metadata,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}

func convertResourcesToDTO(r []*entity.Resource) []*dto.Resource {
	out := make([]*dto.Resource, 0, len(r))
	for i := range r {
		out = append(out, convertResourceToDTO(r[i]))
	}

	return out
}
