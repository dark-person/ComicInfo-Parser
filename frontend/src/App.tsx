// CSS Import
import "./App.css";
import "bootstrap/dist/css/bootstrap.min.css";

// React Component
import React, { useEffect, useState } from "react";
import Button from "react-bootstrap/Button";
import { Row, Col } from "react-bootstrap";

// Project Specified Component
import FolderSelect from "./pages/folderSelect";
import { LoadingModal } from "./modal";
import InputPanel from "./pages/inputPanel";
import { DataPass } from "./data";

const mode_select_folder = 1;
const mode_input_data = 2;

function App() {
	const [mode, setMode] = useState<number>(mode_select_folder);

	const [isLoading, setIsLoading] = useState<boolean>(false);

	const [data, setData] = useState<DataPass | undefined>(undefined);

	/**
	 * Set value of selected folder. Used in communicate with other components.
	 * @param folder the absolute path to the folder
	 */
	function passingFolder(folder: string) {
		console.log("passing folder: " + folder);

		let temp: DataPass;
		if (data != undefined) {
			temp = { folder: folder, comicInfo: data.comicInfo };
			setData(temp);
		} else {
			temp = { folder: folder, comicInfo: undefined };
			setData(temp);
		}
	}

	// TODO: Only for debugging purposes, should be removed later
	useEffect(() => {
		console.log("after pass:" + JSON.stringify(data, null, 4));
	}, [data]);

	// Handling confirm button for select folder
	function handleConfirm() {
		console.log("confirm clicked");
		// setIsLoading(true);

		// console.log("Delayed for 2 second.");
		// setTimeout(() => {
		// 	setIsLoading(false);
		// 	setMode(mode_input_data);
		// }, 2000);
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
			<Row className="min-vh-100">
				<Col xs={1} className="mt-4">
					{mode > 1 && (
						<Button variant="secondary" onClick={backward}>
							{"<"}
						</Button>
					)}
				</Col>
				<Col>
					{mode == mode_select_folder && (
						<FolderSelect
							handleConfirm={handleConfirm}
							handleFolder={passingFolder}
						/>
					)}
					{mode == mode_input_data && <InputPanel />}
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
