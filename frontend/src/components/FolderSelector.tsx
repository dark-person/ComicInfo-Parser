import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import InputGroup from "react-bootstrap/InputGroup";

import { GetDirectory } from "../../wailsjs/go/application/App";

type FolderSelectorProps = {
	/** ID of root element. */
	id?: string;
	/** Class name for root element. */
	className?: string;

	/** Text act as label for this folder selector element. */
	label?: string;
	/** Absoulte path of selected directory. */
	directory: string;
	/** React state hook setter for selected directory. */
	setDirectory: (value: string) => void;
	/** Default directory to open if no directory selected. */
	defaultDirectory?: string;

	/** ID of readonly input element, that display selected directory. */
	inputId?: string;
	/** Class name for readonly input element, that display selected directory.*/
	inputClassName?: string;

	/** Id of button that open dialog to select path. */
	buttonId?: string;
	/** Class name of button that open dialog to select path. */
	buttonClassName?: string;
	/** Button variant. */
	buttonVariant?: string;
};

/** UI component for select folder. */
export default function FolderSelector(props: Readonly<FolderSelectorProps>) {
	/** Handler for click "Select Folder". This will open file chooser for choose a file. */
	function handleSelect() {
		let dir = props.directory;

		// Use default directory if no directory selected.
		if (dir === "" && props.defaultDirectory !== undefined) {
			dir = props.defaultDirectory;
		}

		GetDirectory(dir).then((resp) => {
			if (resp.ErrMsg !== "") {
				console.error(resp.ErrMsg);
			}

			let dirToUse = resp.SelectedDir;

			// When user failed to select/cancel, enfore to use default directory if any
			if (dirToUse === "" && props.defaultDirectory !== undefined) {
				dirToUse = props.defaultDirectory;
			}

			props.setDirectory(dirToUse);
		});
	}

	return (
		<InputGroup id={props.id} className={props.className ?? "my-2"}>
			{props.label !== undefined && <InputGroup.Text>{props.label}</InputGroup.Text>}

			<Form.Control
				id={props.inputId}
				className={props.inputClassName}
				type="text"
				placeholder="select folder.."
				value={props.directory}
				readOnly
			/>

			<Button
				id={props.buttonId}
				className={props.buttonClassName}
				variant={props.buttonVariant ?? "secondary"}
				onClick={handleSelect}>
				Select Folder
			</Button>
		</InputGroup>
	);
}
