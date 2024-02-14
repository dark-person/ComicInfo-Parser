import { Button } from "react-bootstrap";
import "./Tags.css";

type TagsAreaProps = {
	rawTags: string | undefined;
	handleDelete: (arg0: number) => void;
};

export function TagsArea({ rawTags, handleDelete }: TagsAreaProps) {
	/**
	 * Parse the raw string that contains tags into arrays.
	 * <p>
	 * Please note that, split may able to return `[""]` strings array,
	 * which may result generated a empty tag item.
	 * Therefore, this function will turn this case in to empty array instead.
	 *
	 * @param raw raw string of tags, e.g `"tag1, tag2"`
	 * @returns array of tags, e.g. `["tag1", "tag2"]`
	 */
	function getTagsList(raw: string | undefined): string[] {
		if (raw === undefined) return [];

		let temp = raw.split(",");

		// Prevent Empty Tag list
		if (temp.length === 1 && temp[0].length == 0) {
			console.log("empty tag list");
			return [];
		}

		return temp;
	}

	return (
		<div className="tag-area text-start p-2  ">
			<div className="d-inline-flex flex-wrap">
				{getTagsList(rawTags).map((item, index) => (
					<Tag tag={item} key={index} index={index} handleDelete={handleDelete} />
				))}
			</div>
		</div>
	);
}

type TagProps = {
	tag: string;
	index: number;
	handleDelete: (arg0: number) => void;
};

export function Tag({ tag, index, handleDelete }: TagProps) {
	return (
		<div className="me-1 mb-1 tag-item bg-secondary p-1">
			<span className="p-1">{tag}</span>
			<Button className="btn-close remove-icon" onClick={() => handleDelete(index)}></Button>
		</div>
	);
}
