// React
import { useState } from "react";

// React Component
import Form from "react-bootstrap/Form";
import InputGroup from "react-bootstrap/InputGroup";
import Button from "react-bootstrap/Button";
import { Row, Col } from "react-bootstrap";

// Wails
import { GetDirectory, GetDirectoryWithDefault } from "../../wailsjs/go/main/App";

export default function ExportPanel() {
	// Since this is the final step, could ignore the interaction with App.tsx
	const [exportDir, setExportDir] = useState("");

	function handleSelect() {
		if (exportDir != "") {
			GetDirectoryWithDefault(exportDir).then((input) => {
				setExportDir(input);
			});
		} else {
			GetDirectory().then((input) => {
				setExportDir(input);
			});
		}
	}
	return (
		<div id="Export-Panel" className="mt-5">
			<h5 className="mb-4">Export to .cbz</h5>

			<InputGroup className="mb-3">
				<InputGroup.Text>Export Folder</InputGroup.Text>
				<Form.Control
					aria-describedby="btn-select-export-folder"
					type="text"
					placeholder="select folder.."
					value={exportDir}
					readOnly
				/>
				<Button variant="secondary" id="btn-select-export-folder" onClick={handleSelect}>
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
