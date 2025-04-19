// React Component
import { Col, Row } from "react-bootstrap";
import Form from "react-bootstrap/Form";

/** Props for `FormSelectRow`. */
type FormSelectRowProps = {
    /** the title/label of this input group */
    title: string;
    /** The JSX.Element of `<select>`. */
    selectElement: React.JSX.Element;
};

/**
 * Create a uniform Form.Group Element as Row.
 * This element specify designed for select element.
 */
export default function FormSelectRow({ title, selectElement }: Readonly<FormSelectRowProps>) {
    return (
        <Form.Group as={Row} className="mb-3">
            <Form.Label column sm="2">
                {title}
            </Form.Label>
            <Col sm="9">{selectElement}</Col>
        </Form.Group>
    );
}
