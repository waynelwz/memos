package api

import (
	"encoding/json"
	"errors"
	"fmt"

	"golang.org/x/exp/slices"
)

type SystemSettingName string

const (
	// SystemSettingServerID is the key type of server id.
	SystemSettingServerID SystemSettingName = "serverId"
	// SystemSettingSecretSessionName is the key type of secret session name.
	SystemSettingSecretSessionName SystemSettingName = "secretSessionName"
	// SystemSettingAllowSignUpName is the key type of allow signup setting.
	SystemSettingAllowSignUpName SystemSettingName = "allowSignUp"
	// SystemSettingDisablePublicMemosName is the key type of disable public memos setting.
	SystemSettingDisablePublicMemosName SystemSettingName = "disablePublicMemos"
	// SystemSettingAdditionalStyleName is the key type of additional style.
	SystemSettingAdditionalStyleName SystemSettingName = "additionalStyle"
	// SystemSettingAdditionalScriptName is the key type of additional script.
	SystemSettingAdditionalScriptName SystemSettingName = "additionalScript"
	// SystemSettingCustomizedProfileName is the key type of customized server profile.
	SystemSettingCustomizedProfileName SystemSettingName = "customizedProfile"
	// SystemSettingStorageServiceIDName is the key type of storage service ID.
	SystemSettingStorageServiceIDName SystemSettingName = "storageServiceId"
	// SystemSettingLocalStoragePathName is the key type of local storage path.
	SystemSettingLocalStoragePathName SystemSettingName = "localStoragePath"
	// SystemSettingOpenAIConfigName is the key type of OpenAI config.
	SystemSettingOpenAIConfigName SystemSettingName = "openAIConfig"
)

// CustomizedProfile is the struct definition for SystemSettingCustomizedProfileName system setting item.
type CustomizedProfile struct {
	// Name is the server name, default is `memos`
	Name string `json:"name"`
	// LogoURL is the url of logo image.
	LogoURL string `json:"logoUrl"`
	// Description is the server description.
	Description string `json:"description"`
	// Locale is the server default locale.
	Locale string `json:"locale"`
	// Appearance is the server default appearance.
	Appearance string `json:"appearance"`
	// ExternalURL is the external url of server. e.g. https://usermemos.com
	ExternalURL string `json:"externalUrl"`
}

type OpenAIConfig struct {
	Key  string `json:"key"`
	Host string `json:"host"`
}

func (key SystemSettingName) String() string {
	switch key {
	case SystemSettingServerID:
		return "serverId"
	case SystemSettingSecretSessionName:
		return "secretSessionName"
	case SystemSettingAllowSignUpName:
		return "allowSignUp"
	case SystemSettingDisablePublicMemosName:
		return "disablePublicMemos"
	case SystemSettingAdditionalStyleName:
		return "additionalStyle"
	case SystemSettingAdditionalScriptName:
		return "additionalScript"
	case SystemSettingCustomizedProfileName:
		return "customizedProfile"
	case SystemSettingStorageServiceIDName:
		return "storageServiceId"
	case SystemSettingLocalStoragePathName:
		return "localStoragePath"
	case SystemSettingOpenAIConfigName:
		return "openAIConfig"
	}
	return ""
}

type SystemSetting struct {
	Name SystemSettingName `json:"name"`
	// Value is a JSON string with basic value.
	Value       string `json:"value"`
	Description string `json:"description"`
}

type SystemSettingUpsert struct {
	Name        SystemSettingName `json:"name"`
	Value       string            `json:"value"`
	Description string            `json:"description"`
}

func (upsert SystemSettingUpsert) Validate() error {
	if upsert.Name == SystemSettingServerID {
		return errors.New("update server id is not allowed")
	} else if upsert.Name == SystemSettingAllowSignUpName {
		value := false
		err := json.Unmarshal([]byte(upsert.Value), &value)
		if err != nil {
			return fmt.Errorf("failed to unmarshal system setting allow signup value")
		}
	} else if upsert.Name == SystemSettingDisablePublicMemosName {
		value := false
		err := json.Unmarshal([]byte(upsert.Value), &value)
		if err != nil {
			return fmt.Errorf("failed to unmarshal system setting disable public memos value")
		}
	} else if upsert.Name == SystemSettingAdditionalStyleName {
		value := ""
		err := json.Unmarshal([]byte(upsert.Value), &value)
		if err != nil {
			return fmt.Errorf("failed to unmarshal system setting additional style value")
		}
	} else if upsert.Name == SystemSettingAdditionalScriptName {
		value := ""
		err := json.Unmarshal([]byte(upsert.Value), &value)
		if err != nil {
			return fmt.Errorf("failed to unmarshal system setting additional script value")
		}
	} else if upsert.Name == SystemSettingCustomizedProfileName {
		customizedProfile := CustomizedProfile{
			Name:        "memos",
			LogoURL:     "",
			Description: "",
			Locale:      "en",
			Appearance:  "system",
			ExternalURL: "",
		}
		err := json.Unmarshal([]byte(upsert.Value), &customizedProfile)
		if err != nil {
			return fmt.Errorf("failed to unmarshal system setting customized profile value")
		}
		if !slices.Contains(UserSettingLocaleValue, customizedProfile.Locale) {
			return fmt.Errorf("invalid locale value")
		}
		if !slices.Contains(UserSettingAppearanceValue, customizedProfile.Appearance) {
			return fmt.Errorf("invalid appearance value")
		}
	} else if upsert.Name == SystemSettingStorageServiceIDName {
		value := 0
		err := json.Unmarshal([]byte(upsert.Value), &value)
		if err != nil {
			return fmt.Errorf("failed to unmarshal system setting storage service id value")
		}
		return nil
	} else if upsert.Name == SystemSettingLocalStoragePathName {
		value := ""
		err := json.Unmarshal([]byte(upsert.Value), &value)
		if err != nil {
			return fmt.Errorf("failed to unmarshal system setting local storage path value")
		}
	} else if upsert.Name == SystemSettingOpenAIConfigName {
		value := OpenAIConfig{}
		err := json.Unmarshal([]byte(upsert.Value), &value)
		if err != nil {
			return fmt.Errorf("failed to unmarshal system setting openai api config value")
		}
	} else {
		return fmt.Errorf("invalid system setting name")
	}

	return nil
}

type SystemSettingFind struct {
	Name SystemSettingName `json:"name"`
}
