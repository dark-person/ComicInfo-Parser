// React
import { useEffect, useState } from "react";

// React Component
import Button from "react-bootstrap/Button";

// Project Component
import FolderSelector from "../components/FolderSelector";
import { ModalControl } from "../controls/ModalControl";

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

/**
 * The panel to export comic info to cbz/xml file.
 * @returns JSX Component
 */
export default function ExportPanel({ comicInfo: info, originalDirectory, modalControl }: Readonly<ExportProps>) {
	// Since this is the final step, could ignore the interaction with App.tsx
	const [exportDir, setExportDir] = useState<string>("");

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
	function handleExportCbz(isWrap: boolean) {
		if (originalDirectory === undefined) {
			console.log("[ERR] No original directory");
			return;
		}

		if (info === undefined) {
			console.log("[ERR] No original comicinfo");
			return;
		}

		// Open Modal
		modalControl.loading();

		// Start Running
		ExportCbz(originalDirectory, exportDir, info, isWrap).then((msg) => {
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

			{/* Button to Export. Use d-grid to create block button, use w-25 to smaller size. */}
			<div className="w-25 mx-auto d-grid gap-2">
				<Button variant="outline-warning" id="btn-export-xml" onClick={() => handleExportCbz(false)}>
					Export .cbz file only
				</Button>

				<Button variant="outline-info" id="btn-export-xml" onClick={() => handleExportCbz(true)}>
					Export whole .cbz folder
				</Button>
			</div>
		</div>
	);
}
