// React
import { useState } from "react";

// Component
import Button from "react-bootstrap/Button";

// Project Specific Component
import FolderSelector from "../components/FolderSelector";
import { ModalControl } from "../controls/ModalControl";

// Wails
import { QuickExportKomga } from "../../wailsjs/go/application/App";

/** Props Interface for FolderSelect */
type FolderProps = {
	/** function called when process to next step. This function is not applied to Quick Export.*/
	handleFolder: (folder: string) => void;

	/** Function to called when help button clicked. */
	showHelpPanel: () => void;

	/** Modal controller. */
	modalControl: ModalControl;
};

/** Page for Selecting Folder to process. */
export default function FolderSelect({ handleFolder, showHelpPanel, modalControl }: Readonly<FolderProps>) {
	/** The Directory Absolute Path selected by User. */
	const [directory, setDirectory] = useState("");

	/** Handler when user clicked Quick Export Komga Button. Start quick export process. */
	function handleQuickExport() {
		modalControl.loading();

		QuickExportKomga(directory).then((err) => {
			if (err !== "") {
				modalControl.showErr(err);
			} else {
				modalControl.complete();
			}
		});
	}

	/** Handle when user click "Generate ComicInfo". Move to another pages for display/edit comicinfo content. */
	function handleProcess() {
		handleFolder(directory);
	}

	return (
		<div id="Folder-Select" className="mt-5">
			{/* Main Content Start*/}
			<h5 className="mb-4">Select Folder to Start:</h5>

			{/* Folder Chooser */}
			<FolderSelector
				className="mb-3"
				label="Image Folder"
				buttonId="btn-select-folder"
				directory={directory}
				setDirectory={setDirectory}
			/>

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
