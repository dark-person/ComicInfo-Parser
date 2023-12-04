// CSS Import
import "./App.css";
import "bootstrap/dist/css/bootstrap.min.css";

// React Component
import React, { useEffect, useState } from "react";
import Button from "react-bootstrap/Button";
import { Row, Col } from "react-bootstrap";

// Project Specified Component
import { DataPass } from "./data";
import { ErrorModal, LoadingModal } from "./modal";
import FolderSelect from "./pages/folderSelect";
import InputPanel from "./pages/inputPanel";

// Wails
import { GetComicInfo } from "../wailsjs/go/main/App";

const mode_select_folder = 1;
const mode_input_data = 2;

function App() {
	const [mode, setMode] = useState<number>(mode_select_folder);

	const [isLoading, setIsLoading] = useState<boolean>(false);
	const [errMsg, setErrMsg] = useState<string>("");

	const [data, setData] = useState<DataPass | undefined>(undefined);

	/**
	 * Set value of selected folder. Used in communicate with other components.
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
				let temp = { folder: folder, comicInfo: response.ComicInfo };
				setData(temp);

				// Pass to another panel
				setMode(mode_input_data);
			}
		});
	}

	/**
	 * Return to previous page.
	 * @param event React.MouseEvent
	 */
	function backward(event: React.MouseEvent) {
		// Get Current Mode
		let temp = mode;

		// Perform Mode subtraction
		temp = Math.max(1, temp - 1);

		// Set Mode
		setMode(temp);
	}

	return (
		<div id="App" className="container-fluid">
			<ErrorModal
				show={errMsg != ""}
				errorMessage={errMsg}
				disposeFunc={() => {
					setErrMsg("");
					return {};
				}}
			/>

			<Row className="min-vh-100">
				<Col xs={1} className="mt-4">
					{mode > 1 && (
						<Button variant="secondary" onClick={backward}>
							{"<"}
						</Button>
					)}
				</Col>
				<Col>
					{mode == mode_select_folder && <FolderSelect handleFolder={passingFolder} />}
					{mode == mode_input_data && <InputPanel comicInfo={data?.comicInfo} />}
				</Col>
				<Col xs={1} className="align-self-center">
					{/* <Button variant="secondary">{">"}</Button> */}
				</Col>
			</Row>

			<LoadingModal show={isLoading} />
		</div>
	);
}

export default App;
