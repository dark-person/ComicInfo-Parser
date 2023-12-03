// React Component
import Tab from "react-bootstrap/Tab";
import Tabs from "react-bootstrap/Tabs";
import { Form } from "react-bootstrap";

// Project Specified Component
import { FormDateRow, FormRow } from "../formRow";

/** Button Props Interface for InputPanel */
// type InputProps = {
// 	// returnFunc: (event: React.MouseEvent) => void;
// };

/**
 * The interface for show/edit book metadata.
 * @returns JSX Element
 */
function BookMetadata() {
	return (
		<div>
			<Form>
				<FormRow title={"Title"} disabled />
				<FormRow title={"Summary"} textareaRow={3} disabled />
				<FormRow title={"Number"} inputType="number" disabled />
				<FormDateRow title={"Year/Month/Day"} disabled />
				<FormRow title={"Web"} disabled />
				<FormRow title={"GTIN"} disabled />
			</Form>
		</div>
	);
}

/**
 * The interface for show/edit creator metadata.
 * @returns JSX Element
 */
function CreatorMetadata() {
	return (
		<div>
			<Form>
				<FormRow title={"Writer"} disabled />
				<FormRow title={"Penciller"} disabled />
				<FormRow title={"Inker"} disabled />
				<FormRow title={"Colorist"} disabled />
				<FormRow title={"Letterer"} disabled />
				<FormRow title={"CoverArtist"} disabled />
				<FormRow title={"Editor"} disabled />
				<FormRow title={"Translator"} disabled />
				<FormRow title={"Publisher"} disabled />
			</Form>
		</div>
	);
}

/**
 * The interface for show/edit tags metadata.
 * @returns JSX Element
 */
function TagMetadata() {
	return (
		<div>
			<Form>
				<FormRow title={"Tag"} textareaRow={10} disabled />
			</Form>
		</div>
	);
}

/**
 * The panel for input/edit content of ComicInfo.xml
 * @returns JSX Element
 */
export default function InputPanel() {
	return (
		<div id="Input-Panel" className="mt-5">
			<h5 className="mb-4">Modify ComicInfo.xml</h5>
			<Tabs
				defaultActiveKey="Main"
				id="uncontrolled-tab-example"
				className="mb-3">
				<Tab eventKey="Main" title="Book Metadata">
					<BookMetadata />
				</Tab>

				<Tab eventKey="Creator" title="Creator">
					<CreatorMetadata />
				</Tab>
				<Tab eventKey="Tags" title="Tags">
					<TagMetadata />
				</Tab>
			</Tabs>
		</div>
	);
}
