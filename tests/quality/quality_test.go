package quality

import (
	"go/ast"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type variableRule struct {
	Prefix string
	Type   string
}

var variableRules = []variableRule{
	{Prefix: "LogScope", Type: "LogScope"},
	{Prefix: "LogField", Type: "LogField"},
	{Prefix: "LogLevel", Type: "LogLevel"},
	{Prefix: "TranslatableArgument", Type: "TranslatableArgument"},
	{Prefix: "Translatable", Type: "Translatable"},
	{Prefix: "PermissionScope", Type: "PermissionScope"},
	{Prefix: "Permission", Type: "Permission"},
	{Prefix: "Role", Type: "Role"},
}

type variable struct {
	Name string
	Type string

	data map[string]interface{}
}

func TestQuality(t *testing.T) {
	vars, err := getVariables("../../internal/core/domain")
	assert.NoError(t, err)

	t.Run("Checking variable names", func(t *testing.T) {
		for _, v := range vars {
			found := false
			for _, rule := range variableRules {
				if strings.HasPrefix(v.Name, rule.Prefix) {
					assert.Equal(t, rule.Type, v.Type, "Variable named:"+v.Name+" should be of type "+rule.Type+"(current type: "+v.Type+")")
					found = true
					break
				}
				if found {
					break
				}
			}
		}
	})

	t.Run("Checking variable types", func(t *testing.T) {
		for _, v := range vars {
			found := false
			for _, rule := range variableRules {
				if v.Type == rule.Type && !strings.HasPrefix(v.Name, rule.Prefix) {
					assert.Equal(t, rule.Type, v.Type, "Variable named:"+v.Name+" with type "+rule.Type+" should be prefixed with: "+rule.Prefix)
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	})

	t.Run("All declared translatable should have a matching translation", func(t *testing.T) {
		files := listLanguageFiles("../../localization")
		assert.GreaterOrEqual(t, len(files), 1, "There should be at least one language file in the localization directory")

		for _, filePath := range files {
			file, err := loadTomlFile(filePath)
			assert.NoError(t, err)
			assert.NotNil(t, file)

			fileName := strings.TrimRight(filePath, "/")
			fileName = strings.Split(fileName, "/")[len(strings.Split(fileName, "/"))-1]

			for _, v := range vars {
				if v.Type == "Translatable" {
					kv, ok := v.data["fields"].(*ast.KeyValueExpr)
					assert.True(t, ok)

					key, keyIsOk := kv.Key.(*ast.Ident)
					value, valueIsOk := kv.Value.(*ast.BasicLit)

					if key.Name == "Key" && keyIsOk && valueIsOk {
						field := value.Value[1 : len(value.Value)-1]
						str, ok := file[field]
						assert.True(t, ok, fileName+": missing key: "+field)
						if ok {
							assert.NotEmpty(t, str, fileName+": empty value for key: "+field)
						}
					}

				}
			}
		}
	})
}
