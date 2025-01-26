import "./exportPanel.css";

// React
import { useEffect, useState } from "react";

// React Component
import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";

// Project Component
import { ColoredRadio } from "../components/ColoredFormCheck";
import FolderSelector from "../components/FolderSelector";
import { ModalControl } from "../controls/ModalControl";

// Wails
import {
	ExportCbzOnly,
	ExportCbzWithDefaultWrap,
	ExportCbzWithWrap,
	GetDefaultOutputDirectory,
} from "../../wailsjs/go/application/App";
import { comicinfo } from "../../wailsjs/go/models";
import { ExportMethod } from "../controls/SessionData";
import { basename } from "../filename";

/** Props Interface for FolderSelect */
type ExportProps = {
	/** The comic info content. */
	comicInfo: comicinfo.ComicInfo | undefined;
	/** The directory of original input, contains comic images. */
	originalDirectory: string | undefined;
	/** Modal Controller. */
	modalControl: ModalControl;
	/** Export method to be used. */
	exportMethod: ExportMethod;
	/** Method to set export method as react hook. */
	setExportMethod: (val: ExportMethod) => void;
};

/**
 * The panel to export comic info to cbz/xml file.
 * @returns JSX Component
 */
export default function ExportPanel({
	comicInfo: info,
	originalDirectory,
	modalControl,
	exportMethod,
	setExportMethod,
}: Readonly<ExportProps>) {
	// Since this is the final step, could ignore the interaction with App.tsx
	const [defaultDir, setDefaultDir] = useState<string>("");
	const [exportDir, setExportDir] = useState<string>("");
	const [customWrap, setCustomWrap] = useState<string>("");

	// Set the export directory to input directory if it exists
	useEffect(() => {
		if (originalDirectory !== undefined) {
			// Load config from file
			GetDefaultOutputDirectory(originalDirectory).then((dir) => {
				setDefaultDir(dir);
				setExportDir(dir);
			});

			// Set custom wrap value
			setCustomWrap(basename(originalDirectory));
		}
	}, []);

	/**
	 * Handler for click export .cbz only, export path will be the folder chosen by file chooser.
	 * @param isWrap is using wrap folder. If true, then export will include a folder warping cbz file, otherwise only cbz file will be exported.
	 * @returns nothing
	 */
	function handleExportCbz() {
		if (originalDirectory === undefined) {
			console.error("No original directory");
			return;
		}

		if (info === undefined) {
			console.error("No original comicinfo");
			return;
		}

		// Open Modal
		modalControl.loading();

		// Decide which promise to use
		let promise: Promise<string>;

		switch (exportMethod) {
			case ExportMethod.CBZ_ONLY:
				promise = ExportCbzOnly(originalDirectory, exportDir, info);
				break;

			case ExportMethod.DEFAULT_WRAP_CBZ:
				promise = ExportCbzWithDefaultWrap(originalDirectory, exportDir, info);
				break;

			case ExportMethod.CUSTOM_WRAP_CBZ:
				if (customWrap === "") {
					modalControl.showErr("Custom Wrap folder cannot be empty");
					return;
				}

				promise = ExportCbzWithWrap(originalDirectory, exportDir, customWrap, info);
				break;

			default:
				modalControl.showErr("Unhandled export method");
				return;
		}

		// Start running
		promise.then((msg) => {
			if (msg !== "") {
				modalControl.showErr(msg);
			} else {
				modalControl.completeAndReset();
			}
			console.log(`cbz return: '${msg}'`);
		});
	}

	return (
		<div id="Export-Panel" className="mt-5">
			{/* Main Content of this panel */}
			<h5 className="mb-4">Export to .cbz</h5>

			{/* File Chooser */}
			<FolderSelector
				className={"mb-3"}
				label={"Export Folder"}
				directory={exportDir}
				setDirectory={setExportDir}
				defaultDirectory={defaultDir}
			/>

			{/* Radio Buttons */}
			<div className="mx-auto ps-9em pe-3em">
				<ColoredRadio
					id="export-type-cbz"
					name="export-type"
					color="dark-orange"
					label={"Export .cbz file only"}
					checked={exportMethod === ExportMethod.CBZ_ONLY}
					onChange={() => setExportMethod(ExportMethod.CBZ_ONLY)}
				/>
				<ColoredRadio
					id="export-type-wrapped"
					name="export-type"
					color="dark-green"
					label={"Export .cbz wrapped by default folder"}
					checked={exportMethod === ExportMethod.DEFAULT_WRAP_CBZ}
					onChange={() => setExportMethod(ExportMethod.DEFAULT_WRAP_CBZ)}
				/>
				<ColoredRadio
					id="export-type-custom-wrapped"
					name="export-type"
					color="dark-blue"
					label={"Export .cbz wrapped by custom folder"}
					checked={exportMethod === ExportMethod.CUSTOM_WRAP_CBZ}
					onChange={() => setExportMethod(ExportMethod.CUSTOM_WRAP_CBZ)}
				/>

				{exportMethod === ExportMethod.CUSTOM_WRAP_CBZ && (
					<Form.Control
						className="ms-1-5em"
						type="text"
						value={customWrap}
						onChange={(e) => setCustomWrap(e.currentTarget.value)}
					/>
				)}
			</div>

			{/* Button to Export. Use d-grid to create block button, use w-25 to smaller size. */}
			<div className="mt-4">
				<Button variant="success" id="btn-export" className="w-25" onClick={() => handleExportCbz()}>
					Export
				</Button>
			</div>
		</div>
	);
}
