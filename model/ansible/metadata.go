package ansible

type Metadata struct {
	Dependencies []interface{}       `yaml:"dependencies"`
	GalaxyInfo   struct {
		RoleName          string     `yaml:"role_name"`
		Author            string     `yaml:"author"`
		Description       string     `yaml:"description"`
		Company           string     `yaml:"company"`
		License           string     `yaml:"license"`
		MinAnsibleVersion float64    `yaml:"min_ansible_version"`
		GalaxyTags        []string   `yaml:"galaxy_tags"`
		Platforms         []struct {
			Name          string     `yaml:"name"`
			Versions      []string   `yaml:"versions"`
		}                            `yaml:"platforms"`
	}                                `yaml:"galaxy_info"`
}
