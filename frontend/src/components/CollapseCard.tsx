// React
import { useState } from "react";

// Component
import Card from "react-bootstrap/Card";
import Collapse from "react-bootstrap/Collapse";

/** Props Interface for CollapseCard */
type CardProps = {
	/** the unique key for this component, used to generate id*/
	myKey: number;
	/** the title to display in Card.Title */
	title: string;
	/** the body inside the Card.Body */
	body?: React.ReactNode;
};

/**
 * A Card with collapse functionality. The collapsed content will be shown/hidden when click the card title.
 * @param myKey the unique key for this component, used to generate id
 * @param title the title to display in Card.Title
 * @param body the body inside the Card.Body
 * @returns a Card Component with Collapse ability for card body.
 */
export default function CollapseCard({ myKey, title, body }: Readonly<CardProps>) {
	const [open, setOpen] = useState(false);

	/** Handler for user click card header. Collapse/Open the card. */
	function handleCollapse() {
		setOpen(!open);
	}

	return (
		<Card className="text-start">
			<Card.Header onClick={handleCollapse} aria-controls={"collapse-text-" + String(myKey)} aria-expanded={open}>
				<span className="me-2">{open ? "▼" : ">"}</span>
				{title}
			</Card.Header>
			<Collapse in={open}>
				<div>
					<Card.Body id={"collapse-text-" + String(myKey)}>
						<Card.Text as="div" className="newLine">
							{body}
						</Card.Text>
					</Card.Body>
				</div>
			</Collapse>
		</Card>
	);
}
