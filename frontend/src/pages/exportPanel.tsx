// React Component
import Form from "react-bootstrap/Form";
import Tab from "react-bootstrap/Tab";
import Tabs from "react-bootstrap/Tabs";
import InputGroup from "react-bootstrap/InputGroup";
import Button from "react-bootstrap/Button";
import { useState } from "react";

export default function ExportPanel() {
	return (
		<div id="Export-Panel">
			<Button variant="outline-secondary" id="btn-export-xml">
				Export
			</Button>
		</div>
	);
}
