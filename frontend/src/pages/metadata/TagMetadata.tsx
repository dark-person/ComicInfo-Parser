// React
import { ChangeEvent, useState } from "react";

// React Component
import { Button, Col, Form, Row } from "react-bootstrap";

// Project Specified Component
import { TagsArea } from "../../components/Tags";
import { MetadataProps } from "./MetadataProps";

/**
 * The interface for show/edit tags metadata.
 * @returns JSX Element
 */
export default function TagMetadata({ comicInfo: info, infoSetter }: Readonly<MetadataProps>) {
	/** Hooks of tag that to be added. Only allow single tag to be added at a time. */
	const [singleTag, setSingleTag] = useState<string>("");

	/** Function for handing enter key pressed when entering custom tags. */
	const handleKeyDown = (event: React.KeyboardEvent) => {
		if (event.key === "Enter") {
			event.preventDefault();
			handleAdd();
		}
	};

	/**
	 * Function to handle the textfield of tag to be added.
	 * @param e the react event
	 */
	function handleChanges(e: ChangeEvent<HTMLInputElement>) {
		setSingleTag(e.target.value);
	}

	/** Function for handling add button click */
	function handleAdd() {
		// Prevent Empty tags
		if (singleTag === "") {
			return;
		}

		// Retrieve Tags
		let temp = info?.Tags;

		// Append Tags to comic info
		if (temp === "" || temp === undefined) {
			temp = singleTag;
		} else {
			temp += ", " + singleTag;
		}

		// Apply Change to ComicInfo
		infoSetter("Tags", temp);

		// Empty Tags in textfield
		setSingleTag("");
	}

	/** Function for handling delete button click */
	function handleDelete(id: number) {
		if (info?.Tags === undefined) {
			return;
		}

		// Parse tags to array of strings
		let array = info.Tags.split(",");

		// Remove tag by index
		array.splice(id, 1);

		// Concat array of string to string
		let str = array.join(", ");

		// Set by info setter
		infoSetter("Tags", str);
	}

	return (
		<div>
			<Form>
				{/* A Text Area for holding lines of tags. */}
				<Row>
					<Col sm={2} className="mt-1">
						{"Tags"}
					</Col>
					<Col sm={9}>
						<TagsArea rawTags={info?.Tags} handleDelete={handleDelete} />
					</Col>
				</Row>

				{/* Empty Rows for margin */}
				<Row className="mb-3" />

				<Row>
					{/* Empty Column */}
					<Col sm={2} className="mt-1">
						Add Custom Tag
					</Col>

					{/* Column of adding tags */}
					<Col sm={8}>
						<Form.Control value={singleTag} onChange={handleChanges} onKeyDown={handleKeyDown} />
					</Col>

					{/* Column of add button */}
					<Col sm={1}>
						<Button variant="outline-info" onClick={handleAdd}>
							Add
						</Button>
					</Col>
				</Row>
			</Form>
		</div>
	);
}
