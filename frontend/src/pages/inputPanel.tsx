// React Component
import Tab from "react-bootstrap/Tab";
import Tabs from "react-bootstrap/Tabs";
import InputGroup from "react-bootstrap/InputGroup";
import Button from "react-bootstrap/Button";
import { useState } from "react";
import { Row, Col } from "react-bootstrap";

/** Button Props Interface for InputPanel */
type InputProps = {
	// returnFunc: (event: React.MouseEvent) => void;
};

function FormRow(props: { title: string; inputType?: string; value?: string }) {
	return (
		<Form.Group as={Row} className="mb-3">
			<Form.Label column sm="2">
				{props.title}
			</Form.Label>
			<Col sm="9">
				<Form.Control type={props.inputType} value={props.value} />
			</Col>
		</Form.Group>
	);
}

function BookMetadata() {
	return (
		<div>
			<Form>
				<FormRow title={"Title"} />
				<FormRow title={"Summary"} />
				<FormRow title={"Number"} inputType="number" />
				<FormRow title={"Year"} inputType="number" />
				<FormRow title={"Month"} inputType="number" />
				<FormRow title={"Day"} inputType="number" />
				<FormRow title={"Web"} />
				<FormRow title={"GTIN"} />
			</Form>
		</div>
	);
}

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
					Tab content for Profile
				</Tab>
				<Tab eventKey="Tags" title="Tags" disabled>
					Tab content for Contact
				</Tab>
			</Tabs>
		</div>
	);
}
