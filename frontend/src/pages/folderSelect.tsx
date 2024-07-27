// React
import { useState } from "react";

// Component
import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import InputGroup from "react-bootstrap/InputGroup";

// Project Specific Component
import { CompleteModal, ErrorModal, LoadingModal } from "../components/modal";
import CollapseCard from "../components/CollapseCard";

// Wails
import { GetDirectory, GetDirectoryWithDefault, QuickExportKomga } from "../../wailsjs/go/application/App";

/** Props Interface for FolderSelect */
type FolderProps = {
	/** function called when process to next step. This function is not applied to Quick Export.*/
	processFunc: (folder: string) => void;
};

/**
 * Page for Selecting Folder to process.
 * This page also contains some basic tutorial for folder structure.
 *
 * @param processFunc handler when process button is clicked
 * @returns Page for selecting Folder
 */
export default function FolderSelect({ processFunc: handleFolder }: Readonly<FolderProps>) {
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
		if (directory != "") {
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
			if (err != "") {
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
			<CompleteModal
				show={isCompleted}
				disposeFunc={() => {
					setIsCompleted(false);
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
			<Button variant="success" className="mx-2" id="btn-confirm-folder" onClick={handleProcess}>
				Generate ComicInfo.xml
			</Button>
			<Button variant="outline-info" className="mx-2" id="btn-quick-export" onClick={handleQuickExport}>
				Quick Export (Komga)
			</Button>
			{/* <Button
				variant="secondary"
				className="mx-2"
				onClick={() => setErrMsg("Testing.")}>
				Test
			</Button> */}
			{/* Tutorial/Instruction  */}
			<div className="mt-5">
				<CollapseCard
					myKey={0}
					title={"Example of Your Image Folder"}
					body={
						<>
							<p>{" 📦 <Manga Name>\n" + " ┣ 📜01.jpg\n" + " ┣ 📜02.jpg\n" + " ┗ <other images>"}</p>
							<p>No ComicInfo.xml is needed. It will be overwrite if exist.</p>
						</>
					}
				/>
				<CollapseCard
					myKey={1}
					title={"Quick Export (Komga)"}
					body={
						<>
							<p>Directly Export .cbz file with ComicInfo.xml inside. The generated file with be like:</p>
							<p>
								{" 📦 <Manga Name>\n" +
									" ┣ 📦 <Manga Name>  <-- Copy This Folder into Komga Comic Library\n" +
									" ┃  ┣  📜<Manga Name>.cbz    <--- Generated .cbz\n" +
									" ┣ 📜01.jpg\n" +
									" ┣ 📜02.jpg\n" +
									" ┣ <other images>\n" +
									" ┗ 📜ComicInfo.xml\n"}
							</p>
						</>
					}
				/>
			</div>
		</div>
	);
}
