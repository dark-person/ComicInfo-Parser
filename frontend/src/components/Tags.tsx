import "./Tags.css";

type TagsAreaProps = {
	rawTags: string | undefined;
};

export function TagsArea({ rawTags }: TagsAreaProps) {
	/**
	 * Parse the raw string that contains tags into arrays.
	 * @param raw raw string of tags, e.g `"tag1, tag2"`
	 * @returns array of tags, e.g. `["tag1", "tag2"]`
	 */
	function getTagsList(raw: string | undefined): string[] {
		if (raw === undefined) return [];

		return raw.split(",");
	}

	return (
		<div className="tag-area text-start p-2  ">
			<div className="d-inline-flex flex-wrap">
				{getTagsList(rawTags).map((item, index) => (
					<Tag tag={item} key={index} />
				))}
			</div>
		</div>
	);
}

type TagProps = {
	tag: string;
};

export function Tag({ tag }: TagProps) {
	return <div className="me-1">{tag}</div>;
}
