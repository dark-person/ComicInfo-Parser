import Button from "react-bootstrap/Button";
import Modal from "react-bootstrap/Modal";

type LoadingModalProps = {
	/** determine the modal to display or not */
	show: boolean;
};

/**
 * A Modal that display "Please wait..." message and block output.
 * @returns React Function Component
 */
export function LoadingModal({ show }: Readonly<LoadingModalProps>) {
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

/** The Props for CompleteModal */
type CompleteModalProps = {
	/**  determine the modal to show or not*/
	show: boolean;
	/** the function to dispose this modal. Usually like setShow(false). */
	disposeFunc: () => {};
};

/**
 * A Modal that display Success Complete Message.
 * @param show determine the modal to show or not
 * @param disposeFunc
 * @returns React Function Component
 */
export function CompleteModal({ show, disposeFunc }: Readonly<CompleteModalProps>) {
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

/** The Props for Error Modal. */
type ErrorModalProps = {
	/** determine the modal to show or not */
	show: boolean;
	/** the error message to displayed */
	errorMessage: string;
	/** the function to dispose this modal. Usually like setShow(false) */
	disposeFunc: () => {};
};

/**
 * A Modal that display Success Error Message.
 * @returns React Function Component
 */
export function ErrorModal({ show, errorMessage, disposeFunc }: Readonly<ErrorModalProps>) {
	/**
	 * Convert Error Message Text to Human readable string with foot stop & Capital Letter
	 * @param msg the original Error Message
	 * @returns the human readable with foot stop & Capital Letter
	 */
	function humanReadable(msg: string): string {
		return `${msg.charAt(0).toUpperCase() + msg.slice(1)}.`;
	}

	return (
		<Modal show={show} aria-labelledby="contained-modal-title-vcenter" backdrop="static" keyboard={false} centered>
			<Modal.Header className="justify-content-center text-danger-emphasis">
				<Modal.Title id="contained-modal-title-vcenter">Failed!</Modal.Title>
			</Modal.Header>
			<Modal.Body className="text-nowrap text-start">
				<p>The Process is failed because: </p>
				<p>{humanReadable(errorMessage)}</p>
			</Modal.Body>
			<Modal.Footer className="justify-content-center">
				<Button variant="secondary" onClick={disposeFunc}>
					Close
				</Button>
			</Modal.Footer>
		</Modal>
	);
}
