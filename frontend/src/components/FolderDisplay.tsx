// React Component
import { Button, Col, InputGroup, Row } from "react-bootstrap";
import Form from "react-bootstrap/Form";

// Project Specified Component
import { basename } from "../filename";

/** Props for FolderNameDisplay. */
type FolderNameDisplayProps = {
	/** Actual absolue path for folder. */
	folderPath: string | undefined;
};

/** A UI component for display folder name. */
export default function FolderNameDisplay({ folderPath }: Readonly<FolderNameDisplayProps>) {
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
					<Button variant="secondary" id="folder-display-open">
						{"🗁"}
					</Button>
				</InputGroup>
			</Col>
		</Row>
	);
}
