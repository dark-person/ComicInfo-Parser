// React Component
import Form from "react-bootstrap/Form";
import InputGroup from "react-bootstrap/InputGroup";
import Button from "react-bootstrap/Button";
import { Row, Col } from "react-bootstrap";

// Wails
import { GetDirectory } from "../../wailsjs/go/main/App";

export default function ExportPanel() {
	return (
		<div id="Export-Panel" className="mt-5">
			<h5 className="mb-4">Export to .cbz</h5>

			<InputGroup className="mb-3">
				<InputGroup.Text>Export Folder</InputGroup.Text>
				<Form.Control
					aria-describedby="btn-select-export-folder"
					type="text"
					placeholder="select folder.."
					readOnly
				/>
				<Button variant="secondary" id="btn-select-export-folder">
					Select Folder
				</Button>
			</InputGroup>

			<Row className="mb-3">
				<Col>
					<Button variant="outline-secondary" id="btn-export-xml">
						Export ComicInfo.xml Only
					</Button>
				</Col>
			</Row>

			<Row className="mb-3">
				<Col>
					<Button variant="outline-secondary" id="btn-export-xml">
						Export whole .cbz folder
					</Button>
				</Col>
			</Row>
		</div>
	);
}
