import Form from "react-bootstrap/Form";
import Tab from "react-bootstrap/Tab";
import Tabs from "react-bootstrap/Tabs";
import InputGroup from "react-bootstrap/InputGroup";
import Button from "react-bootstrap/Button";
import { useState } from "react";
import Card from "react-bootstrap/Card";
import Collapse from "react-bootstrap/Collapse";
import Modal from "react-bootstrap/Modal";

export default function LoadingModal({ show }: { show: boolean }) {
	return (
		<Modal
			show={show}
			size="lg"
			aria-labelledby="contained-modal-title-vcenter"
			centered>
			<Modal.Header>
				<Modal.Title id="contained-modal-title-vcenter">Loading</Modal.Title>
			</Modal.Header>
			<Modal.Body>
				<p>Please Wait...</p>
			</Modal.Body>
		</Modal>
	);
}
