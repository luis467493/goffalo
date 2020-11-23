package models

import "github.com/gobuffalo/uuid"

type Product struct {
	ID    uuid.UUID `json:"ID"`
	Name  string    `json:"name,omitempty"`
	Color uuid.UUID    `json:"color,omitempty"`
	Type  uuid.UUID    `json:"type,omitempty"`
}

type ProductDto struct {
	ID    uuid.UUID `json:"ID"`
	Name  string    `json:"name,omitempty"`
	Color string    `json:"color,omitempty"`
	Type  string    `json:"type,omitempty"`
}

type Color struct {
	ID   uuid.UUID `json:"ID"`
	Name string    `json:"name,omitempty"`
}

type Type struct {
	ID   uuid.UUID `json:"ID"`
	Name string    `json:"name,omitempty"`
}

