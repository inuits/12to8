package api

// Models represents all the lists we can view,
// edit, etc
var Models = &models{}

type models struct {
	Models           []ModelList
	IndividualModels []Model
}

func (c *models) register(m ModelList, i Model) {
	c.Models = append(c.Models, m)
	c.IndividualModels = append(c.IndividualModels, i)
}

// List return the models that match the slug
func (c *models) GetBySlug(slug string) ModelList {
	for _, m := range c.Models {
		if m.Slug() == slug {
			return m
		}
	}
	return nil
}
