import { useState } from "react";
import { Button } from "react-bootstrap";

import CollapseCard from "../components/CollapseCard";

/** Props Interface for FolderSelect */
type HelpPanelProps = {
	/** Function for back to home page */
	backToHome: () => void;
};

/** Page for common FAQ / Help section. */
export default function HelpPanel({ backToHome }: Readonly<HelpPanelProps>) {
	const NONE_ACTIVE = -1;

	/** React state for panel index that is currently active. */
	const [active, setActive] = useState<number>(NONE_ACTIVE);

	/**
	 * Function to handle onClick in different panel.
	 * @param index the index of panel, `-1` as no panel to active
	 */
	function handleClick(index: number) {
		if (active === index) {
			setActive(NONE_ACTIVE);
			return;
		}

		setActive(index);
	}

	return (
		<div id="Help-Panel" className="mt-5">
			{/* Main Content of this panel */}
			<h5 className="mb-4">Common FAQ</h5>

			{/* Tutorial/Instruction */}
			<div>
				<CollapseCard
					myKey={0}
					title={'What is "folder structure is not correct"?'}
					body={
						<>
							<p>The folder selected cannot contain any sub-folder, and follow this structure:</p>
							<p>{" 📦 <Manga Name>\n" + " ┣ 📜01.jpg\n" + " ┣ 📜02.jpg\n" + " ┗ <other images>"}</p>
							<p>No ComicInfo.xml is needed. It will be overwrite if exist.</p>
						</>
					}
					isOpen={active === 0}
					onClick={() => handleClick(0)}
				/>
				<CollapseCard
					myKey={1}
					title={'What is "Quick Export (Komga)"?'}
					body={
						<>
							<p>Directly Export .cbz file with ComicInfo.xml inside. The generated file with be like:</p>
							<p>
								{" 📦 <Manga Name>\n" +
									" ┣ 📦 <Manga Name>  <-- Copy This Folder into Komga Comic Library\n" +
									" ┃  ┣  📜<Manga Name>.cbz    <--- Generated .cbz\n" +
									" ┣ 📜01.jpg\n" +
									" ┣ 📜02.jpg\n" +
									" ┣ <other images>\n" +
									" ┗ 📜ComicInfo.xml\n"}
							</p>
						</>
					}
					isOpen={active === 1}
					onClick={() => handleClick(1)}
				/>

				{/** Back to Home page */}
				<Button variant="success" id="btn-return" onClick={backToHome} className="mt-5">
					Back to Select Folder
				</Button>
			</div>
		</div>
	);
}
