package requests

type NewOrganization struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}
