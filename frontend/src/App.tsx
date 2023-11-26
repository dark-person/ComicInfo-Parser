// CSS Import
import "./App.css";
import "bootstrap/dist/css/bootstrap.min.css";

// React Component
import { useState } from "react";
import Form from "react-bootstrap/Form";
import Tab from "react-bootstrap/Tab";
import Tabs from "react-bootstrap/Tabs";
import InputGroup from "react-bootstrap/InputGroup";
import Button from "react-bootstrap/Button";
import FolderSelect from "./pages/folderSelect";
import LoadingModal from "./loading";

function App() {
	const [isLoading, setIsLoading] = useState<boolean>(false);

	function handleConfirm(event: React.MouseEvent) {
		console.log("config clicked");
		setIsLoading(true);

		setTimeout(() => {
			console.log("Delayed for 5 second.");
			setIsLoading(false);
		}, 5000);
	}

	return (
		<div id="App" className="container-fluid">
			<FolderSelect handleConfirm={handleConfirm} />
			<LoadingModal show={isLoading} />
		</div>
	);
}

export default App;
