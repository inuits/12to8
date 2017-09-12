package api

// Models represents all the lists we can view,
// edit, etc
var Models = &models{}

type models struct {
	models []modelList
}

func (c *models) register(m modelList) {
	c.models = append(c.models, m)
}

// List return the list of models
func (c *models) List() []string {
	var r []string
	for _, m := range c.models {
		r = append(r, m.slug())
	}
	return r
}

// List return the models that match the slug
func (c *models) GetBySlug(slug string) modelList {
	for _, m := range c.models {
		if m.slug() == slug {
			return m
		}
	}
	return nil
}
