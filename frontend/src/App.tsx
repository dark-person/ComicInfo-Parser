// CSS Import
import "./App.css";
import "bootstrap/dist/css/bootstrap.min.css";

// React Component
import React, { useState } from "react";
import Button from "react-bootstrap/Button";
import FolderSelect from "./pages/folderSelect";
import { LoadingModal } from "./modal";
import InputPanel from "./pages/inputPanel";
import { Row, Col } from "react-bootstrap";

const mode_select_folder = 1;
const mode_input_data = 2;

const largest_mode = mode_input_data;

function App() {
	const [mode, setMode] = useState<number>(mode_select_folder);

	const [isLoading, setIsLoading] = useState<boolean>(false);

	function handleConfirm(event: React.MouseEvent) {
		console.log("config clicked");
		setIsLoading(true);
		console.log("Delayed for 2 second.");
		setTimeout(() => {
			setIsLoading(false);
			setMode(mode_input_data);
		}, 2000);
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
						<FolderSelect handleConfirm={handleConfirm} />
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
