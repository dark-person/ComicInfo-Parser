// React
import { useState } from "react";

// Component
import Form from "react-bootstrap/Form";
import InputGroup from "react-bootstrap/InputGroup";
import Button from "react-bootstrap/Button";
import Card from "react-bootstrap/Card";
import Collapse from "react-bootstrap/Collapse";
import { CompleteModal, ErrorModal, LoadingModal } from "../modal";

// Wails
import {
	GetDirectory,
	QuickExportKomga,
	GetDirectoryWithDefault,
} from "../../wailsjs/go/main/App";

/** Button Props Interface for FolderSelect */
type ButtonProps = {
	handleConfirm: (event: React.MouseEvent) => void;
};

/**
 * A Card with collapse functionality. The collapsed content will be shown/hidden when click the card title.
 * @param key the unique key for this component, used to generate id
 * @param title the title to display in Card.Title
 * @param body the body inside the Card.Body
 * @returns a Card Component with Collapse ability for card body.
 */
function CollapseCard(props: {
	myKey: number;
	title: string;
	body?: React.ReactNode;
}) {
	const [open, setOpen] = useState(false);

	function handleCollapse() {
		setOpen(!open);
	}

	return (
		<Card className="text-start">
			<Card.Header
				onClick={handleCollapse}
				aria-controls={"collapse-text-" + String(props.myKey)}
				aria-expanded={open}>
				<span className="me-2">{open == true ? "▼" : ">"}</span>
				{props.title}
			</Card.Header>
			<Collapse in={open}>
				<div>
					<Card.Body id={"collapse-text-" + String(props.myKey)}>
						<Card.Text as="div" className="newLine">
							{props.body}
						</Card.Text>
					</Card.Body>
				</div>
			</Collapse>
		</Card>
	);
}

/**
 * Page for Selecting Folder to process.
 * This page also contains some basic tutorial for folder structure.
 *
 * @param handleConfirm handler when confirm button is clicked
 * @returns Page for selecting Folder
 */
export default function FolderSelect({ handleConfirm }: ButtonProps) {
	const [directory, setDirectory] = useState("");
	const [isLoading, setIsLoading] = useState(false);
	const [isCompleted, setIsCompleted] = useState(false);
	const [errMsg, setErrMsg] = useState("");

	function handleSelect() {
		if (directory != "") {
			GetDirectoryWithDefault(directory).then((input) => {
				setDirectory(input);
			});
		} else {
			GetDirectory().then((input) => {
				setDirectory(input);
			});
		}
	}

	function handleQuickExport() {
		setIsLoading(true);

		QuickExportKomga(directory).then((err) => {
			setIsLoading(false);
			if (err != "") {
				setErrMsg(err);
			} else {
				setIsCompleted(true);
			}
		});
	}

	return (
		<div id="Folder-Select" className="mt-5">
			<LoadingModal show={isLoading} />
			<CompleteModal
				show={isCompleted}
				disposeFunc={() => {
					setIsCompleted(false);
					return {};
				}}
			/>
			<ErrorModal
				show={errMsg != ""}
				errorMessage={errMsg}
				disposeFunc={() => {
					setErrMsg("");
					return {};
				}}
			/>

			<h5 className="mb-4">Select Folder to Start:</h5>
			<InputGroup className="mb-3">
				<InputGroup.Text>Image Folder</InputGroup.Text>
				<Form.Control
					aria-describedby="btn-select-folder"
					type="text"
					placeholder="select folder.."
					value={directory}
					readOnly
				/>
				<Button
					variant="secondary"
					id="btn-select-folder"
					onClick={handleSelect}>
					Select Folder
				</Button>
			</InputGroup>
			<Button
				variant="success"
				className="mx-2"
				id="btn-confirm-folder"
				onClick={handleConfirm}>
				Generate ComicInfo.xml
			</Button>
			<Button
				variant="outline-info"
				className="mx-2"
				id="btn-quick-export"
				onClick={handleQuickExport}>
				Quick Export (Komga)
			</Button>
			{/* <Button
				variant="secondary"
				className="mx-2"
				onClick={() => setErrMsg("Testing.")}>
				Test
			</Button> */}

			<div className="mt-5">
				<CollapseCard
					myKey={0}
					title={"Example of Your Image Folder"}
					body={
						<>
							<p>
								{" 📦 <Manga Name>\n" +
									" ┣ 📜01.jpg\n" +
									" ┣ 📜02.jpg\n" +
									" ┗ <other images>"}
							</p>
							<p>No ComicInfo.xml is needed. It will be overwrite if exist.</p>
						</>
					}
				/>
				<CollapseCard
					myKey={1}
					title={"Quick Export (Komga)"}
					body={
						<>
							<p>
								Directly Export .cbz file with ComicInfo.xml inside. The
								generated file with be like:
							</p>
							<p>
								{" 📦 <Manga Name>\n" +
									" ┣ 📦 <Manga Name>  <-- Copy This Folder into Komga Comic Library\n " +
									" ┃  ┣  📜<Manga Name>.cbz    <--- Generated .cbz\n" +
									" ┣ 📜01.jpg\n" +
									" ┣ 📜02.jpg\n" +
									" ┣ <other images>\n" +
									" ┗ 📜ComicInfo.xml\n"}
							</p>
						</>
					}
				/>
			</div>
		</div>
	);
}
