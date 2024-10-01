// CSS Import
import "bootstrap/dist/css/bootstrap.min.css";
import "./App.css";

// React Component
import { useState } from "react";
import { Col, Row } from "react-bootstrap";
import Button from "react-bootstrap/Button";

// Project Specified Component
import { CompleteModal, ErrorModal, LoadingModal } from "./components/modal";
import { AppMode } from "./controls/AppMode";
import { ModalControl } from "./controls/ModalControl";
import { defaultModalState, ModalState } from "./controls/ModalState";
import ExportPanel from "./pages/exportPanel";
import FolderSelect from "./pages/folderSelect";
import HelpPanel from "./pages/helpPanel";
import InputPanel from "./pages/inputPanel";

// Wails
import { GetComicInfo } from "../wailsjs/go/application/App";
import { comicinfo } from "../wailsjs/go/models";

/**
 * The main component to be displayed. It will handle all pages data & timing to display.
 *
 * There will have two column with width=1 at left and right side, for center the panel/page component.
 *
 * There has a return button at the top-left corner in this window.
 */
function App() {
	/** Decide which panel will be displayed */
	const [mode, setMode] = useState<AppMode>(AppMode.SELECT_FOLDER);

	/** Modal State to control display which dialog. */
	const [modalState, setModalState] = useState<ModalState>(defaultModalState);

	/** The ComicInfo model. For communicate with different panel. */
	const [info, setInfo] = useState<comicinfo.ComicInfo | undefined>(undefined);

	/** The directory of initial input, which is the folder contain image. */
	const [inputDir, setInputDir] = useState<string | undefined>(undefined);

	/** Controller of modal. */
	const modalController: ModalControl = {
		showErr: (err) => setModalState({ ...defaultModalState, errMsg: err }),
		loading: () => setModalState({ ...defaultModalState, isLoading: true }),
		complete: () => setModalState({ ...defaultModalState, isCompleted: true }),
		completeAndReset: () => setModalState({ ...defaultModalState, isCompleted: true, resetOnComplete: true }),
		closeAll: () => setModalState({ ...defaultModalState }),
	};

	/**
	 * Set value of selected folder, then pass selected folder to next panel.
	 * @param folder the absolute path to the folder
	 */
	function passingFolder(folder: string) {
		console.log("passing folder: " + folder);

		// Set Loading Modal
		modalController.loading();

		// Get ComicInfo
		GetComicInfo(folder).then((response) => {
			const error = response.ErrorMessage;

			if (error !== "") {
				// Print Error Message
				modalController.showErr(error);
			} else {
				// Reset all modal
				modalController.closeAll();

				// Set data with info
				setInfo(response.ComicInfo);
				setInputDir(folder);

				// Pass to another panel
				setMode(AppMode.INPUT_DATA);
			}
		});
	}

	/** Change the panel in app to export panel. */
	function showExportPanel() {
		setMode(AppMode.EXPORT);
	}

	/** Show help panel in App. */
	function showHelpPanel() {
		setMode(AppMode.HELP);
	}

	/**
	 * Return to previous page.
	 * Only `AppMode.INPUT_DATA` & `AppMode.EXPORT` is supported.
	 */
	function backward() {
		switch (mode) {
			case AppMode.INPUT_DATA:
				setMode(AppMode.SELECT_FOLDER);
				return;
			case AppMode.EXPORT:
				setMode(AppMode.INPUT_DATA);
				return;
			default:
				throw new Error("Invalid mode");
		}
	}

	/** Return to the home panel. In current version, it is select folder panel. */
	function backToHomePanel() {
		setMode(AppMode.SELECT_FOLDER);
	}

	/**
	 * Set the value by its field name.
	 * @param data the data to be modify
	 * @param key  the field name
	 * @param value new value of that field
	 */
	function setValue<T, K extends keyof T>(data: T, key: K, value: T[K]) {
		data[key] = value;
	}

	/**
	 * Setter for changing value of comicInfo.
	 *
	 * Note that this function will treat "Summary" field as special case.
	 * @param field the field name, must be same as ComicInfo field name
	 * @param value new value of that field
	 */
	function infoSetter(field: string, value: string | number) {
		// Prepare an object of ComicInfo
		const temp = { ...info } as comicinfo.ComicInfo;

		// Treat Summary field name as special case
		if (field === "Summary" && typeof value === "string") {
			temp["Summary"]["InnerXML"] = value;
		} else {
			// Normal Change
			const key = field as keyof comicinfo.ComicInfo;
			setValue(temp, key, value);
		}

		// Set the changed value to data
		setInfo(temp);
	}

	return (
		<div id="App" className="container-fluid">
			{/* Modal Part */}
			<LoadingModal show={modalState.isLoading} />

			<ErrorModal
				show={modalState.errMsg !== ""}
				errorMessage={modalState.errMsg}
				disposeFunc={modalController.closeAll}
			/>

			<CompleteModal
				show={modalState.isCompleted}
				disposeFunc={() => {
					modalController.closeAll();

					// Redirect to first page only if export cbz is clicked
					if (modalState.resetOnComplete) {
						backToHomePanel();
					}
				}}
			/>

			{/* Main Panel of this app */}
			<Row className="min-vh-100">
				{/* Back Button, return to previous panel */}
				<Col xs={1} className="mt-4">
					{/* Only Allow backward when export page / input data page */}
					{(mode === AppMode.EXPORT || mode === AppMode.INPUT_DATA) && (
						<Button variant="secondary" onClick={backward}>
							{"<"}
						</Button>
					)}
				</Col>

				{/* Area to display panel */}
				<Col>
					{mode === AppMode.SELECT_FOLDER && (
						<FolderSelect
							handleFolder={passingFolder}
							showHelpPanel={showHelpPanel}
							modalControl={modalController}
						/>
					)}
					{mode === AppMode.INPUT_DATA && (
						<InputPanel
							comicInfo={info}
							exportFunc={showExportPanel}
							infoSetter={infoSetter}
							folderName={inputDir}
						/>
					)}
					{mode === AppMode.EXPORT && (
						<ExportPanel comicInfo={info} originalDirectory={inputDir} backToHomeFunc={backToHomePanel} />
					)}
					{mode === AppMode.HELP && <HelpPanel backToHome={backToHomePanel} />}
				</Col>

				{/* Use as alignment */}
				<Col xs={1} className="align-self-center"></Col>
			</Row>
		</div>
	);
}

export default App;
