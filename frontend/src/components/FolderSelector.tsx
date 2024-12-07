import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import InputGroup from "react-bootstrap/InputGroup";

import { GetDirectory, GetDirectoryWithDefault } from "../../wailsjs/go/application/App";

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
		if (props.directory !== "") {
			GetDirectoryWithDefault(props.directory).then((input) => {
				props.setDirectory(input);
			});
		} else {
			GetDirectory().then((input) => {
				props.setDirectory(input);
			});
		}
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
