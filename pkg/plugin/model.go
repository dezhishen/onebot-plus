package plugin

type OnebotPlusPluginSetting struct {
	Key          string
	Description  string
	DefaultValue string
}

type OnebotPlusPluginSettingBuilder struct {
	settings []OnebotPlusPluginSetting
}

func NewOnebotPlusPluginSettingBuilder() *OnebotPlusPluginSettingBuilder {
	return &OnebotPlusPluginSettingBuilder{}
}

func (m *OnebotPlusPluginSettingBuilder) Add(key, description, defaultValue string) *OnebotPlusPluginSettingBuilder {
	m.settings = append(m.settings, OnebotPlusPluginSetting{
		key,
		description,
		defaultValue,
	})
	return m
}

func (m *OnebotPlusPluginSettingBuilder) Build() []OnebotPlusPluginSetting {
	return m.settings
}
