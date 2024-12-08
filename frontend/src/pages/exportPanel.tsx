// React
import { useEffect, useState } from "react";

// React Component
import Button from "react-bootstrap/Button";

// Project Component
import FolderSelector from "../components/FolderSelector";
import { ModalControl } from "../controls/ModalControl";
import GreenRadio from "../components/GreenRadio";

// Wails
import { ExportCbz, GetDefaultOutputDirectory } from "../../wailsjs/go/application/App";
import { comicinfo } from "../../wailsjs/go/models";

/** Props Interface for FolderSelect */
type ExportProps = {
	/** The comic info content. */
	comicInfo: comicinfo.ComicInfo | undefined;
	/** The directory of original input, contains comic images. */
	originalDirectory: string | undefined;
	/** Modal Controller. */
	modalControl: ModalControl;
};

/** Method for export comicinfo cbz. */
enum ExportMethod {
	/** Only export single cbz file. */
	CBZ_ONLY,
	/** Export cbz file with a folder wrapped. */
	FOLDER_WRAP_CBZ,
}

/**
 * The panel to export comic info to cbz/xml file.
 * @returns JSX Component
 */
export default function ExportPanel({ comicInfo: info, originalDirectory, modalControl }: Readonly<ExportProps>) {
	// Since this is the final step, could ignore the interaction with App.tsx
	const [exportDir, setExportDir] = useState<string>("");
	const [exportMethod, setExportMethod] = useState<ExportMethod>(ExportMethod.FOLDER_WRAP_CBZ);

	// Set the export directory to input directory if it exists
	useEffect(() => {
		if (originalDirectory !== undefined) {
			// Load config from file
			GetDefaultOutputDirectory(originalDirectory).then((dir) => {
				setExportDir(dir);
			});
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

		// Start Running
		ExportCbz(originalDirectory, exportDir, info, exportMethod === ExportMethod.FOLDER_WRAP_CBZ).then((msg) => {
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
			/>

			{/* Radio Buttons */}
			<div className="w-50 mx-auto d-grid justify-content-center">
				<GreenRadio
					id="export-type-cbz"
					name="export-type"
					label={"Export .cbz file only"}
					checked={exportMethod === ExportMethod.CBZ_ONLY}
					onChange={() => setExportMethod(ExportMethod.CBZ_ONLY)}
				/>
				<GreenRadio
					id="export-type-wrapped"
					name="export-type"
					label={"Export .cbz wrapped by folder"}
					checked={exportMethod === ExportMethod.FOLDER_WRAP_CBZ}
					onChange={() => setExportMethod(ExportMethod.FOLDER_WRAP_CBZ)}
				/>
			</div>

			{/* Button to Export. Use d-grid to create block button, use w-25 to smaller size. */}
			<div className="w-25 mx-auto d-grid gap-2 mt-4">
				<Button variant="success" id="btn-export" onClick={() => handleExportCbz()}>
					Export
				</Button>
			</div>
		</div>
	);
}
