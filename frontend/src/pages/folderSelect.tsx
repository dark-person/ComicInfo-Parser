// React Component
import Form from "react-bootstrap/Form";
import Tab from "react-bootstrap/Tab";
import Tabs from "react-bootstrap/Tabs";
import InputGroup from "react-bootstrap/InputGroup";
import Button from "react-bootstrap/Button";
import { useState } from "react";
import Card from "react-bootstrap/Card";
import Collapse from "react-bootstrap/Collapse";

type ButtonProps = {
	handleConfirm: (event: React.MouseEvent) => void;
};

export default function FolderSelect({ handleConfirm }: ButtonProps) {
	const [open, setOpen] = useState(false);

	return (
		<div id="Folder-Select" className="mt-5">
			<h5 className="mb-4">Select Folder to Start:</h5>
			<InputGroup className="mb-3">
				<InputGroup.Text>Image Folder</InputGroup.Text>
				<Form.Control
					aria-describedby="btn-select-folder"
					type="text"
					placeholder="select folder.."
					readOnly
				/>
				<Button variant="secondary" id="btn-select-folder">
					Select Folder
				</Button>
			</InputGroup>
			<Button variant="success" id="btn-confirm-folder" onClick={handleConfirm}>
				Confirm
			</Button>

			<div className="mt-5">
				<Card className="text-start">
					<Card.Header
						onClick={() => setOpen(!open)}
						aria-controls="example-collapse-text"
						aria-expanded={open}>
						<span className="me-2">{open == true ? "▼" : ">"}</span>
						Example of Your Image Folder
					</Card.Header>
					<Collapse in={open}>
						<Card.Body id="example-collapse-text">
							<Card.Text className="newLine">
								<p>
									{" 📦 <Manga Name>\n" +
										" ┣ 📜01.jpg\n" +
										" ┣ 📜02.jpg\n" +
										" ┗ <other images>"}
								</p>
								<p>
									No ComicInfo.xml is needed. It will be overwrite if exist.
								</p>
							</Card.Text>
						</Card.Body>
					</Collapse>
				</Card>
			</div>
		</div>
	);
}
