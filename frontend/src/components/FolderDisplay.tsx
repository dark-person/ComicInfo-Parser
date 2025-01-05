// React Component
import { Col, Row } from "react-bootstrap";
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
		<Form.Group as={Row} className="mb-3">
			<Form.Label column sm="2" className={"fst-italic"}>
				{"Folder Name"}
			</Form.Label>
			<Col sm="9">
				<Form.Control
					value={folderPath !== undefined ? basename(folderPath) : "(N/A)"}
					title={"Folder Name"}
					disabled={true}
				/>
			</Col>
		</Form.Group>
	);
}
