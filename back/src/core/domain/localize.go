package domain

type Language struct {
	Tag         string `json:"tag"`
	Flag        string `json:"flag"`
	DisplayName string `json:"displayName"`
}

type TranslatableArgument struct {
	Key string `json:"key"`
}

type Translatable struct {
	Key       string `json:"key"`
	Arguments []TranslatableArgument
}

func (t Translatable) RawKey() string {
	return t.Key
}

type TranslatableArgumentMap map[TranslatableArgument]interface{}

// Translation keys.
var (
	TranslatableDefaultOrganization = Translatable{Key: "DefaultOrganizationDescription", Arguments: []TranslatableArgument{TranslatableArgumentUserFullName}}
)

// Translation Arguments
var (
	TranslatableArgumentUserFullName = TranslatableArgument{Key: "UserFullName"}
)
