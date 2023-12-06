// React
import { useEffect, useState } from "react";

// React Component
import Form from "react-bootstrap/Form";
import InputGroup from "react-bootstrap/InputGroup";
import Button from "react-bootstrap/Button";
import { Row, Col } from "react-bootstrap";

// Wails
import {
	GetDirectory,
	GetDirectoryWithDefault,
	ExportXml,
	ExportCbz,
} from "../../wailsjs/go/main/App";
import { comicinfo } from "../../wailsjs/go/models";
import { LoadingModal, CompleteModal, ErrorModal } from "../modal";

/** Props Interface for FolderSelect */
type ExportProps = {
	comicInfo: comicinfo.ComicInfo | undefined;
	originalDirectory: string | undefined;
	backToHomeFunc: () => void;
};

/** The modal state constant. */
type ModalState = "loading" | "complete" | undefined;

/** The button name state constant, for last clicked button */
type buttonName = "xml" | "cbz" | undefined;

export default function ExportPanel({
	comicInfo: info,
	originalDirectory,
	backToHomeFunc,
}: ExportProps) {
	// Since this is the final step, could ignore the interaction with App.tsx
	const [exportDir, setExportDir] = useState<string>("");

	// Modal Controller
	const [modalState, setModalState] = useState<ModalState>(undefined);
	const [errMsg, setErrMsg] = useState<string>("");

	// The button name of last clicked button
	const [btnClicked, setBtnClicked] = useState<buttonName>(undefined);

	// Set the export directory to input directory if it exists
	useEffect(() => {
		if (originalDirectory != undefined) {
			setExportDir(originalDirectory);
		}
	}, []);

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

		// Open Modal
		setModalState("loading");

		// Set button state
		setBtnClicked("xml");

		// Start Running
		ExportXml(originalDirectory, info).then((msg) => {
			if (msg != "") {
				setErrMsg(msg);
				setModalState(undefined);
			} else {
				setModalState("complete");
			}

			console.log("xml return: '" + msg + "'");
		});
	}

	function handleExportCbz() {
		if (originalDirectory == undefined) {
			console.log("[ERR] No original directory");
			return;
		} else if (info == undefined) {
			console.log("[ERR] No original comicinfo");
			return;
		}

		// Open Modal
		setModalState("loading");

		// Set button state
		setBtnClicked("cbz");

		// Start Running
		ExportCbz(originalDirectory, exportDir, info).then((msg) => {
			if (msg != "") {
				setErrMsg(msg);
				setModalState(undefined);
			} else {
				setModalState("complete");
			}
			console.log("cbz return: '" + msg + "'");
		});
	}

	return (
		<div id="Export-Panel" className="mt-5">
			{/* Modal Part */}
			<LoadingModal show={modalState === "loading"} />
			<CompleteModal
				show={modalState === "complete"}
				disposeFunc={() => {
					setModalState(undefined);
					// Redirect to first page only if export cbz is clicked
					if (btnClicked == "cbz") {
						backToHomeFunc();
					}
					return {};
				}}
			/>
			<ErrorModal
				show={errMsg != ""}
				errorMessage={errMsg}
				disposeFunc={() => {
					setErrMsg("");
					return {};
				}}
			/>

			{/* Main Content of this panel */}
			<h5 className="mb-4">Export to .cbz</h5>

			{/* File Chooser */}
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

			{/* Button to Export */}
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
					<Button
						variant="outline-secondary"
						id="btn-export-xml"
						onClick={handleExportCbz}>
						Export whole .cbz folder
					</Button>
				</Col>
			</Row>
		</div>
	);
}
