package comicinfo

// Add Tags to the comic info container.
// This function will handle the comma separation automatically.
func (c *ComicInfo) AddTags(tags ...string) {
	c.Tags = AddValue(c.Tags, tags...)
}

func (c *ComicInfo) AddGenre(genre ...string) {
	c.Genre = AddValue(c.Genre, genre...)
}

func (c *ComicInfo) AddPublisher(publisher ...string) {
	c.Publisher = AddValue(c.Publisher, publisher...)
}

func (c *ComicInfo) AddTranslator(translator ...string) {
	c.Translator = AddValue(c.Translator, translator...)
}
