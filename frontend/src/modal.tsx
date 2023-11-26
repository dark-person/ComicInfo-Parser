import Modal from "react-bootstrap/Modal";
import Button from "react-bootstrap/Button";

/**
 * A Modal that display "Please wait..." message and block output.
 * @param show determine the modal to display or not
 * @returns React Function Component
 */
export function LoadingModal({ show }: { show: boolean }) {
	return (
		<Modal
			show={show}
			size="lg"
			aria-labelledby="contained-modal-title-vcenter"
			backdrop="static"
			keyboard={false}
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

/**
 * A Modal that display Success Complete Message.
 * @param show determine the modal to show or not
 * @param disposeFunc the function to dispose this modal. Usually like setShow(false).
 * @returns React Function Component
 */
export function CompleteModal({
	show,
	disposeFunc,
}: {
	show: boolean;
	disposeFunc: () => {};
}) {
	return (
		<Modal
			show={show}
			size="sm"
			aria-labelledby="contained-modal-title-vcenter"
			backdrop="static"
			keyboard={false}
			centered>
			<Modal.Header className="justify-content-center">
				<Modal.Title id="contained-modal-title-vcenter">Completed!</Modal.Title>
			</Modal.Header>
			<Modal.Body className="text-nowrap">
				<p>The Process is complete successfully.</p>
			</Modal.Body>
			<Modal.Footer className="justify-content-center">
				<Button variant="success" onClick={disposeFunc}>
					Close
				</Button>
			</Modal.Footer>
		</Modal>
	);
}

/**
 * A Modal that display Success Error Message.
 * @param show determine the modal to show or not
 * @param errorMessage the error message to displayed
 * @param disposeFunc the function to dispose this modal. Usually like setShow(false).
 * @returns React Function Component
 */
export function ErrorModal({
	show,
	errorMessage,
	disposeFunc,
}: {
	show: boolean;
	errorMessage: string;
	disposeFunc: () => {};
}) {
	return (
		<Modal
			show={show}
			aria-labelledby="contained-modal-title-vcenter"
			backdrop="static"
			keyboard={false}
			centered>
			<Modal.Header className="justify-content-center text-danger-emphasis">
				<Modal.Title id="contained-modal-title-vcenter">Failed!</Modal.Title>
			</Modal.Header>
			<Modal.Body className="text-nowrap text-start">
				<p>The Process is failed because: </p>
				<p>{errorMessage}</p>
			</Modal.Body>
			<Modal.Footer className="justify-content-center">
				<Button variant="secondary" onClick={disposeFunc}>
					Close
				</Button>
			</Modal.Footer>
		</Modal>
	);
}
