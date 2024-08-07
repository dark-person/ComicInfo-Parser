// React Component
import { Button } from "react-bootstrap";
import Tab from "react-bootstrap/Tab";
import Tabs from "react-bootstrap/Tabs";

// Project Specified Component
import { basename } from "../filename";
import FormRow from "../forms/FormRow";
import BookMetadata from "./metadata/BookMetadata";
import CreatorMetadata from "./metadata/CreatorMetadata";
import MiscMetadata from "./metadata/MiscMetadata";
import SeriesMetadata from "./metadata/SeriesMetadata";
import TagMetadata from "./metadata/TagMetadata";

// Wails binding
import { comicinfo } from "../../wailsjs/go/models";

/** Props Interface for InputPanel */
type InputProps = {
	/** The comic info object. Accept undefined value. */
	comicInfo: comicinfo.ComicInfo | undefined;

	/** The folder name for reference. */
	folderName?: string;

	/** The function to change display panel to export panel. */
	exportFunc: () => void;

	/** The info Setter. This function should be doing setting value, but no verification. */
	infoSetter: (field: string, value: string | number) => void;
};

/**
 * The panel for input/edit content of ComicInfo.xml
 * @returns JSX Element
 */
export default function InputPanel({ comicInfo, folderName, exportFunc, infoSetter }: Readonly<InputProps>) {
	return (
		<div id="Input-Panel" className="mt-5">
			<h5 className="mb-4">Modify ComicInfo.xml</h5>

			{/* Component for showing folder name (with basename only) */}
			<FormRow
				title={"Folder Name"}
				titleClass="fst-italic"
				value={folderName != undefined ? basename(folderName) : "(N/A)"}
				disabled
			/>

			{/* The Tabs Group to display metadata. */}
			<Tabs defaultActiveKey="Main" id="uncontrolled-tab-example" className="mb-3">
				<Tab eventKey="Main" title="Book Metadata">
					<BookMetadata comicInfo={comicInfo} infoSetter={infoSetter} />
				</Tab>

				<Tab eventKey="Creator" title="Creator">
					<CreatorMetadata comicInfo={comicInfo} infoSetter={infoSetter} />
				</Tab>

				<Tab eventKey="Tags" title="Tags">
					<TagMetadata comicInfo={comicInfo} infoSetter={infoSetter} />
				</Tab>

				<Tab eventKey="Series" title="Series">
					<SeriesMetadata comicInfo={comicInfo} infoSetter={infoSetter} />
				</Tab>

				<Tab eventKey="Misc" title="Collection & ReadList">
					<MiscMetadata comicInfo={comicInfo} infoSetter={infoSetter} />
				</Tab>
			</Tabs>

			{/* The button that will always at the bottom of screen. Should ensure there has enough space */}
			<div className="fixed-bottom mb-3">
				<Button variant="success" className="mx-2 " id="btn-export-cbz" onClick={exportFunc}>
					Export to .cbz
				</Button>
			</div>
		</div>
	);
}
