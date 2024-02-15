// CSS Import
import "bootstrap/dist/css/bootstrap.min.css";
import "./App.css";

// React Component
import { useState } from "react";
import { Col, Row } from "react-bootstrap";
import Button from "react-bootstrap/Button";

// Project Specified Component
import { ErrorModal, LoadingModal } from "./components/modal";
import FolderSelect from "./pages/folderSelect";
import InputPanel from "./pages/inputPanel";

// Wails
import { GetComicInfo } from "../wailsjs/go/main/App";
import { comicinfo } from "../wailsjs/go/models";
import ExportPanel from "./pages/exportPanel";

/** Const, for Display page of select folder */
const mode_select_folder = 1;
/** Const, for Display page of input panel */
const mode_input_data = 2;
/** Const, for Display page of export panel */
const mode_export = 3;

/**
 * The main component to be displayed. It will handle all pages data & timing to display.
 *
 * There will have two column with width=1 at left and right side, for center the panel/page component.
 *
 * There has a return button at the top-left corner in this window.
 */
function App() {
	/** Decide which panel will be displayed */
	const [mode, setMode] = useState<number>(mode_select_folder);

	/** True if need to display loading dialog. Should be show when change to another page. */
	const [isLoading, setIsLoading] = useState<boolean>(false);

	/** Error Message, will modal be display when not empty string. Empty Strings mean not error at all. */
	const [errMsg, setErrMsg] = useState<string>("");

	/** The ComicInfo model. For communicate with different panel. */
	const [info, setInfo] = useState<comicinfo.ComicInfo | undefined>(undefined);

	/** The directory of initial input, which is the folder contain image1. */
	const [inputDir, setInputDir] = useState<string | undefined>(undefined);

	/**
	 * Set value of selected folder, then pass selected folder to next panel.
	 * @param folder the absolute path to the folder
	 */
	function passingFolder(folder: string) {
		console.log("passing folder: " + folder);

		// Set Loading Modal
		setIsLoading(true);

		// Get ComicInfo
		GetComicInfo(folder).then((response) => {
			// Remove loading modal
			setIsLoading(false);

			let error = response.ErrorMessage;
			if (error != "") {
				// Print Error Message
				setErrMsg(error);
			} else {
				// Set data with info
				setInfo(response.ComicInfo);
				setInputDir(folder);

				// Pass to another panel
				setMode(mode_input_data);
			}
		});
	}

	/** Change the panel in app to export panel. */
	function showExportPanel() {
		setMode(mode_export);
	}

	/**
	 * Return to previous page.
	 */
	function backward() {
		// Get Current Mode
		let temp = mode;

		// Perform Mode subtraction
		temp = Math.max(1, temp - 1);

		// Set Mode
		setMode(temp);
	}

	/**
	 * Return to the home panel. In current version, it is select folder panel.
	 */
	function backToHomePanel() {
		setMode(mode_select_folder);
	}

	/**
	 * Set the value by its field name.
	 * @param data the data to be modify
	 * @param key  the field name
	 * @param value new value of that field
	 */
	function setValue<T, K extends keyof T>(data: T, key: K, value: any) {
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
		let temp = { ...info } as comicinfo.ComicInfo;

		// Treat Summary field name as special case
		if (field === "Summary" && typeof value === "string") {
			temp["Summary"]["InnerXML"] = value;
		} else {
			// Normal Change
			let key = field as keyof comicinfo.ComicInfo;
			setValue(temp, key, value);
		}

		// Set the changed value to data
		setInfo(temp);
	}

	return (
		<div id="App" className="container-fluid">
			{/* Modal Part */}
			<LoadingModal show={isLoading} />
			<ErrorModal
				show={errMsg != ""}
				errorMessage={errMsg}
				disposeFunc={() => {
					setErrMsg("");
					return {};
				}}
			/>

			{/* Main Panel of this app */}
			<Row className="min-vh-100">
				{/* Back Button, return to previous panel */}
				<Col xs={1} className="mt-4">
					{mode > 1 && (
						<Button variant="secondary" onClick={backward}>
							{"<"}
						</Button>
					)}
				</Col>

				{/* Area to display panel */}
				<Col>
					{mode == mode_select_folder && <FolderSelect processFunc={passingFolder} />}
					{mode == mode_input_data && (
						<InputPanel comicInfo={info} exportFunc={showExportPanel} infoSetter={infoSetter} />
					)}
					{mode == mode_export && (
						<ExportPanel comicInfo={info} originalDirectory={inputDir} backToHomeFunc={backToHomePanel} />
					)}
				</Col>

				{/* Button to next panel, development only, otherwise use as alignment */}
				<Col xs={1} className="align-self-center">
					{/* <Button
						variant="danger"
						onClick={() => {
							// Perform Mode subtraction
							let temp = Math.min(mode + 1, mode_export);

							// Set Mode
							setMode(temp);
						}}>
						{">"}
					</Button> */}
				</Col>
			</Row>
		</div>
	);
}

export default App;
