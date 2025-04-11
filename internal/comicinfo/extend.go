package comicinfo

// Add Tags to the comic info container.
// This function will handle the comma separation automatically.
func (c *ComicInfo) AddTags(tags ...string) {
	c.Tags = AddValue(c.Tags, tags...)
}
