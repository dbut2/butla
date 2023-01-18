package configs

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"

	"github.com/dbut2/butla/pkg/configs"
)

func TestConfigs(t *testing.T) {
	configs.SkipLoading()

	files, err := envs.ReadDir(".")
	assert.NoError(t, err)

	r, err := regexp.Compile(`^.*\.yaml$`)
	assert.NoError(t, err)

	for _, file := range files {
		t.Run(file.Name(), func(t *testing.T) {
			assert.Regexp(t, r, file.Name())

			entry, err := envs.ReadFile(file.Name())
			assert.NoError(t, err)

			err = yaml.Unmarshal(entry, &Config{})
			assert.NoError(t, err)
		})
	}
}

func TestLoadConfig(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		c, err := LoadConfig("local")
		assert.NoError(t, err)
		assert.NotNil(t, c)
	})

	t.Run("NotExist", func(t *testing.T) {
		c, err := LoadConfig("fake")
		assert.Error(t, err)
		assert.Nil(t, c)
	})
}
