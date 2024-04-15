package domain

// MARK: - Objects

type Language struct {
	Tag         string `json:"tag"`
	Flag        string `json:"flag"`
	DisplayName string `json:"displayName"`
}

type TranslatableError interface {
	RawKey() Translatable
}

type TranslatableArgument struct {
	Key string `json:"key"`
}

type Translatable struct {
	Key       string `json:"key"`
	Arguments []TranslatableArgument
}

type TranslatableArgumentMap map[TranslatableArgument]interface{}

// MARK: - Translatable

var (
	TranslatableDefaultOrganization = Translatable{Key: "DefaultOrganizationDescription", Arguments: []TranslatableArgument{TranslatableArgumentUserFullName}}
)

// MARK: - Translatable Args

var (
	TranslatableArgumentUserFullName = TranslatableArgument{Key: "UserFullName"}
)

// MARK: - Functions

func (t Translatable) RawKey() string {
	return t.Key
}
