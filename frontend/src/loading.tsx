import Modal from "react-bootstrap/Modal";

/**
 * A Modal that display "Please wait..." message and block output.
 * @param show determine the modal to display or not
 * @returns React Function Component
 */
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
