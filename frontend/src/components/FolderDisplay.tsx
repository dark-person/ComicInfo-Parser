// React Component
import { useState } from "react";
import { Button, Col, InputGroup, Row } from "react-bootstrap";
import Form from "react-bootstrap/Form";

// Project Specified Component
import { OpenFolder } from "../../wailsjs/go/application/App";
import { basename } from "../filename";

/** Props for FolderNameDisplay. */
type FolderNameDisplayProps = {
	/** Actual absolue path for folder. */
	folderPath: string | undefined;
};

/** A UI component for display folder name. */
export default function FolderNameDisplay({ folderPath }: Readonly<FolderNameDisplayProps>) {
	const [isDisabled, setIsDisabled] = useState(false);

	function handleOpen() {
		setIsDisabled(true);

		// 1 second delay before re-enabling button after successful open operation
		setTimeout(() => setIsDisabled(false), 1000);

		if (folderPath === undefined) {
			console.error("Folder name is undefined");
			return;
		}

		OpenFolder(folderPath);
	}

	return (
		<Row className="mb-3">
			<Col sm="11">
				<InputGroup>
					<InputGroup.Text id="folder-display-label" className={"fst-italic"}>
						{"Folder Name"}
					</InputGroup.Text>
					<Form.Control
						id="folder-display-text"
						value={folderPath !== undefined ? basename(folderPath) : "(N/A)"}
						title={"Folder Name"}
						disabled={true}
					/>
					<Button variant="secondary" id="folder-display-open" onClick={handleOpen} disabled={isDisabled}>
						{"🗁"}
					</Button>
				</InputGroup>
			</Col>
		</Row>
	);
}
