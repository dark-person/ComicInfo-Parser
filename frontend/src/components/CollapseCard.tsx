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
    /** Flag control collapse part to be open or not. Should be React state.*/
    isOpen: boolean;
    /**
     * Function called when user clicked on card header.
     *
     * Developer are suggested to include set `isOpen` value inside.
     */
    onClick: () => void;
};

/**
 * A Card with collapse functionality.
 * The collapsed content will be shown/hidden when click the card title.
 */
export default function CollapseCard({ myKey, title, body, isOpen, onClick }: Readonly<CardProps>) {
    return (
        <Card className="text-start">
            <Card.Header onClick={onClick} aria-controls={"collapse-text-" + String(myKey)} aria-expanded={isOpen}>
                <span className="me-2">{isOpen ? "▼" : ">"}</span>
                {title}
            </Card.Header>
            <Collapse in={isOpen}>
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
