// React
import { useState } from "react";

// Component
import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import InputGroup from "react-bootstrap/InputGroup";

// Project Specific Component
import { CompleteModal, ErrorModal, LoadingModal } from "../components/modal";

// Wails
import { GetDirectory, GetDirectoryWithDefault, QuickExportKomga } from "../../wailsjs/go/application/App";

/** Props Interface for FolderSelect */
type FolderProps = {
	/** function called when process to next step. This function is not applied to Quick Export.*/
	handleFolder: (folder: string) => void;

	/** Function to called when help button clicked. */
	showHelpPanel: () => void;
};

/** Page for Selecting Folder to process. */
export default function FolderSelect({ handleFolder, showHelpPanel }: Readonly<FolderProps>) {
	/** The Directory Absolute Path selected by User. */
	const [directory, setDirectory] = useState("");

	/** True if loading screen should appear */
	const [isLoading, setIsLoading] = useState(false);

	/** True if completed screen should appear */
	const [isCompleted, setIsCompleted] = useState(false);

	/** Error Message for completed process. If has string, then appear modal for display message. */
	const [errMsg, setErrMsg] = useState("");

	/** Handler when user clicked select folder. It will use different function depend on there are already selected folder or not */
	function handleSelect() {
		if (directory !== "") {
			GetDirectoryWithDefault(directory).then((input) => {
				setDirectory(input);
			});
		} else {
			GetDirectory().then((input) => {
				setDirectory(input);
			});
		}
	}

	/** Handler when user clicked Quick Export Komga Button. Start quick export process. */
	function handleQuickExport() {
		setIsLoading(true);

		QuickExportKomga(directory).then((err) => {
			setIsLoading(false);
			if (err !== "") {
				setErrMsg(err);
			} else {
				setIsCompleted(true);
			}
		});
	}

	/** Handle when user click "Generate ComicInfo". Move to another pages for display/edit comicinfo content. */
	function handleProcess() {
		handleFolder(directory);
	}

	return (
		<div id="Folder-Select" className="mt-5">
			{/* Model Part */}
			<LoadingModal show={isLoading} />
			<CompleteModal show={isCompleted} disposeFunc={() => setIsCompleted(false)} />
			<ErrorModal show={errMsg !== ""} errorMessage={errMsg} disposeFunc={() => setErrMsg("")} />

			{/* Main Content Start*/}
			<h5 className="mb-4">Select Folder to Start:</h5>

			{/* Folder Chooser */}
			<InputGroup className="mb-3">
				<InputGroup.Text>Image Folder</InputGroup.Text>
				<Form.Control
					aria-describedby="btn-select-folder"
					type="text"
					placeholder="select folder.."
					value={directory}
					readOnly
				/>
				<Button variant="secondary" id="btn-select-folder" onClick={handleSelect}>
					Select Folder
				</Button>
			</InputGroup>

			{/* Button Group */}
			<div className="w-25 mx-auto d-grid gap-2">
				<Button variant="success" id="btn-confirm-folder" onClick={handleProcess}>
					Generate ComicInfo.xml
				</Button>

				<Button variant="outline-info" id="btn-quick-export" onClick={handleQuickExport}>
					Quick Export (Komga)
				</Button>

				{/* Tutorial/Instruction */}
				<Button variant="outline-warning" id="btn-help" onClick={showHelpPanel}>
					Help
				</Button>
			</div>
		</div>
	);
}
