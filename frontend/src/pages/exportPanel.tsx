// React
import { useState } from "react";

// React Component
import Form from "react-bootstrap/Form";
import InputGroup from "react-bootstrap/InputGroup";
import Button from "react-bootstrap/Button";
import { Row, Col } from "react-bootstrap";

// Wails
import { GetDirectory, GetDirectoryWithDefault, ExportXml } from "../../wailsjs/go/main/App";
import { comicinfo } from "../../wailsjs/go/models";

/** Props Interface for FolderSelect */
type ExportProps = {
	comicInfo: comicinfo.ComicInfo | undefined;
	originalDirectory: string | undefined;
};

export default function ExportPanel({ comicInfo: info, originalDirectory }: ExportProps) {
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

	function handleExportXml() {
		if (originalDirectory == undefined) {
			console.log("[ERR] No original directory");
			return;
		} else if (info == undefined) {
			console.log("[ERR] No original comicinfo");
			return;
		}

		ExportXml(originalDirectory, info).then((msg) => {
			console.log("xml return: " + msg);
		});
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
					<Button
						variant="outline-secondary"
						id="btn-export-xml"
						onClick={handleExportXml}>
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
