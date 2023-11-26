import { useState } from "react";
import "./App.css";
import Form from "react-bootstrap/Form";
import Tab from "react-bootstrap/Tab";
import Tabs from "react-bootstrap/Tabs";
import InputGroup from "react-bootstrap/InputGroup";
import Button from "react-bootstrap/Button";

import "bootstrap/dist/css/bootstrap.min.css";

function App() {
	return (
		<div id="App">
			<div id="Folder-Select">
				<InputGroup className="mb-3">
					<Button variant="outline-secondary" id="btn-select-folder">
						Select Folder
					</Button>
					<Form.Control
						aria-describedby="btn-select-folder"
						type="text"
						placeholder="select folder.."
						readOnly
					/>
				</InputGroup>
			</div>
			<div id="Input-Panel">
				<Tabs
					defaultActiveKey="Main"
					id="uncontrolled-tab-example"
					className="mb-3">
					<Tab eventKey="Main" title="Main">
						Tab content for Home
					</Tab>
					<Tab eventKey="Creator" title="Creator">
						Tab content for Profile
					</Tab>
					<Tab eventKey="Tags" title="Tags" disabled>
						Tab content for Contact
					</Tab>
				</Tabs>
			</div>
			<div id="Export-Panel">
				<Button variant="outline-secondary" id="btn-export-xml">
					Export
				</Button>
			</div>
		</div>
	);
}

export default App;
